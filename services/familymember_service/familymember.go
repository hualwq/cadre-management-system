package familymember_service

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
}

type FamilyMember_mod struct {
	ID              int    `json:"id"`
	CadreID         string `json:"user_id"`
	Relation        string `json:"relation"`
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date,omitempty"`
	PoliticalStatus string `json:"political_status,omitempty"`
	WorkUnit        string `json:"work_unit,omitempty"`
}

func (fm *FamilyMember_mod) Add_familymember_mod() error {
	Cinfo := map[string]interface{}{
		"user_id":          fm.CadreID,
		"relation":         fm.Relation,
		"name":             fm.Name,
		"birth_date":       fm.BirthDate,
		"political_status": fm.PoliticalStatus,
		"work_unit":        fm.WorkUnit,
	}

	if err := models.Add_familymember_mod(Cinfo); err != nil {
		return err
	}
	return nil
}

func (f *FamilyMember_mod) ExistByID() (bool, error) {
	return models.ExistByID(f.ID)
}

func (fm *FamilyMember_mod) EditFamilyMemberMod() error {
	data := map[string]interface{}{
		"user_id":          fm.CadreID,
		"relation":         fm.Relation,
		"name":             fm.Name,
		"birth_date":       fm.BirthDate,
		"political_status": fm.PoliticalStatus,
		"work_unit":        fm.WorkUnit,
	}
	return models.EditFamilyMemberModification(fm.ID, data)
}

type FamilyMemberModifications struct {
	ID int
}

func (fmm *FamilyMemberModifications) ExistByID() (bool, error) {
	return models.ExistFamilyMemberModificationByID(fmm.ID)
}

func (fmm *FamilyMemberModifications) Get() (*models.FamilyMember_modifications, error) {
	return models.GetFamilyMemberModificationByID(fmm.ID)
}

func (fmm *FamilyMemberModifications) Delete() error {
	return models.DeleteFamilyMemberModificationByID(fmm.ID)
}

type FamilyMemberModifications_cadreinfo struct {
	CadreID string
}

// GetByCadreID 根据 CadreID 获取家庭成员修改记录
func (fmm *FamilyMemberModifications_cadreinfo) GetByCadreID() ([]models.FamilyMember_modifications, error) {
	return models.GetFamilyMemberModificationsByCadreID(fmm.CadreID)
}

type FamilyMemberDelete struct {
	ID int
}

func (fmd *FamilyMemberDelete) Delete() error {
	return models.DeleteFamilyMemberByID(fmd.ID)
}

type Comfirmfamilymember struct {
	ID int
	// 可以添加其他必要的字段
}

func (c Comfirmfamilymember) Comfirmfamilymember() error {
	if err := models.Comfirmfamilymember(c.ID); err != nil {
		return err
	}

	return nil
}
