package assessment_mod_service

import (
	"cadre-management/models"
)

type Assessment_mod struct {
	ID          int
	Name        string
	CadreID     string
	Phone       string
	Email       string
	Department  string
	Category    string
	AssessDept  string
	Year        int
	WorkSummary string
	Grade       string
	Audited     *bool

	PageNum  int
	PageSize int
}

type GetAssessment_modID struct {
	ID int
}

func (a *GetAssessment_modID) ExistByID() (bool, error) {
	return models.ExistAssessmentModByID(a.ID)
}

func (a *GetAssessment_modID) Get() (*models.Assessment_mod, error) {
	return models.GetAssesement(a.ID)
}

type ComfirmAssessment struct {
	ID    int    `json:"id"`
	Grade string `json:"result"`
}

func (c ComfirmAssessment) ComfirmAssessment() error {
	if err := models.ComfirmAssessment(c.ID, c.Grade); err != nil {
		return err
	}

	return nil
}

func (a *Assessment_mod) AddAssessment_mod() error {
	assessment := map[string]interface{}{
		"user_id":      a.CadreID,
		"department":   a.Department,
		"category":     a.Category,
		"assess_dept":  a.AssessDept,
		"work_summary": a.WorkSummary,
		"year":         a.Year,
	}

	if err := models.AddAssessment_mod(assessment); err != nil {
		return err
	}

	return nil
}

func (a *Assessment_mod) Count() (int64, error) {
	return models.GetAssessmentModTotal(a.getMaps())
}

func (a *Assessment_mod) GetAll() ([]models.Assessment_mod, error) {
	return models.GetAssessmentsMod(a.PageNum, a.PageSize, a.getMaps())
}

func (a *Assessment_mod) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if a.Name != "" {
		maps["name"] = a.Name
	}
	if a.CadreID != "" {
		maps["user_id"] = a.CadreID
	}
	if a.Phone != "" {
		maps["phone"] = a.Phone
	}
	if a.Email != "" {
		maps["email"] = a.Email
	}
	if a.Department != "" {
		maps["department"] = a.Department
	}
	if a.Category != "" {
		maps["category"] = a.Category
	}
	if a.AssessDept != "" {
		maps["assess_dept"] = a.AssessDept
	}
	if a.Year != 0 {
		maps["year"] = a.Year
	}
	if a.Audited != nil {
		maps["is_audited"] = *a.Audited
	}

	return maps
}

type DeleteAssessmentModByID struct {
	ID int
}

func (d *DeleteAssessmentModByID) Delete() error {
	return models.DeleteAssessmentByID(d.ID)
}

type DeleteAssessmentByID struct {
	ID int
}

func (d *DeleteAssessmentByID) Delete() error {
	return models.DeleteAssessmentByID(d.ID)
}

func (a *Assessment_mod) ExistByID() (bool, error) {
	return models.ExistAssessmentModByID(a.ID)
}

func (a *Assessment_mod) EditAssessmentMod() error {
	data := map[string]interface{}{
		"name":         a.Name,
		"user_id":      a.CadreID,
		"phone":        a.Phone,
		"email":        a.Email,
		"department":   a.Department,
		"category":     a.Category,
		"assess_dept":  a.AssessDept,
		"year":         a.Year,
		"work_summary": a.WorkSummary,
		"grade":        a.Grade,
	}

	// 调用 model 层更新
	if err := models.EditAssessmentModByID(a.ID, data); err != nil {
		return err
	}

	return nil
}
