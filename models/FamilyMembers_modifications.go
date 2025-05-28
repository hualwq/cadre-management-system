package models

import (
	"fmt"

	"gorm.io/gorm"
)

type FamilyMember_modifications struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID         string `gorm:"not null;column:user_id" json:"user_id"`
	Relation        string `gorm:"type:varchar(20);not null" json:"relation"`
	Name            string `gorm:"type:varchar(50);not null" json:"name"`
	BirthDate       string `gorm:"type:date" json:"birth_date,omitempty"`
	PoliticalStatus string `gorm:"type:varchar(50)" json:"political_status,omitempty"`
	WorkUnit        string `gorm:"type:varchar(200)" json:"work_unit,omitempty"`
	Audited         bool   `gorm:"default:false;column:is_audited"`
}

func (FamilyMember_modifications) TableName() string {
	return "cadm_family_members_mod"
}

func Add_familymember_mod(data map[string]interface{}) error {
	// 检查必填字段是否存在
	requiredFields := []string{"user_id", "relation", "name"}
	for _, field := range requiredFields {
		if _, ok := data[field]; !ok {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	familyMember := FamilyMember_modifications{
		CadreID:  data["user_id"].(string),
		Relation: data["relation"].(string),
		Name:     data["name"].(string),
	}

	// 可选字段，检查是否存在
	if birthDate, ok := data["birth_date"]; ok && birthDate != nil {
		familyMember.BirthDate = birthDate.(string)
	}
	if politicalStatus, ok := data["political_status"]; ok && politicalStatus != nil {
		familyMember.PoliticalStatus = politicalStatus.(string)
	}
	if workUnit, ok := data["work_unit"]; ok && workUnit != nil {
		familyMember.WorkUnit = workUnit.(string)
	}

	if err := db.Create(&familyMember).Error; err != nil {
		return fmt.Errorf("failed to create family member record: %v", err)
	}

	return nil
}

func ExistByID(id int) (bool, error) {
	var count int64

	// 查询是否存在匹配的记录
	err := db.Model(&FamilyMember_modifications{}).
		Where("id = ?", id).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	// count > 0 表示记录存在
	return count > 0, nil
}

// ExistFamilyMemberModificationByID checks if a family member modification exists based on ID
func ExistFamilyMemberModificationByID(id int) (bool, error) {
	var member FamilyMember_modifications
	err := db.Select("id").Where("id = ?", id).First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if member.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetFamilyMemberModificationByID Get a single family member modification based on ID
func GetFamilyMemberModificationByID(id int) (*FamilyMember_modifications, error) {
	var member FamilyMember_modifications
	err := db.Where("id = ?", id).First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &member, nil
}

// DeleteFamilyMemberModificationByID delete a single family member modification
func DeleteFamilyMemberModificationByID(id int) error {
	if err := db.Where("id = ?", id).Delete(FamilyMember_modifications{}).Error; err != nil {
		return err
	}

	return nil
}

// GetFamilyMemberModificationsByCadreID 根据 CadreID 获取家庭成员修改记录
func GetFamilyMemberModificationsByCadreID(cadreID string) ([]FamilyMember_modifications, error) {
	var members []FamilyMember_modifications
	err := db.Where("user_id = ?", cadreID).Find(&members).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return members, nil
}
