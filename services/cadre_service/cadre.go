package cadre_service

import (
	"cadre-management/models"
)

type PositionHistory_mod struct {
	CadreID      string `json:"user_id"`
	Department   string `json:"department"`
	Category     string `json:"category"`
	Office       string `json:"office"`
	AcademicYear string `json:"academic_year"`
	Year         uint   `json:"applied_at_year"`
	Month        uint   `json:"applied_at_month"`
	Day          uint   `json:"applied_at_day"`
}

type GetPositionHistory_mod struct {
	Name         string
	Department   string
	Category     string
	Office       string
	AcademicYear string
	Audited      *bool

	PageNum  int
	PageSize int
}

type PositionHistory struct {
	Name         string
	Department   string
	Category     string
	Office       string
	AcademicYear string
	Positions    string
	Year         uint
	Month        uint
	Day          uint

	ID       int
	PageNum  int
	PageSize int
}

type Posexp struct {
	CadreID    string `json:"user_id"`
	Posyear    string `json:"year"`
	Department string `json:"department"`
	Pos        string `json:"position"`

	PageNum  int
	PageSize int
}

type PositionHistoryModService struct {
	ID int `json:"id"`
}

func (p *PositionHistoryModService) Get() (*models.PositionHistory_mod, error) {
	return models.GetPositionHistoryModByID(p.ID)
}

func (p *PositionHistory_mod) AddPositionHistory_mod() error {
	positionHistory := map[string]interface{}{
		"user_id":          p.CadreID,
		"department":       p.Department,
		"category":         p.Category,
		"office":           p.Office,
		"academic_year":    p.AcademicYear,
		"applied_at_year":  p.Year,
		"applied_at_month": p.Month,
		"applied_at_day":   p.Day,
	}

	if err := models.AddPositionHistory_mod(positionHistory); err != nil {
		return err
	}

	return nil
}

type CadreInfo_mod struct {
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

type Posexp_mod struct {
	CadreID    string `json:"user_id"`
	Posyear    string `json:"year"`
	Department string `json:"department"`
	Pos        string `json:"position"`
}

func (c *CadreInfo_mod) DeleteByID() error {
	return models.DeleteCadreInfoByID(c.ID)
}

func (c *CadreInfo_mod) GetCadreInfo() (*models.Cadre_Modification, error) {
	return models.GetCadre(c.ID)
}

func (c *CadreInfo_mod) AddCadreInfo() error {
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

	if err := models.AddCadreInfo_mod(Cinfo); err != nil {
		return err
	}
	return nil
}

func (p *Posexp_mod) Addyearposition_mod() error {
	Pos := map[string]interface{}{
		"user_id":    p.CadreID,
		"year":       p.Posyear,
		"department": p.Department,
		"position":   p.Pos,
	}

	if err := models.Addyearpositon(Pos); err != nil {
		return err
	}
	return nil
}

func (c *CadreInfo_mod) ExistByID() (bool, error) {
	return models.ExistCadreInfoByID(c.ID)
}

func (c *CadreInfo_mod) Edit() error {
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

func (c *CadreInfo_mod) ComfirmCadreInfo() error {
	cadreid := c.ID
	return models.AddCadreInfofromMod(cadreid)
}

func (p *PositionHistory) GetAll() ([]models.PositionHistory, error) {
	return models.GetPositionHistories(p.PageNum, p.PageSize, p.getMaps())
}

func (p *PositionHistory) Count() (int64, error) {
	return models.GetPositionHistoryTotal(p.getMaps())
}

func (p *PositionHistory) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if p.Name != "" {
		maps["name"] = p.Name
	}
	if p.Department != "" {
		maps["department"] = p.Department
	}
	if p.Category != "" {
		maps["category"] = p.Category
	}
	if p.Office != "" {
		maps["office"] = p.Office
	}
	if p.AcademicYear != "" {
		maps["academic_year"] = p.AcademicYear
	}
	if p.Positions != "" {
		maps["positions"] = p.Positions
	}
	if p.Year > 0 {
		maps["applied_at_year"] = p.Year
	}
	if p.Month > 0 {
		maps["applied_at_month"] = p.Month
	}
	if p.Day > 0 {
		maps["applied_at_day"] = p.Day
	}

	return maps
}

func (p *Posexp) GetAll() ([]models.Posexp, error) {
	return models.GetPosexps(p.PageNum, p.PageSize, p.getMaps())
}

func (p *Posexp) Count() (int64, error) {
	return models.GetPosexpTotal(p.getMaps())
}

func (p *Posexp) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if p.CadreID != "" {
		maps["user_id"] = p.CadreID
	}
	if p.Posyear != "" {
		maps["year"] = p.Posyear
	}
	if p.Department != "" {
		maps["department"] = p.Department
	}
	if p.Pos != "" {
		maps["position"] = p.Pos
	}

	return maps
}

func (p *GetPositionHistory_mod) GetAll() ([]models.PositionHistory_mod, error) {
	return models.GetPositionHistoriesMod(p.PageNum, p.PageSize, p.getMaps())
}

func (p *GetPositionHistory_mod) Count() (int64, error) {
	return models.GetPositionHistoryModTotal(p.getMaps())
}

func (p *GetPositionHistory_mod) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if p.Name != "" {
		maps["name"] = p.Name
	}
	if p.Department != "" {
		maps["department"] = p.Department
	}
	if p.Category != "" {
		maps["category"] = p.Category
	}
	if p.Office != "" {
		maps["office"] = p.Office
	}
	if p.AcademicYear != "" {
		maps["academic_year"] = p.AcademicYear
	}
	if p.Audited != nil {
		maps["is_audited"] = *p.Audited
	}

	return maps
}

type PositionHistoryModEdit struct {
	ID           int    `json:"id"`
	CadreID      string `json:"user_id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Department   string `json:"department"`
	Category     string `json:"category"`
	Office       string `json:"office"`
	AcademicYear string `json:"academic_year"`
	Positions    string `json:"positions"`
	Year         uint   `json:"applied_at_year"`
	Month        uint   `json:"applied_at_month"`
	Day          uint   `json:"applied_at_day"`
}

func (p *PositionHistoryModEdit) EditPositionhistorymod() error {
	data := map[string]interface{}{
		"user_id":          p.CadreID,
		"name":             p.Name,
		"phone_number":     p.PhoneNumber,
		"email":            p.Email,
		"department":       p.Department,
		"category":         p.Category,
		"office":           p.Office,
		"academic_year":    p.AcademicYear,
		"positions":        p.Positions,
		"applied_at_year":  p.Year,
		"applied_at_month": p.Month,
		"applied_at_day":   p.Day,
	}

	return models.EditPositionHistoryMod(p.ID, data)
}

func (p *PositionHistoryModEdit) ExistByID() (bool, error) {
	return models.ExistPositionHistoryByID(p.ID)
}

func (p *PositionHistoryModEdit) DeleteByID() error {
	return models.DeleteFamilyMemberModificationByID(p.ID)
}
