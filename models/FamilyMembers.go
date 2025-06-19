package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Familymember struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID         string `gorm:"not null;column:user_id" json:"user_id"`
	Relation        string `gorm:"type:varchar(20);not null" json:"relation"`
	Name            string `gorm:"type:varchar(50);not null" json:"name"`
	BirthDate       string `gorm:"type:date" json:"birth_date,omitempty"`
	PoliticalStatus string `gorm:"type:varchar(50)" json:"political_status,omitempty"`
	WorkUnit        string `gorm:"type:varchar(200)" json:"work_unit,omitempty"`
	IsAudited       int    `gorm:"default:0;column:is_audited" json:"is_audited"`
}

func (Familymember) TableName() string {
	return "cadm_family_members"
}

func Addfamilymember(data map[string]interface{}) error {
	// 检查必填字段是否存在
	requiredFields := []string{"user_id", "relation", "name"}
	for _, field := range requiredFields {
		if _, ok := data[field]; !ok {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	familyMember := Familymember{
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
	err := db.Model(&Familymember{}).
		Where("id = ? and is_audited = ?", id, 0).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	// count > 0 表示记录存在
	return count > 0, nil
}

func EditFamilyMember(id int, data map[string]interface{}) error {
	var familyMember Familymember
	result := db.Model(&familyMember).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// ExistFamilyMemberModificationByID checks if a family member modification exists based on ID
func ExistFamilyMemberByID(id int) (bool, error) {
	var member Familymember
	err := db.Select("id").Where("id = ?  and is_audited = ?", id, 0).First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if member.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetFamilyMemberModificationByID Get a single family member modification based on ID
func GetFamilyMemberByID(id int) (*Familymember, error) {
	var member Familymember
	err := db.Where("id = ?", id).First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &member, nil
}

// DeleteFamilyMemberModificationByID delete a single family member modification
func DeleteFamilyMemberByID(id int) error {
	if err := db.Where("id = ?", id).Delete(Familymember{}).Error; err != nil {
		return err
	}

	return nil
}

// GetFamilyMemberModificationsByCadreID 根据 CadreID 获取家庭成员修改记录
func GetFamilyMembersByCadreID(cadreID string) ([]Familymember, error) {
	var members []Familymember
	err := db.Where("user_id = ?", cadreID).Find(&members).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return members, nil
}

func Comfirmfamilymember(id int) error {
	result := db.Model(&Familymember{}).
		Where("id = ? and is_aidited = 0", id).
		Updates(map[string]interface{}{
			"is_audited": 1, // Using column name from struct tag
		})

	if result.Error != nil {
		return result.Error
	}

	// Check if any record was actually updated
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
