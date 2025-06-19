package Admin_service

import (
	"cadre-management/models"
)

type Comfirmpoexp struct {
	CadreID string
}

func (c *Comfirmpoexp) Comfirmpoexp() error {
	if err := models.Comfirmpoexp(c.CadreID); err != nil {
		return err
	}
	return nil
}

type GetCadreInfoModByPage struct {
	ID         string
	Name       string
	Department string
	Gender     string
	Audited    int

	PageNum  int
	PageSize int
}

func (g *GetCadreInfoModByPage) GetAll() ([]models.Cadre, error) {
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
	if g.Audited != 0 {
		maps["is_audited"] = g.Audited
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
	return models.DeleteCadreByID(cid.ID)
}
