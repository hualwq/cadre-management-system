package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Cadre struct {
	ID                        string `gorm:"primaryKey;type:varchar(50);column:user_id" json:"user_id"`
	PhotoUrl                  string `gorm:"type:varchar(100); column:photourl" json:"photourl"`
	Name                      string `gorm:"type:varchar(50); column:name" json:"name"`
	Gender                    string `gorm:"type:ENUM('男','女'); column:gender" json:"gender"`
	BirthDate                 string `gorm:"type:date; column:birth_date" json:"birth_date"`
	Age                       uint8  `gorm:"type:tinyint unsigned;column:age" json:"age"`
	EthnicGroup               string `gorm:"type:varchar(20); column:ethnic_group" json:"ethnic_group"`
	NativePlace               string `gorm:"type:varchar(100); column:native_place" json:"native_place"`
	BirthPlace                string `gorm:"type:varchar(100);column:birth_place" json:"birth_place"`
	PoliticalStatus           string `gorm:"type:ENUM('中共党员','中共预备党员','共青团员');column:political_status" json:"political_status"`
	WorkStartDate             string `gorm:"type:date; column:work_start_date" json:"work_start_date"`
	HealthStatus              string `gorm:"type:varchar(20);column:health_status" json:"health_status"`
	ProfessionalTitle         string `gorm:"type:varchar(100);column:professional_title" json:"professional_title"`
	Specialty                 string `gorm:"type:varchar(200);column:specialty" json:"specialty"`
	Phone                     string `gorm:"type:varchar(20); column:phone" json:"phone"`
	CurrentPosition           string `gorm:"type:varchar(200); column:current_position" json:"current_position"`
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
	IsAudited                 int    `gorm:"default:0;column:is_audited" json:"is_audited"`
}

func (Cadre) TableName() string {
	return "cadm_cadreinfo"
}

func GetCadre(cadreID string) (*Cadre, error) {
	if cadreID == "" {
		return nil, fmt.Errorf("干部ID不能为空")
	}

	var cadreInfo Cadre

	err := db.
		Select("*"). // 或者指定具体字段，如："user_id, name, department"
		Where("user_id = ?", cadreID).
		First(&cadreInfo).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("干部信息不存在，干部ID: %s", cadreID)
			return nil, fmt.Errorf("干部信息不存在")
		}

		log.Printf("查询干部信息失败，干部ID: %s，错误: %v", cadreID, err)
		return nil, fmt.Errorf("查询干部信息失败: %v", err)
	}

	return &cadreInfo, nil
}

func (c *Cadre) CalculateAge() error {
	if c.BirthDate == "" {
		return fmt.Errorf("出生日期为必选项")
	}

	// 分割字符串为年和月
	parts := strings.Split(c.BirthDate, ".")
	if len(parts) != 2 {
		return fmt.Errorf("invalid birth date format, expected 'YYYY.M'")
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid year in birth date: %v", err)
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		return fmt.Errorf("invalid month in birth date: %v", err)
	}

	now := time.Now()
	age := now.Year() - year

	// 如果今年生日还没到，年龄减1
	if int(now.Month()) < month || (int(now.Month()) == month && now.Day() < 1) {
		age--
	}

	c.Age = uint8(age)
	return nil
}

func AddCadre(data map[string]interface{}) error {
	cadreInfo := Cadre{
		ID:                        data["user_id"].(string),
		Name:                      data["name"].(string),
		Gender:                    data["gender"].(string),
		BirthDate:                 data["birth_date"].(string),
		EthnicGroup:               data["ethnic_group"].(string),
		NativePlace:               data["native_place"].(string),
		BirthPlace:                data["birth_place"].(string),
		PoliticalStatus:           data["political_status"].(string),
		WorkStartDate:             data["work_start_date"].(string),
		HealthStatus:              data["health_status"].(string),
		ProfessionalTitle:         data["professional_title"].(string),
		Specialty:                 data["specialty"].(string),
		Phone:                     data["phone"].(string),
		CurrentPosition:           data["current_position"].(string),
		AwardsAndPunishments:      data["awards_and_punishments"].(string),
		AnnualAssessment:          data["annual_assessment"].(string),
		Email:                     data["email"].(string),
		FilledBy:                  data["filled_by"].(string),
		FullTimeEducationDegree:   data["full_time_education_degree"].(string),
		FullTimeEducationSchool:   data["full_time_education_school"].(string),
		OnTheJobEducationDegree:   data["on_the_job_education_degree"].(string),
		OnTheJobEducationSchool:   data["on_the_job_education_school"].(string),
		ReportingUnit:             data["reporting_unit"].(string),
		ApprovalAuthority:         data["approval_authority"].(string),
		AdministrativeAppointment: data["administrative_appointment"].(string),
	}

	if err := cadreInfo.CalculateAge(); err != nil {
		return err
	}

	if err := db.Create(&cadreInfo).Error; err != nil {
		return err
	}

	return nil
}

func ExistCadreInfoByID(id string) (bool, error) {
	var count int64
	err := db.Model(&Cadre{}).Where("user_id = ? and is_audited = ?", id, 0).Count(&count).Error
	return count > 0, err
}

func EditCadreInfoByID(id string, data map[string]interface{}) error {
	return db.Model(&Cadre{}).Where("user_id = ?", id).Updates(data).Error
}

func DeleteCadreByID(id string) error {

	if err := db.Where("user_id = ?", id).Delete(&Familymember{}).Error; err != nil {
		return err
	}

	if err := db.Where("user_id = ?", id).Delete(&ResumeEntry{}).Error; err != nil {
		return err
	}

	if err := db.Where("user_id = ?", id).Delete(&Positionhistory{}).Error; err != nil {
		return err
	}

	if err := db.Where("user_id = ?", id).Delete(&Cadre{}).Error; err != nil {
		return err
	}

	return nil
}

func GetCadreInfoModByPage(pageNum int, pageSize int, maps interface{}) ([]Cadre, error) {
	var (
		cadreInfos []Cadre
		err        error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&cadreInfos).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&cadreInfos).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return cadreInfos, nil
}

func GetCadreInfoModTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Cadre{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// models/Cadres.go 或其他合适的模型文件
func UpdateCadrePhotoURL(cadreID string, photoURL string) error {
	if cadreID == "" {
		return fmt.Errorf("干部ID不能为空")
	}

	err := db.Model(&Cadre{}).Where("user_id = ?", cadreID).Update("photourl", photoURL).Error
	if err != nil {
		log.Printf("更新干部图片URL失败，干部ID: %s，错误: %v", cadreID, err)
		return fmt.Errorf("更新干部图片URL失败: %v", err)
	}

	return nil
}

func UpdateCadreInfoModPhotoURL(cadreid, photoURL string) error {
	var mod Cadre
	result := db.Where("user_id = ?", cadreid).First(&mod)
	if result.Error != nil {
		return fmt.Errorf("未找到待更新的干部信息修改记录: %v", result.Error)
	}

	mod.PhotoUrl = photoURL
	if err := db.Save(&mod).Error; err != nil {
		return fmt.Errorf("更新干部信息修改记录的 photourl 字段失败: %v", err)
	}

	return nil
}

func ComfirmCadre(cadreid string) error {
	result := db.Model(&Cadre{}).
		Where("user_id = ? and is_audited = 0", cadreid).
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
