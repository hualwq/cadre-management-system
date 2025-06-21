package Positionhistory_service

import (
	"cadre-management/models"
)

type Positionhistory struct {
	CadreID      string `json:"user_id"`
	Department   string `json:"department"`
	Category     string `json:"category"`
	Office       string `json:"office"`
	AcademicYear string `json:"academic_year"`
	Year         uint   `json:"applied_at_year"`
	Month        uint   `json:"applied_at_month"`
	Day          uint   `json:"applied_at_day"`
	Positions    string `json:"positions"`
	Audited      int    `json:"is_audited"`
	DepartmentID int    `json:"department_id"`

	ID       int
	PageNum  int
	PageSize int
}

type Posexp struct {
	CadreID    string `json:"user_id"`
	Posyear    string `json:"year"`
	Department string `json:"department"`
	Pos        string `json:"position"`
	PosID      int    `json:"posid"`

	ID       int
	PageNum  int
	PageSize int
}

func (p *Positionhistory) Get() (*models.Positionhistory, error) {
	return models.GetPositionHistoryModByID(p.ID)
}

func (p *Posexp) ExistByID() (bool, error) {
	return models.ExistPosexpByID(p.ID)
}

func (p *Positionhistory) AddPositionhistory() (int, error) {
	positionHistory := map[string]interface{}{
		"user_id":          p.CadreID,
		"department":       p.Department,
		"category":         p.Category,
		"office":           p.Office,
		"academic_year":    p.AcademicYear,
		"applied_at_year":  p.Year,
		"applied_at_month": p.Month,
		"applied_at_day":   p.Day,
		"department_id":    p.DepartmentID,
	}

	id, err := models.AddPositionhistory(positionHistory)
	if id == -1 || err != nil {
		return -1, err
	}

	return id, nil
}

func (s *Posexp) GetAll() ([]models.Posexp, error) {
	return models.GetPosExpByPosID(s.PosID)
}

func (s *Posexp) Count() (int64, error) {
	return models.GetPosExpTotalByPosID(s.PosID)
}

func (s *Posexp) Get() ([]models.Posexp, error) {
	return models.GetPosExpByPosID(s.PosID)
}

func (p *Posexp) Addyearposition() error {
	Pos := map[string]interface{}{
		"user_id":    p.CadreID,
		"year":       p.Posyear,
		"department": p.Department,
		"position":   p.Pos,
		"posid":      p.PosID,
	}

	if err := models.Addyearpositon(Pos); err != nil {
		return err
	}
	return nil
}

func (p *Positionhistory) GetAll() ([]models.Positionhistory, error) {
	return models.GetPositionHistories(p.PageNum, p.PageSize, p.getMaps())
}

func (p *Positionhistory) Count() (int64, error) {
	return models.GetPositionHistoryModTotal(p.getMaps())
}

func (p *Positionhistory) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

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
	if p.Audited != 0 {
		maps["is_audited"] = p.Audited
	}
	if p.DepartmentID != 0 {
		maps["department_id"] = p.DepartmentID
	}

	return maps
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

func (p *Positionhistory) EditPositionhistorymod() error {
	data := map[string]interface{}{
		"user_id":          p.CadreID,
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

func (p *Positionhistory) ExistByID() (bool, int, error) {
	return models.ExistPositionHistoryByID(p.ID)
}

func (p *Positionhistory) DeleteByID() error {
	return models.DeletePositionHistoryByID(p.ID)
}

func (p *Posexp) DeleteByID() error {
	return models.DeletePosexpByID(p.ID)
}

func (p *Positionhistory) ConfirmPositionHistory() error {
	if err := models.ComfirmPositionhistory(p.ID); err != nil {
		return err
	}

	return nil
}
