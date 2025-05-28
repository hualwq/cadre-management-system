package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type FamilyMember struct {
	ID              int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID         string `gorm:"not null" json:"user_id"`
	Relation        string `gorm:"type:varchar(20);not null" json:"relation"`
	Name            string `gorm:"type:varchar(50);not null" json:"name"`
	BirthDate       string `gorm:"type:date" json:"birth_date,omitempty"`
	PoliticalStatus string `gorm:"type:varchar(50)" json:"political_status,omitempty"`
	WorkUnit        string `gorm:"type:varchar(200)" json:"work_unit,omitempty"`
}

func (FamilyMember) TableName() string {
	return "cadm_family_members"
}

func GetFamilyMembers(CadreID string) ([]FamilyMember, error) {
	var familyMembers []FamilyMember

	err := db.Where(" id = ?", CadreID).Find(&familyMembers).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return familyMembers, nil
}

func AddFamilyMembers(data map[string]interface{}) error {
	var members []FamilyMember

	if fm, ok := data["family_members"]; ok {
		// Handle multiple members
		membersData := fm.([]map[string]interface{})
		for _, memberData := range membersData {
			member, err := createFamilyMemberFromMap(memberData)
			if err != nil {
				return err
			}
			members = append(members, member)
		}
	} else {
		// Handle single member
		member, err := createFamilyMemberFromMap(data)
		if err != nil {
			return err
		}
		members = append(members, member)
	}

	if err := db.Create(&members).Error; err != nil {
		return err
	}

	return nil
}

func createFamilyMemberFromMap(data map[string]interface{}) (FamilyMember, error) {

	return FamilyMember{
		CadreID:         data["user_id"].(string),
		Relation:        data["relation"].(string),
		Name:            data["name"].(string),
		BirthDate:       data["birth_date"].(string),
		PoliticalStatus: data["political_status"].(string),
		WorkUnit:        data["work_unit"].(string),
	}, nil
}

func DeleteFamilyMember(ID string) (bool, error) {
	if ID == "" {
		return false, errors.New("家庭成员ID不能为空")
	}

	result := db.Where("id = ?", ID).Delete(&FamilyMember{})

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("指定的家庭成员不存在")
		}
		return false, fmt.Errorf("删除家庭成员失败: %v", err)
	}

	if result.RowsAffected == 0 {
		return false, errors.New("未找到匹配的家庭成员记录")
	}

	return true, nil
}
