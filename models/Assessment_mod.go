package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Assessment 干部考核模型
type Assessment_mod struct {
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
	Audited     bool   `gorm:"default:false;column:is_audited"`
}

func (Assessment_mod) TableName() string {
	return "cadm_assessments_mod"
}

func GetAssessmentsMod(pageNum int, pageSize int, maps interface{}) ([]Assessment_mod, error) {
	var (
		assessments []Assessment_mod
		err         error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&assessments).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&assessments).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return assessments, nil
}

func GetAssessmentModTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Assessment_mod{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddAssessment_mod(data map[string]interface{}) error {
	// 1. 从 data 中解析必需字段
	userID, ok := data["user_id"].(string)
	if !ok || userID == "" {
		return errors.New("invalid or missing user_id")
	}

	// 2. 查询 cadre_info 表获取干部基本信息
	var cadre CadreInfo
	if err := db.Where("user_id = ?", userID).First(&cadre).Error; err != nil {
		return fmt.Errorf("failed to find cadre info: %v", err)
	}

	// 4. 处理年份（优先使用传入的年份，否则使用当前年份）
	var year int
	if inputYear, ok := data["year"].(float64); ok {
		year = int(inputYear)
	} else {
		year = time.Now().Year()
	}

	// 5. 检查必需字段
	requiredFields := []string{"department", "category", "work_summary"}
	for _, field := range requiredFields {
		if _, ok := data[field].(string); !ok {
			return fmt.Errorf("missing or invalid required field: %s", field)
		}
	}

	// 6. 创建 Assessment 对象
	assessment := Assessment_mod{
		Name:        cadre.Name,
		CadreID:     userID,
		Phone:       cadre.Phone,
		Email:       cadre.Email,
		Department:  data["department"].(string),
		Category:    data["category"].(string),
		AssessDept:  data["assess_dept"].(string),
		Year:        year, // 使用处理后的年份
		WorkSummary: data["work_summary"].(string),
		Grade:       "待评定", // 默认值
	}

	// 7. 检查是否已存在相同记录
	var existing Assessment_mod
	err := db.Where("user_id = ? AND year = ?", userID, year).First(&existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing assessment: %v", err)
	}

	// 8. 执行创建或更新
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&assessment).Error; err != nil {
			return fmt.Errorf("failed to create assessment: %v", err)
		}
	} else {
		assessment.ID = existing.ID
		if err := db.Save(&assessment).Error; err != nil {
			return fmt.Errorf("failed to update assessment: %v", err)
		}
	}

	return nil
}

func ExistAssesementByID(id int) (bool, error) {
	var assesement Assessment_mod
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&assesement).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if assesement.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetAssesement Get a single assessment based on ID
func GetAssesement(id int) (*Assessment_mod, error) {
	var assesement Assessment_mod
	err := db.Where("id = ?", id).First(&assesement).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &assesement, nil
}

func DeleteAssessmentModByID(id int) error {
	if id <= 0 {
		return errors.New("无效的考核记录 ID")
	}
	result := db.Where("id = ?", id).Delete(&Assessment_mod{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到匹配的考核记录")
	}
	return nil
}

func ExistAssessmentModByID(id int) (bool, error) {
	var count int64
	err := db.Model(&Assessment_mod{}).Where("id = ? and is_audited = false", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func EditAssessmentModByID(id int, data map[string]interface{}) error {
	return db.Model(&Assessment_mod{}).Where("id = ?", id).Updates(data).Error
}


