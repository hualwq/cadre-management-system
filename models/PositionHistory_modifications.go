package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PositionHistory_mod struct {
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
	Audited      bool   `gorm:"default:false;column:is_audited"`
}

type PositionHistoryBrief_mod struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Posexp_mod struct {
	CadreID    string `gorm:"size:50;column:user_id" json:"user_id"` // ✅ 用空格分隔两个标签
	Posyear    string `gorm:"size:20" json:"year"`
	Department string `gorm:"size:100" json:"department"`
	Pos        string `gorm:"size:50" json:"position"`
	Audited    bool   `gorm:"default:false;column:is_audited"`
}

func (PositionHistory_mod) TableName() string {
	return "cadm_position_histories_mod"
}

func (Posexp_mod) TableName() string {
	return "cadm_posexp_mod"
}

func GetPositionHistory_mod(cadreID string) (*PositionHistory_mod, error) {
	// 查找 PositionHistory 表中该干部的任职记录
	var position PositionHistory_mod
	if err := db.Where("id = ?", cadreID).First(&position).Error; err != nil {
		return nil, fmt.Errorf("failed to find position history: %v", err)
	}

	// 查找 cadre_info 表获取姓名、电话、邮箱等信息
	var cadre CadreInfo
	if err := db.Where("id = ?", cadreID).First(&cadre).Error; err != nil {
		return nil, fmt.Errorf("failed to find cadre info: %v", err)
	}

	// 补充信息
	position.Name = cadre.Name
	position.CadreID = cadreID
	position.PhoneNumber = cadre.Phone
	position.Email = cadre.Email

	return &position, nil
}

func AddPositionHistory_mod(data map[string]interface{}) error {
	// 从 data 中解析直接提供的字段
	cadreID, ok := data["user_id"].(string)
	if !ok {
		return errors.New("invalid or missing cadre ID")
	}

	// 查询 cadre_info 表获取其他信息
	var cadre CadreInfo
	if err := db.Where("user_id = ?", cadreID).First(&cadre).Error; err != nil {
		return fmt.Errorf("请先编辑基本信息或等管理员审核信息")
	}

	// 创建 PositionHistory_mod 对象
	positionHistory := PositionHistory_mod{
		CadreID:      cadreID,
		Name:         cadre.Name,
		PhoneNumber:  cadre.Phone,
		Email:        cadre.Email,
		Department:   data["department"].(string),
		Category:     data["category"].(string),
		Office:       data["office"].(string),
		AcademicYear: data["academic_year"].(string),
		Year:         data["applied_at_year"].(uint),
		Month:        data["applied_at_month"].(uint),
		Day:          data["applied_at_day"].(uint),
	}

	// 检查是否存在相同 CadreID 和 AcademicYear 的记录
	var existing PositionHistory_mod
	err := db.Where("user_id = ? AND academic_year = ?", cadreID, positionHistory.AcademicYear).First(&existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing position history: %v", err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在旧记录，执行插入
		if err := db.Create(&positionHistory).Error; err != nil {
			return fmt.Errorf("failed to create position history: %v", err)
		}
	} else {
		// 存在旧记录，更新
		if err := db.Save(&positionHistory).Error; err != nil {
			return fmt.Errorf("failed to update position history: %v", err)
		}
	}

	return nil
}

func GetPositionhistoryList_mod_page(page, pageSize int) ([]PositionHistoryBrief_mod, int64, error) {
	var assessments []PositionHistory_mod
	var count int64

	// 获取总数
	if err := db.Model(&PositionHistory_mod{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&assessments).Error; err != nil {
		return nil, 0, err
	}

	// 转换结果
	var result []PositionHistoryBrief_mod
	for _, a := range assessments {
		result = append(result, PositionHistoryBrief_mod{
			Name:       a.Name,
			Department: a.Department,
			// 可以添加更多字段
		})
	}

	return result, count, nil
}

func Addyearpositon(data map[string]interface{}) error {
	posexp := Posexp_mod{
		CadreID:    data["user_id"].(string),
		Posyear:    data["year"].(string),
		Department: data["department"].(string),
		Pos:        data["position"].(string),
	}

	// 3. 数据库操作
	if err := db.Create(&posexp).Error; err != nil {
		return fmt.Errorf("database create failed: %v", err)
	}

	return nil
}

func GetPositionHistoryModByID(id int) (*PositionHistory_mod, error) {
	var positionHistoryMod PositionHistory_mod
	if err := db.Where("id = ?", id).First(&positionHistoryMod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &positionHistoryMod, nil
}

// GetPositionHistories gets a list of position histories based on paging and constraints
func GetPositionHistoriesMod(pageNum int, pageSize int, maps interface{}) ([]PositionHistory_mod, error) {
	var (
		positionHistories []PositionHistory_mod
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
func GetPositionHistoryModTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&PositionHistory_mod{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}


// DeletePositionHistoryByID delete a single position history
func DeletePositionHistoryByID(id int) error {
	if err := db.Where("id = ?", id).Delete(PositionHistory_mod{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistPositionHistoryByID(id int) (bool, error) {
	var positionHistory PositionHistory_mod
	err := db.Select("id").Where("id = ?", id, 0).First(&positionHistory).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if positionHistory.ID > 0 {
		return true, nil
	}

	return false, nil
}
