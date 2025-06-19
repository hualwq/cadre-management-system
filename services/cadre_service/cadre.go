package Cadre_service

import (
	"cadre-management/models"
)

type Cadre struct {
	ID                        string `json:"user_id"`
	Name                      string `json:"name"`
	Gender                    string `json:"gender"`
	BirthDate                 string `json:"birth_date"`
	Age                       uint8  `json:"age"`
	EthnicGroup               string `json:"ethnic_group"`
	NativePlace               string `json:"native_place"`
	BirthPlace                string `json:"birth_place"`
	PoliticalStatus           string `json:"political_status"`
	WorkStartDate             string `json:"work_start_date"`
	HealthStatus              string `json:"health_status"`
	ProfessionalTitle         string `json:"professional_title"`
	Specialty                 string `json:"specialty"`
	Phone                     string `json:"phone"`
	CurrentPosition           string `json:"current_position"`
	AwardsAndPunishments      string `json:"awards_and_punishments"`
	AnnualAssessment          string `json:"annual_assessment"`
	Email                     string `json:"email"`
	FilledBy                  string `json:"filled_by"`
	FullTimeEducationDegree   string `json:"full_time_education_degree"`
	FullTimeEducationSchool   string `json:"full_time_education_school"`
	OnTheJobEducationDegree   string `json:"on_the_job_education_degree"`
	OnTheJobEducationSchool   string `json:"on_the_job_education_school"`
	ReportingUnit             string `json:"reporting_unit"`
	ApprovalAuthority         string `json:"approval_authority"`
	AdministrativeAppointment string `json:"administrative_appointment"`
}

func (c *Cadre) DeleteByID() error {
	return models.DeleteCadreByID(c.ID)
}

func (c *Cadre) GetCadreInfo() (*models.Cadre, error) {
	return models.GetCadre(c.ID)
}

func (c *Cadre) AddCadreInfo() error {
	Cinfo := map[string]interface{}{
		"user_id":                     c.ID,
		"name":                        c.Name,
		"gender":                      c.Gender,
		"birth_date":                  c.BirthDate,
		"ethnic_group":                c.EthnicGroup,
		"native_place":                c.NativePlace,
		"birth_place":                 c.BirthPlace,
		"political_status":            c.PoliticalStatus,
		"work_start_date":             c.WorkStartDate,
		"health_status":               c.HealthStatus,
		"professional_title":          c.ProfessionalTitle,
		"specialty":                   c.Specialty,
		"phone":                       c.Phone,
		"current_position":            c.CurrentPosition,
		"awards_and_punishments":      c.AwardsAndPunishments,
		"annual_assessment":           c.AnnualAssessment,
		"email":                       c.Email,
		"filled_by":                   c.FilledBy,
		"full_time_education_degree":  c.FullTimeEducationDegree,
		"full_time_education_school":  c.FullTimeEducationSchool,
		"on_the_job_education_degree": c.OnTheJobEducationDegree,
		"on_the_job_education_school": c.OnTheJobEducationSchool,
		"reporting_unit":              c.ReportingUnit,
		"approval_authority":          c.ApprovalAuthority,
		"administrative_appointment":  c.AdministrativeAppointment,
	}

	if err := models.AddCadre(Cinfo); err != nil {
		return err
	}
	return nil
}

func (c *Cadre) ExistByID() (bool, error) {
	return models.ExistCadreInfoByID(c.ID)
}

func (c *Cadre) Edit() error {
	// 构造基本字段
	data := map[string]interface{}{
		"name":                        c.Name,
		"gender":                      c.Gender,
		"birth_date":                  c.BirthDate,
		"age":                         c.Age,
		"ethnic_group":                c.EthnicGroup,
		"native_place":                c.NativePlace,
		"birth_place":                 c.BirthPlace,
		"political_status":            c.PoliticalStatus,
		"work_start_date":             c.WorkStartDate,
		"health_status":               c.HealthStatus,
		"professional_title":          c.ProfessionalTitle,
		"specialty":                   c.Specialty,
		"phone":                       c.Phone,
		"current_position":            c.CurrentPosition,
		"awards_and_punishments":      c.AwardsAndPunishments,
		"annual_assessment":           c.AnnualAssessment,
		"email":                       c.Email,
		"filled_by":                   c.FilledBy,
		"full_time_education_degree":  c.FullTimeEducationDegree,
		"full_time_education_school":  c.FullTimeEducationSchool,
		"on_the_job_education_degree": c.OnTheJobEducationDegree,
		"on_the_job_education_school": c.OnTheJobEducationSchool,
		"reporting_unit":              c.ReportingUnit,
		"approval_authority":          c.ApprovalAuthority,
		"administrative_appointment":  c.AdministrativeAppointment,
	}

	// 调用 model 层基础信息更新
	if err := models.EditCadreInfoByID(c.ID, data); err != nil {
		return err
	}

	return nil
}

func (c *Cadre) ComfirmCadre() error {
	cadreid := c.ID
	return models.ComfirmCadre(cadreid)
}
