package Assessment_service

import (
	"cadre-management/models"
	"fmt"
)

type Assessment struct {
	ID           int
	Name         string
	Phone        string
	Email        string
	Department   string
	Category     string
	AssessDept   string
	Year         int
	WorkSummary  string
	Grade        string
	Audited      int
	UserID       string
	DepartmentID int

	PageNum  int
	PageSize int
}

func (a *Assessment) ExistByID() (bool, error) {
	return models.ExistAssessmentModByID(a.ID)
}

func (a *Assessment) Get() (*models.Assessment, error) {
	return models.GetAssesement(a.ID)
}

type ComfirmAssessment struct {
	ID    int    `json:"id"`
	Grade string `json:"grade"`
}

func (c ComfirmAssessment) ComfirmAssessment() error {
	if err := models.ComfirmAssessment(c.ID, c.Grade); err != nil {
		return err
	}
	fmt.Println("ComfirmAssessmentaaaaa", c.ID, c.Grade)

	return nil
}

func (a *Assessment) AddAssessment() error {
	assessment := map[string]interface{}{
		"user_id":       a.UserID,
		"department":    a.Department,
		"category":      a.Category,
		"assess_dept":   a.AssessDept,
		"work_summary":  a.WorkSummary,
		"year":          a.Year,
		"department_id": a.DepartmentID,
	}

	if err := models.AddAssessment(assessment); err != nil {
		return err
	}

	return nil
}

func (a *Assessment) Count() (int64, error) {
	return models.GetAssessmentModTotal(a.getMaps())
}

func (a *Assessment) GetAll() ([]models.Assessment, error) {
	return models.GetAssessmentsMod(a.PageNum, a.PageSize, a.getMaps())
}

func (a *Assessment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if a.UserID != "" {
		maps["user_id"] = a.UserID
	}

	if a.Name != "" {
		maps["name"] = a.Name
	}
	if a.UserID != "" {
		maps["user_id"] = a.UserID
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
	if a.Audited != 0 {
		maps["is_audited"] = a.Audited
	}
	if a.DepartmentID != 0 {
		maps["department_id"] = a.DepartmentID
	}

	return maps
}

func (d *Assessment) Delete() error {
	return models.DeleteAssessmentByID(d.ID)
}

func (a *Assessment) EditAssessmentMod() error {
	data := map[string]interface{}{
		"name":         a.Name,
		"user_id":      a.UserID,
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
