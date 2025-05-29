package models

import (
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

func DeleteFamilyMemberByID(id int) error {
	if err := db.Where("id = ?", id).Delete(FamilyMember{}).Error; err != nil {
		return err
	}
	return nil
}

func Comfirmfamilymember(id int) error {
	var mod FamilyMember_modifications

	// 查询待审核的家庭成员修改记录
	result := db.Where("id = ?", id).First(&mod)
	if result.Error != nil {
		return fmt.Errorf("未找到待审核的家庭成员信息: %v", result.Error)
	}

	// 创建正式的家庭成员记录
	familyMember := FamilyMember{
		CadreID:         mod.CadreID,
		Relation:        mod.Relation,
		Name:            mod.Name,
		BirthDate:       mod.BirthDate,
		PoliticalStatus: mod.PoliticalStatus,
		WorkUnit:        mod.WorkUnit,
	}

	// 插入到正式的家庭成员表中
	if err := db.Create(&familyMember).Error; err != nil {
		return fmt.Errorf("插入家庭成员信息失败: %v", err)
	}

	// 更新审核状态
	mod.Audited = true
	if err := db.Save(&mod).Error; err != nil {
		return fmt.Errorf("更新家庭成员审核状态失败: %v", err)
	}

	return nil
}
