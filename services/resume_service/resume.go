package resume_service

import "cadre-management/models"

type ResumeEntry struct {
	CadreID      string `json:"user_id"`
	StartDate    string `json:"start_date"`           // 格式：2007.09 或 2019.12
	EndDate      string `json:"end_date"`             // 格式：2011.07 或 "至今"
	Organization string `json:"organization"`         // 工作单位或学校
	Department   string `json:"department,omitempty"` // 学院/部门，可选
	Position     string `json:"position,omitempty"`   // 职务/身份，可选
}

type ResumeEntry_mod struct {
	CadreID      string `json:"user_id"`
	StartDate    string `json:"start_date"`           // 格式：2007.09 或 2019.12
	EndDate      string `json:"end_date"`             // 格式：2011.07 或 "至今"
	Organization string `json:"organization"`         // 工作单位或学校
	Department   string `json:"department,omitempty"` // 学院/部门，可选
	Position     string `json:"position,omitempty"`   // 职务/身份，可选
}

func (r *ResumeEntry_mod) Add_resume_mod() error {
	Cinfo := map[string]interface{}{
		"user_id":      r.CadreID,
		"start_date":   r.StartDate,
		"end_date":     r.EndDate,
		"organization": r.Organization,
		"department":   r.Department,
		"position":     r.Position,
	}

	if err := models.Add_resume_mod(Cinfo); err != nil {
		return err
	}
	return nil
}

type ResumeEntryModifications struct {
	ID      int
	CadreID string
}

func (rem *ResumeEntryModifications) ExistByID() (bool, error) {
	return models.ExistResumeEntryModificationByID(rem.ID)
}

func (rem *ResumeEntryModifications) GetByID() (*models.ResumeEntry_modifications, error) {
	return models.GetResumeEntryModificationByID(rem.ID)
}

func (rem *ResumeEntryModifications) GetByCadreID() ([]models.ResumeEntry_modifications, error) {
	return models.GetResumeEntryModificationsByCadreID(rem.CadreID)
}

func (rem *ResumeEntryModifications) DeleteByID() error {
	return models.DeleteResumeEntryModificationByID(rem.ID)
}
