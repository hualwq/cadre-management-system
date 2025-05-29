package models

import (
	"errors"
	"fmt"
)

// Assessment 干部考核模型
type Assessment struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(50);not null;" json:"name"`
	CadreID     string `gorm:"type:varchar(20);column:user_id" json:"user_id"`
	Phone       string `gorm:"type:varchar(20);" json:"phone"`
	Email       string `gorm:"type:varchar(100);" json:"email"`
	Department  string `gorm:"type:varchar(100);not null;" json:"department"`
	Category    string `gorm:"type:varchar(20);not null;" json:"category"`
	AssessDept  string `gorm:"type:varchar(100);not null;" json:"assess_dept"`
	Year        int    `gorm:"type:int;not null;" json:"year"`
	WorkSummary string `gorm:"type:text;not null;" json:"work_summary"`
	Grade       string `gorm:"type:varchar(10);not null;" json:"grade"`
}

func (Assessment) TableName() string {
	return "cadm_assessments"
}

func ComfirmAssessment(id int, Grade string) error {
	var mod Assessment_mod

	// 查询 Cadre_Modification 中 user_id 为 cadreid 的记录
	result := db.Where("id = ?", id).First(&mod)
	if result.Error != nil {
		return fmt.Errorf("未找到待审核信息: %v", result.Error)
	}
	assessment := Assessment{
		CadreID:     mod.CadreID,
		Department:  mod.Department,
		Category:    mod.Category,
		AssessDept:  mod.AssessDept,
		Year:        mod.Year,
		WorkSummary: mod.WorkSummary,
		Grade:       Grade,
		Name:        mod.Name,
		Phone:       mod.Phone,
		Email:       mod.Email,
	}

	if err := db.Create(&assessment).Error; err != nil {
		return fmt.Errorf("插入 CadreInfo 失败: %v", err)
	}

	mod.Audited = true
	if err := db.Save(&mod).Error; err != nil {
		return fmt.Errorf("更新审核状态失败: %v", err)
	}

	return nil
}

func DeleteAssessmentByID(id int) error {
	if id <= 0 {
		return errors.New("无效的考核记录 ID")
	}
	result := db.Where("id = ?", id).Delete(&Assessment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到匹配的考核记录")
	}
	return nil
}
