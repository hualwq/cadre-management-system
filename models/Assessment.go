package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Assessment 干部考核模型
type Assessment struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string `gorm:"type:varchar(50);not null;" json:"name"`
	CadreID      string `gorm:"type:varchar(20);column:user_id" json:"user_id"`
	Phone        string `gorm:"type:varchar(20);" json:"phone"`
	Email        string `gorm:"type:varchar(100);" json:"email"`
	Department   string `gorm:"type:varchar(100);not null;" json:"department"`
	Category     string `gorm:"type:varchar(20);not null;" json:"category"`
	AssessDept   string `gorm:"type:varchar(100);not null;" json:"assess_dept"`
	Year         int    `gorm:"type:int;not null;" json:"year"`
	WorkSummary  string `gorm:"type:text;not null;" json:"work_summary"`
	Grade        string `gorm:"type:varchar(10);not null;" json:"grade"`
	IsAudited    int    `gorm:"default:0;column:is_audited" json:"is_audited"`
	DepartmentID int    `gorm:"type:int;not null;" json:"department_id"`
}

func (Assessment) TableName() string {
	return "cadm_assessments"
}

func GetAssessmentsMod(pageNum int, pageSize int, maps interface{}) ([]Assessment, error) {
	var (
		assessments []Assessment
		err         error
	)

	fmt.Println("maps", maps)

	query := db.Where(maps)

	if pageSize > 0 && pageNum > 0 {
		query = query.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	err = query.Find(&assessments).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return assessments, nil
}

func GetAssessmentModTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Assessment{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddAssessment(data map[string]interface{}) error {
	// 1. 从 data 中解析必需字段
	userID, ok := data["user_id"].(string)
	if !ok || userID == "" {
		return errors.New("invalid or missing user_id")
	}

	// 2. 查询 cadre_info 表获取干部基本信息
	var cadre Cadre
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
	assessment := Assessment{
		Name:         cadre.Name,
		CadreID:      userID,
		Phone:        cadre.Phone,
		Email:        cadre.Email,
		Department:   data["department"].(string),
		Category:     data["category"].(string),
		AssessDept:   data["assess_dept"].(string),
		Year:         year, // 使用处理后的年份
		WorkSummary:  data["work_summary"].(string),
		DepartmentID: data["department_id"].(int),
		Grade:        "待评定", // 默认值
	}

	// 7. 检查是否已存在相同记录
	var existing Assessment
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
	var assesement Assessment
	err := db.Select("id").Where("id = ?", id).First(&assesement).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if assesement.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetAssesement Get a single assessment based on ID
func GetAssesement(id int) (*Assessment, error) {
	var assesement Assessment
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
	result := db.Where("id = ?", id).Delete(&Assessment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到匹配的考核记录")
	}
	return nil
}

func EditAssessmentModByID(id int, data map[string]interface{}) error {
	return db.Model(&Assessment{}).Where("id = ?", id).Updates(data).Error
}

func ComfirmAssessment(id int, grade string) error {
	// Update the assessment's grade and set audited status to 1
	result := db.Model(&Assessment{}).
		Where("id = ? and is_audited = 0", id).
		Updates(map[string]interface{}{
			"grade":      grade,
			"is_audited": 1, // Using column name from struct tag
		})

	fmt.Println("ComfirmAssessment", id, grade)

	if result.Error != nil {
		return result.Error
	}

	// Check if any record was actually updated
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func DeleteAssessmentByID(id int) error {
	if err := db.Where("id = ?", id, 0).Delete(&Assessment{}).Error; err != nil {
		return err
	}
	return nil
}

func ExistAssessmentModByID(id int) (bool, error) {
	var count int64
	err := db.Model(&Assessment{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}
