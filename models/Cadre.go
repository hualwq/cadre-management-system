package models

import (
	"fmt"

	"gorm.io/gorm"
)

type CadreInfo struct {
	ID                        string `gorm:"primaryKey;type:varchar(50);column:user_id" json:"user_id"`
	Name                      string `gorm:"type:varchar(50);not null;column:name" json:"name"`
	Gender                    string `gorm:"type:ENUM('男','女');not null;column:gender" json:"gender"`
	BirthDate                 string `gorm:"type:date;not null;column:birth_date" json:"birth_date"`
	Age                       uint8  `gorm:"type:tinyint unsigned;column:age" json:"age"`
	EthnicGroup               string `gorm:"type:varchar(20);not null;column:ethnic_group" json:"ethnic_group"`
	NativePlace               string `gorm:"type:varchar(100);not null;column:native_place" json:"native_place"`
	BirthPlace                string `gorm:"type:varchar(100);column:birth_place" json:"birth_place"`
	PoliticalStatus           string `gorm:"type:ENUM('中共党员','中共预备党员','共青团员');column:political_status" json:"political_status"`
	WorkStartDate             string `gorm:"type:date;not null;column:work_start_date" json:"work_start_date"`
	HealthStatus              string `gorm:"type:varchar(20);column:health_status" json:"health_status"`
	ProfessionalTitle         string `gorm:"type:varchar(100);column:professional_title" json:"professional_title"`
	Specialty                 string `gorm:"type:varchar(200);column:specialty" json:"specialty"`
	Phone                     string `gorm:"type:varchar(20);not null;column:phone" json:"phone"`
	CurrentPosition           string `gorm:"type:varchar(200);not null;column:current_position" json:"current_position"`
	AwardsAndPunishments      string `gorm:"type:text;column:awards_and_punishments" json:"awards_and_punishments"`
	AnnualAssessment          string `gorm:"type:text;column:annual_assessment" json:"annual_assessment"`
	Email                     string `gorm:"type:varchar(50);column:email" json:"email"`
	FilledBy                  string `gorm:"type:varchar(50);column:filled_by" json:"filled_by"`
	FullTimeEducationDegree   string `gorm:"type:varchar(50);column:full_time_education_degree" json:"full_time_education_degree"`
	FullTimeEducationSchool   string `gorm:"type:varchar(200);column:full_time_education_school" json:"full_time_education_school"`
	OnTheJobEducationDegree   string `gorm:"type:varchar(50);column:on_the_job_education_degree" json:"on_the_job_education_degree"`
	OnTheJobEducationSchool   string `gorm:"type:varchar(200);column:on_the_job_education_school" json:"on_the_job_education_school"`
	ReportingUnit             string `gorm:"type:varchar(200);column:reporting_unit" json:"reporting_unit"`
	ApprovalAuthority         string `gorm:"type:text;column:approval_authority" json:"approval_authority"`
	AdministrativeAppointment string `gorm:"type:text;column:administrative_appointment" json:"administrative_appointment"`
}

func (CadreInfo) TableName() string {
	return "cadm_cadreinfo"
}

func GetCadreInfo(cadreID string) (*CadreInfo, error) {
	var cadreInfo CadreInfo
	err := db.Preload("CadreInfo").Where("user_id = ?", cadreID).First(&cadreInfo).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("干部信息不存在")
		}
		return nil, fmt.Errorf("查询干部信息失败: %v", err)
	}

	return &cadreInfo, nil
}

func AddCadreInfofromMod(cadreid string) error {
	var mod Cadre_Modification

	// 查询 Cadre_Modification 中 user_id 为 cadreid 的记录
	result := db.Where("user_id = ?", cadreid).First(&mod)
	if result.Error != nil {
		return fmt.Errorf("未找到待审核信息: %v", result.Error)
	}

	// 创建 CadreInfo 实体
	cadreInfo := CadreInfo{
		ID:                        mod.ID,
		Name:                      mod.Name,
		Gender:                    mod.Gender,
		BirthDate:                 mod.BirthDate,
		Age:                       mod.Age,
		EthnicGroup:               mod.EthnicGroup,
		NativePlace:               mod.NativePlace,
		BirthPlace:                mod.BirthPlace,
		PoliticalStatus:           mod.PoliticalStatus,
		WorkStartDate:             mod.WorkStartDate,
		HealthStatus:              mod.HealthStatus,
		ProfessionalTitle:         mod.ProfessionalTitle,
		Specialty:                 mod.Specialty,
		CurrentPosition:           mod.CurrentPosition,
		AwardsAndPunishments:      mod.AwardsAndPunishments,
		AnnualAssessment:          mod.AnnualAssessment,
		Email:                     mod.Email,
		FilledBy:                  mod.FilledBy,
		FullTimeEducationDegree:   mod.FullTimeEducationDegree,
		FullTimeEducationSchool:   mod.FullTimeEducationSchool,
		OnTheJobEducationDegree:   mod.OnTheJobEducationDegree,
		OnTheJobEducationSchool:   mod.OnTheJobEducationSchool,
		ReportingUnit:             mod.ReportingUnit,
		ApprovalAuthority:         mod.ApprovalAuthority,
		AdministrativeAppointment: mod.AdministrativeAppointment,
		Phone:                     mod.Phone,
	}

	// 插入到 CadreInfo 表中
	if err := db.Create(&cadreInfo).Error; err != nil {
		return fmt.Errorf("插入 CadreInfo 失败: %v", err)
	}

	mod.Audited = true
	if err := db.Save(&mod).Error; err != nil {
		return fmt.Errorf("更新审核状态失败: %v", err)
	}

	return nil
}

func DeleteCadreInfoByID(id string) error {
	if err := db.Where("user_id = ?", id).Delete(CadreInfo{}).Error; err != nil {
		return err
	}
	return nil
}
