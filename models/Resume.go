package models

import (
	"fmt"
)

type ResumeEntry struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID      string `gorm:"not null" json:"user_id"`
	StartDate    string `json:"start_date"`           // 格式：2007.09 或 2019.12
	EndDate      string `json:"end_date"`             // 格式：2011.07 或 "至今"
	Organization string `json:"organization"`         // 工作单位或学校
	Department   string `json:"department,omitempty"` // 学院/部门，可选
	Position     string `json:"position,omitempty"`   // 职务/身份，可选
}

func (ResumeEntry) TableName() string {
	return "cadm_resume_entries"
}

func DeleteResumeEntryByID(id int) error {
	if err := db.Where("id = ?", id).Delete(ResumeEntry{}).Error; err != nil {
		return err
	}
	return nil
}

func ComfirmResume(id int) error {
	var mod ResumeEntry_modifications

	// 查询待审核的履历修改记录
	result := db.Where("id = ?", id).First(&mod)
	if result.Error != nil {
		return fmt.Errorf("未找到待审核的履历信息: %v", result.Error)
	}

	// 创建正式的履历条目
	resumeEntry := ResumeEntry{
		CadreID:      mod.CadreID,
		StartDate:    mod.StartDate,
		EndDate:      mod.EndDate,
		Organization: mod.Organization,
		Department:   mod.Department,
		Position:     mod.Position,
	}

	// 插入到正式的履历表中
	if err := db.Create(&resumeEntry).Error; err != nil {
		return fmt.Errorf("插入履历信息失败: %v", err)
	}

	// 更新审核状态
	mod.Audited = true
	if err := db.Save(&mod).Error; err != nil {
		return fmt.Errorf("更新履历审核状态失败: %v", err)
	}

	return nil
}
