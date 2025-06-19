package models

import (
	"gorm.io/gorm"
)

type ResumeEntry_modifications struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID      string `gorm:"column:user_id;not null" json:"user_id"`
	StartDate    string `json:"start_date"`           // 格式：2007.09 或 2019.12
	EndDate      string `json:"end_date"`             // 格式：2011.07 或 "至今"
	Organization string `json:"organization"`         // 工作单位或学校
	Department   string `json:"department,omitempty"` // 学院/部门，可选
	Position     string `json:"position,omitempty"`   // 职务/身份，可选
	Audited      int    `gorm:"default:0;column:audit_status"`
}

func (ResumeEntry_modifications) TableName() string {
	return "cadm_resume_entries_mod"
}

func Add_resume_mod(data map[string]interface{}) error {
	resumeEntry := ResumeEntry_modifications{
		CadreID:      data["user_id"].(string),
		StartDate:    data["start_date"].(string),
		EndDate:      data["end_date"].(string),
		Organization: data["organization"].(string),
	}

	// 可选字段，检查是否存在
	if department, ok := data["department"]; ok {
		resumeEntry.Department = department.(string)
	}
	if position, ok := data["position"]; ok {
		resumeEntry.Position = position.(string)
	}

	if err := db.Create(&resumeEntry).Error; err != nil {
		return err
	}

	return nil
}

func ExistResumeEntryModificationByID(id int) (bool, error) {
	var entry ResumeEntry_modifications
	err := db.Select("id").Where("id = ?  and is_audited = ?", id, 0).First(&entry).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if entry.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetResumeEntryModificationByID 根据 ID 获取单个履历条目修改记录
func GetResumeEntryModificationByID(id int) (*ResumeEntry_modifications, error) {
	var entry ResumeEntry_modifications
	err := db.Where("id = ?", id).First(&entry).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &entry, nil
}

// GetResumeEntryModificationsByCadreID 根据 CadreID 获取履历条目修改记录列表
func GetResumeEntryModificationsByCadreID(cadreID string) ([]ResumeEntry_modifications, error) {
	var entries []ResumeEntry_modifications
	err := db.Where("user_id = ?", cadreID).Find(&entries).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return entries, nil
}

// DeleteResumeEntryModificationByID 根据 ID 删除单个履历条目修改记录
func DeleteResumeEntryModificationByID(id int) error {
	if err := db.Where("id = ?", id).Delete(ResumeEntry_modifications{}).Error; err != nil {
		return err
	}

	return nil
}

func EditResumeEntryModification(id int, data map[string]interface{}) error {
	var resumeEntry ResumeEntry_modifications
	result := db.Model(&resumeEntry).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
