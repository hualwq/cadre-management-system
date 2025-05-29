package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PositionHistory struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID      string `gorm:"not null;column:user_id" json:"user_id"`                   // Associates with the cadre's basic info
	Name         string `gorm:"size:100;not null" json:"name"`                            // 姓名
	PhoneNumber  string `gorm:"size:20" json:"phone_number"`                              // 电话号码
	Email        string `gorm:"size:100" json:"email"`                                    // 电子邮件
	Department   string `gorm:"size:100;not null" json:"department"`                      // 院系
	Category     string `gorm:"size:50;not null" json:"category"`                         // 类别: 专职团干部/兼职团干部/教师/学生
	Office       string `gorm:"type:ENUM('校团委内设部门','学生会','研究生会');not null" json:"office"` // 任职部门
	AcademicYear string `gorm:"size:50;not null" json:"academic_year"`                    // 任职年度格式: "2023-2024第一学期"
	Positions    string `gorm:"size:200" json:"positions"`                                // 职位名称
	Year         uint   `gorm:"not null;column:applied_at_year;type:int unsigned" json:"applied_at_year"`
	Month        uint   `gorm:"column:applied_at_month;type:tinyint unsigned" json:"applied_at_month"`
	Day          uint   `gorm:"column:applied_at_day;type:tinyint unsigned" json:"applied_at_day"`
}

type Posexp struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID    string `gorm:"size:50;column:user_id" json:"user_id"`
	Posyear    string `gorm:"size:20" json:"year"`
	Department string `gorm:"size:100" json:"department"`
	Pos        string `gorm:"size:50" json:"position"`
}

func (PositionHistory) TableName() string {
	return "cadm_position_histories"
}

func EditPositionHistory(data map[string]interface{}) error {
	// Get CadreID from the input data
	cadreID, ok := data["CadreID"].(string)
	if !ok || cadreID == "" {
		return errors.New("CadreID is required and must be a string")
	}

	// Find the corresponding record in PositionHistory_mod
	var modRecord PositionHistory_mod
	if err := db.Where("user_id = ?", cadreID).First(&modRecord).Error; err != nil {
		return err
	}

	// Create the new PositionHistory record
	newRecord := PositionHistory{
		CadreID:      modRecord.CadreID,
		Name:         modRecord.Name,
		PhoneNumber:  modRecord.PhoneNumber,
		Email:        modRecord.Email,
		Department:   modRecord.Department,
		Category:     modRecord.Category,
		Office:       modRecord.Office,
		AcademicYear: modRecord.AcademicYear,
		Positions:    modRecord.Positions,
	}

	// Add any additional fields from the input data
	if name, ok := data["Name"].(string); ok && name != "" {
		newRecord.Name = name
	}
	if phone, ok := data["PhoneNumber"].(string); ok {
		newRecord.PhoneNumber = phone
	}
	if email, ok := data["Email"].(string); ok {
		newRecord.Email = email
	}
	if dept, ok := data["Department"].(string); ok && dept != "" {
		newRecord.Department = dept
	}
	if cat, ok := data["Category"].(string); ok && cat != "" {
		newRecord.Category = cat
	}
	if office, ok := data["Office"].(string); ok && office != "" {
		newRecord.Office = office
	}
	if year, ok := data["AcademicYear"].(string); ok && year != "" {
		newRecord.AcademicYear = year
	}
	if pos, ok := data["Positions"].(string); ok {
		newRecord.Positions = pos
	}

	// Save the new record to the database
	if err := db.Create(&newRecord).Error; err != nil {
		return err
	}

	return nil
}

func GetPositionHistory(cadreID string) ([]PositionHistory, error) {
	if cadreID == "" {
		return nil, errors.New("干部ID不能为空")
	}

	var histories []PositionHistory
	err := db.Model(&PositionHistory{}).Preload("Cadre").Where(" id = ?", cadreID).Order("is_current DESC, start_date DESC").Find(&histories).Error

	switch {
	case err == nil:
		return histories, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return []PositionHistory{}, nil
	default:
		return nil, fmt.Errorf("查询职位历史失败: %w", err)
	}
}

func ComfirmPositionhistory(cadreID string) error {
	var exp_mod Posexp_mod
	var pos_mod PositionHistory_mod

	result_exp := db.Where("user_id = ? AND is_audited = false", cadreID).First(&exp_mod)
	expFound := result_exp.Error == nil

	result_pos := db.Where("user_id = ? AND is_audited = false", cadreID).First(&pos_mod)
	posFound := result_pos.Error == nil

	if !expFound && !posFound {
		return fmt.Errorf("未找到任何待审核信息")
	}

	positionhistory := PositionHistory{
		CadreID:      pos_mod.CadreID,      // 对应 user_id
		Name:         pos_mod.Name,         // 姓名
		PhoneNumber:  pos_mod.PhoneNumber,  // 电话号码
		Email:        pos_mod.Email,        // 电子邮件
		Department:   pos_mod.Department,   // 院系
		Category:     pos_mod.Category,     // 类别
		Office:       pos_mod.Office,       // 任职部门
		AcademicYear: pos_mod.AcademicYear, // 任职年度
		Positions:    pos_mod.Positions,    // 职位名称
		Year:         pos_mod.Year,         // 申请年份
		Month:        pos_mod.Month,        // 申请月份
		Day:          pos_mod.Day,          // 申请日
	}

	pos_exp := Posexp{
		CadreID:    exp_mod.CadreID,
		Posyear:    exp_mod.Posyear,
		Department: exp_mod.Department,
		Pos:        exp_mod.Pos,
	}

	if pos_exp.CadreID != "" {
		if err := db.Create(&pos_exp).Error; err != nil {
			return err
		}

		if err := db.Model(&Posexp_mod{}).
			Where("user_id = ? AND is_audited = false", cadreID).
			Update("is_audited", true).Error; err != nil {
			return err
		}
	}

	if positionhistory.CadreID != "" {
		if err := db.Create(&positionhistory).Error; err != nil {
			return err
		}

		if err := db.Model(&PositionHistory_mod{}).
			Where("user_id = ? AND is_audited = false", cadreID).
			Update("is_audited", true).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetPositionHistories(pageNum int, pageSize int, maps interface{}) ([]PositionHistory, error) {
	var (
		positionHistories []PositionHistory
		err               error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&positionHistories).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&positionHistories).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return positionHistories, nil
}

// GetPositionHistoryTotal counts the total number of position histories based on the constraint
func GetPositionHistoryTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&PositionHistory{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetPosexps(pageNum int, pageSize int, maps interface{}) ([]Posexp, error) {
	var (
		posexps []Posexp
		err     error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&posexps).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&posexps).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return posexps, nil
}

// GetPosexpTotal counts the total number of posexps based on the constraint
func GetPosexpTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Posexp{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func DeletePosexpByID(id int) error {
	if id <= 0 {
		return errors.New("无效的岗位经历记录 ID")
	}
	result := db.Where("id = ?", id).Delete(&Posexp{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到匹配的岗位经历记录")
	}
	return nil
}
