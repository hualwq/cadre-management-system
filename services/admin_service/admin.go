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

type GetPosexpModID struct {
	ID int
}

func (g *GetPosexpModID) ExistByID() (bool, error) {
	return models.ExistPosexpModByID(g.ID)
}

func (g *GetPosexpModID) Get() (*models.Posexp_mod, error) {
	return models.GetPosexpModByID(g.ID)
}

type GetPoexpModByCadreID struct {
	CadreID string
}

func (g *GetPoexpModByCadreID) ExistByCadreID() (bool, error) {
	return models.ExistPoexpModByCadreID(g.CadreID)
}

func (g *GetPoexpModByCadreID) Get() ([]models.Posexp_mod, error) {
	return models.GetPoexpModByCadreID(g.CadreID)
}

type Comfirmpoexp struct {
	CadreID string
}

func (c *Comfirmpoexp) Comfirmpoexp() error {
	if err := models.Comfirmpoexp(c.CadreID); err != nil {
		return err
	}
	return nil
}

type DeletePosexpByID struct {
	ID int
}

func (d *DeletePosexpByID) Delete() error {
	return models.DeletePosexpByID(d.ID)
}

type DeletePosexpModByID struct {
	ID int
}

func (d *DeletePosexpModByID) DeleteMod() error {
	return models.DeletePosexpByID(d.ID)
}

type GetCadreInfoModByPage struct {
	ID         string
	Name       string
	Department string
	Gender     string
	Audited    *bool

	PageNum  int
	PageSize int
}

func (g *GetCadreInfoModByPage) GetAll() ([]models.Cadre_Modification, error) {
	return models.GetCadreInfoModByPage(g.PageNum, g.PageSize, g.getMaps())
}

func (g *GetCadreInfoModByPage) Count() (int64, error) {
	return models.GetCadreInfoModTotal(g.getMaps())
}

func (g *GetCadreInfoModByPage) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if g.Name != "" {
		maps["name"] = g.Name
	}
	if g.Department != "" {
		maps["department"] = g.Department
	}
	if g.Audited != nil {
		maps["is_audited"] = *g.Audited
	}
	if g.ID != "" {
		maps["user_id"] = g.ID
	}
	if g.Gender != "" {
		maps["gender"] = g.Gender
	}

	return maps
}

type PositionHistoryDelete struct {
	ID int
}

func (phd *PositionHistoryDelete) Delete() error {
	return models.DeletePositionHistoryByID(phd.ID)
}

type CadreInfoDelete struct {
	ID string
}

func (cid *CadreInfoDelete) Delete() error {
	return models.DeleteCadreInfoByID(cid.ID)
}
