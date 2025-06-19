package Familymember_service

import (
	"cadre-management/models"
)

type FamilyMember struct {
	ID              int    `json:"id"`
	CadreID         string `json:"cadre_id"`
	Relation        string `json:"relation"`
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date,omitempty"`
	PoliticalStatus string `json:"political_status,omitempty"`
	WorkUnit        string `json:"work_unit,omitempty"`
	IsAudited       int    `json:"is_audited"`
}

func (fm *FamilyMember) AddFamilyMember() error {
	Cinfo := map[string]interface{}{
		"user_id":          fm.CadreID,
		"relation":         fm.Relation,
		"name":             fm.Name,
		"birth_date":       fm.BirthDate,
		"political_status": fm.PoliticalStatus,
		"work_unit":        fm.WorkUnit,
	}

	if err := models.Addfamilymember(Cinfo); err != nil {
		return err
	}
	return nil
}

func (f *FamilyMember) ExistByID() (bool, error) {
	return models.ExistByID(f.ID)
}

func (fm *FamilyMember) EditFamilyMemberMod() error {
	data := map[string]interface{}{
		"user_id":          fm.CadreID,
		"relation":         fm.Relation,
		"name":             fm.Name,
		"birth_date":       fm.BirthDate,
		"political_status": fm.PoliticalStatus,
		"work_unit":        fm.WorkUnit,
	}
	return models.EditFamilyMember(fm.ID, data)
}

func (fmm *FamilyMember) Get() (*models.Familymember, error) {
	return models.GetFamilyMemberByID(fmm.ID)
}

func (fmm *FamilyMember) Delete() error {
	return models.DeleteFamilyMemberByID(fmm.ID)
}

// GetByCadreID 根据 CadreID 获取家庭成员修改记录
func (fmm *FamilyMember) GetByCadreID() ([]models.Familymember, error) {
	return models.GetFamilyMembersByCadreID(fmm.CadreID)
}

func (c *FamilyMember) Comfirmfamilymember() error {
	if err := models.Comfirmfamilymember(c.ID); err != nil {
		return err
	}

	return nil
}
