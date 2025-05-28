package admin_service

import (
	"cadre-management/models"
	"strings"
)

type AssessmentService struct{}
type PositionHistoryService struct{}
type PositionHistory_mod struct{}
type PositionHistory struct {
	Name         string `json:"name"`
	CadreID      string `json:"cadre_id"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	Department   string `json:"department"`
	Category     string `json:"category"`
	Office       string `json:"office"`
	AcademicYear string `json:"academic_year"`
	Positions    string `json:"positions"`
	AppliedAt    string `json:"applied_at"`
}

func (p *PositionHistory) ConfirmPositionHistory() error {
	if err := models.ComfirmPositionhistory(p.CadreID); err != nil {
		return err
	}

	return nil
}

func (p *PositionHistory_mod) GetPositionHistory_mod(CadreID string) (*models.PositionHistory_mod, error) {
	return models.GetPositionHistory_mod(CadreID)
}

func (p *PositionHistory) AddPositionHistory() error {
	positionHistory := map[string]interface{}{
		"cadre_id":      p.CadreID,
		"name":          p.Name,
		"phone_number":  p.PhoneNumber,
		"email":         p.Email,
		"department":    p.Department,
		"category":      p.Category,
		"office":        p.Office,
		"academic_year": p.AcademicYear,
		"positions":     strings.Split(p.Positions, ","), // Convert comma-separated string to slice
		"applied_at":    p.AppliedAt,
	}

	if err := models.EditPositionHistory(positionHistory); err != nil {
		return err
	}

	return nil
}

func (s *PositionHistoryService) GetPositionHistoryList_page(page, pageSize int) ([]models.PositionHistoryBrief_mod, int64, error) {
	return models.GetPositionhistoryList_mod_page(page, pageSize)
}
