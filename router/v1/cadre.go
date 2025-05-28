package v1

import (
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/pkg/setting"
	"cadre-management/pkg/utils"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"cadre-management/services/assessment_mod_service"
	"cadre-management/services/cadre_service"
	"cadre-management/services/familymember_service"
	"cadre-management/services/resume_service"

	"github.com/gin-gonic/gin"
)

type AddCadreInfoForm_modifications struct {
	ID                        string `json:"user_id" binding:"required"`
	Name                      string `json:"name" binding:"required"`
	Gender                    string `json:"gender" binding:"required"`
	BirthDate                 string `json:"birth_date" binding:"required"`
	Age                       uint8  `json:"age"`
	EthnicGroup               string `json:"ethnic_group" binding:"required"`
	NativePlace               string `json:"native_place" binding:"required"`
	BirthPlace                string `json:"birth_place" binding:"required"`
	PoliticalStatus           string `json:"political_status" binding:"required"`
	WorkStartDate             string `json:"work_start_date" binding:"required"`
	HealthStatus              string `json:"health_status" binding:"required"`
	ProfessionalTitle         string `json:"professional_title" binding:"required"`
	Specialty                 string `json:"specialty" binding:"required"`
	Phone                     string `json:"phone"`
	CurrentPosition           string `json:"current_position" binding:"required"`
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

type EditCadreInfoForm_modifications struct {
	ID                        string `json:"user_id" binding:"required"`
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

type AddAssessment_modFrom struct {
	CadreID     string `json:"user_id"`
	Department  string `json:"department"`
	Category    string `json:"category"`
	AssessDept  string `json:"assess_dept"`
	Year        int    `json:"year"`
	WorkSummary string `json:"work_summary"`
}

type AddPositionHistory_modForm struct {
	CadreID      string `json:"user_id"`
	Department   string `json:"department"`
	Category     string `json:"category"`
	Office       string `json:"office"`
	AcademicYear string `json:"academic_year"`
	Year         uint   `json:"applied_at_year"`
	Month        uint   `json:"applied_at_month"`
	Day          uint   `json:"applied_at_day"`
}

type PosexpForm_mod struct {
	CadreID    string `json:"user_id"`
	Posyear    string `json:"year"`
	Department string `json:"department"`
	Pos        string `json:"position"`
}

type ResumeEntry_modForm struct {
	CadreID      string `json:"user_id"`
	StartDate    string `json:"start_date"`           // 格式：2007.09 或 2019.12
	EndDate      string `json:"end_date"`             // 格式：2011.07 或 "至今"
	Organization string `json:"organization"`         // 工作单位或学校
	Department   string `json:"department,omitempty"` // 学院/部门，可选
	Position     string `json:"position,omitempty"`   // 职务/身份，可选
}

type FamilyMember_modForm struct {
	CadreID         string `gorm:"not null" json:"user_id"`
	Relation        string `gorm:"type:varchar(20);not null" json:"relation"`
	Name            string `gorm:"type:varchar(50);not null" json:"name"`
	BirthDate       string `gorm:"type:date" json:"birth_date,omitempty"`
	PoliticalStatus string `gorm:"type:varchar(50)" json:"political_status,omitempty"`
	WorkUnit        string `gorm:"type:varchar(200)" json:"work_unit,omitempty"`
}

type EditResumeEntry_modForm struct {
	CadreID      string `json:"user_id"`
	StartDate    string `json:"start_date"`           // 格式：2007.09 或 2019.12
	EndDate      string `json:"end_date"`             // 格式：2011.07 或 "至今"
	Organization string `json:"organization"`         // 工作单位或学校
	Department   string `json:"department,omitempty"` // 学院/部门，可选
	Position     string `json:"position,omitempty"`   // 职务/身份，可选
}

type EditFamilyMember_modForm struct {
	ID              int    `json:"id"`
	CadreID         string `json:"user_id"`
	Relation        string `json:"relation"`
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date,omitempty"`
	PoliticalStatus string `json:"political_status,omitempty"`
	WorkUnit        string `json:"work_unit,omitempty"`
}

type GetPositionHistoryForm struct {
	CadreID string `json:"user_id" binding:"required"`
}

func GetPositionHistories(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	name := c.Query("name")
	department := c.Query("department")
	category := c.Query("category")
	office := c.Query("office")
	academicYear := c.Query("academic_year")
	positions := c.Query("positions")
	year := uint(com.StrTo(c.Query("year")).MustInt())
	month := uint(com.StrTo(c.Query("month")).MustInt())
	day := uint(com.StrTo(c.Query("day")).MustInt())

	positionHistoryService := cadre_service.PositionHistory{
		ID:           id,
		Name:         name,
		Department:   department,
		Category:     category,
		Office:       office,
		AcademicYear: academicYear,
		Positions:    positions,
		Year:         year,
		Month:        month,
		Day:          day,
		PageNum:      utils.GetPage(c),
		PageSize:     setting.AppSetting.PageSize,
	}
	positionHistories, err := positionHistoryService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POSITION_HISTORIES_FAIL, nil)
		return
	}

	count, err := positionHistoryService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POSITION_HISTORIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": positionHistories,
		"total": count,
	})
}

func AddCadreInfo_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCadreInfoForm_modifications
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	cadreService := cadre_service.CadreInfo_mod{
		ID:                        form.ID,
		Name:                      form.Name,
		Gender:                    form.Gender,
		BirthDate:                 form.BirthDate,
		Age:                       form.Age,
		EthnicGroup:               form.EthnicGroup,
		NativePlace:               form.NativePlace,
		BirthPlace:                form.BirthPlace,
		PoliticalStatus:           form.PoliticalStatus,
		WorkStartDate:             form.WorkStartDate,
		HealthStatus:              form.HealthStatus,
		ProfessionalTitle:         form.ProfessionalTitle,
		Specialty:                 form.Specialty,
		CurrentPosition:           form.CurrentPosition,
		AwardsAndPunishments:      form.AwardsAndPunishments,
		AnnualAssessment:          form.AnnualAssessment,
		Email:                     form.Email,
		FilledBy:                  form.FilledBy,
		FullTimeEducationDegree:   form.FullTimeEducationDegree,
		FullTimeEducationSchool:   form.FullTimeEducationSchool,
		OnTheJobEducationDegree:   form.OnTheJobEducationDegree,
		OnTheJobEducationSchool:   form.OnTheJobEducationSchool,
		ReportingUnit:             form.ReportingUnit,
		ApprovalAuthority:         form.ApprovalAuthority,
		AdministrativeAppointment: form.AdministrativeAppointment,
		Phone:                     form.Phone,
	}

	if err := cadreService.AddCadreInfo(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CADRE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddAssessment_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddAssessment_modFrom
	)

	// 参数校验
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 构造服务层结构体
	CadreService := assessment_mod_service.Assessment_mod{
		CadreID:     form.CadreID,
		Department:  form.Department,
		Category:    form.Category,
		AssessDept:  form.AssessDept,
		Year:        form.Year,
		WorkSummary: form.WorkSummary,
	}

	// 调用添加逻辑
	if err := CadreService.AddAssessment_mod(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ASSESSMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddPositionHistory_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddPositionHistory_modForm
	)

	// 参数校验
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 构造服务层结构体
	PositionService := cadre_service.PositionHistory_mod{
		CadreID:      form.CadreID,
		Department:   form.Department,
		Category:     form.Category,
		Office:       form.Office,
		AcademicYear: form.AcademicYear,
		Year:         form.Year,
		Month:        form.Month,
		Day:          form.Day,
	}

	// 调用添加逻辑
	if err := PositionService.AddPositionHistory_mod(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_POSITIONHISTORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Addyearposition_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PosexpForm_mod
	)

	// 1. 参数绑定和校验
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 2. 构造服务层结构体
	posService := cadre_service.Posexp_mod{
		CadreID:    form.CadreID,
		Posyear:    form.Posyear,
		Department: form.Department,
		Pos:        form.Pos,
	}

	// 3. 调用添加逻辑
	if err := posService.Addyearposition_mod(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_POSITION_FAIL, nil)
		return
	}

	// 4. 返回成功响应
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func EditInfo_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form EditCadreInfoForm_modifications
	)

	// 绑定并验证 JSON 数据
	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 构造 service 层对象
	cadreService := cadre_service.CadreInfo_mod{
		ID:                        form.ID,
		Name:                      form.Name,
		Gender:                    form.Gender,
		BirthDate:                 form.BirthDate,
		Age:                       form.Age,
		EthnicGroup:               form.EthnicGroup,
		NativePlace:               form.NativePlace,
		BirthPlace:                form.BirthPlace,
		PoliticalStatus:           form.PoliticalStatus,
		WorkStartDate:             form.WorkStartDate,
		HealthStatus:              form.HealthStatus,
		ProfessionalTitle:         form.ProfessionalTitle,
		Specialty:                 form.Specialty,
		Phone:                     form.Phone,
		CurrentPosition:           form.CurrentPosition,
		AwardsAndPunishments:      form.AwardsAndPunishments,
		AnnualAssessment:          form.AnnualAssessment,
		Email:                     form.Email,
		FilledBy:                  form.FilledBy,
		FullTimeEducationDegree:   form.FullTimeEducationDegree,
		FullTimeEducationSchool:   form.FullTimeEducationSchool,
		OnTheJobEducationDegree:   form.OnTheJobEducationDegree,
		OnTheJobEducationSchool:   form.OnTheJobEducationSchool,
		ReportingUnit:             form.ReportingUnit,
		ApprovalAuthority:         form.ApprovalAuthority,
		AdministrativeAppointment: form.AdministrativeAppointment,
	}

	// 检查该干部是否存在
	exists, err := cadreService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXSIT_CADREINFO, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusNotFound, e.ERROR_NOT_EXSIT_CADREINFO, nil)
		return
	}

	// 执行修改操作
	if err := cadreService.Edit(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_CADRE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddResume_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ResumeEntry_modForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	resumeservice := resume_service.ResumeEntry_mod{
		CadreID:      form.CadreID,
		StartDate:    form.StartDate,
		EndDate:      form.EndDate,
		Organization: form.Organization,
		Department:   form.Department,
		Position:     form.Position,
	}

	if err := resumeservice.Add_resume_mod(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RESUME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Addfamilymember_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form FamilyMember_modForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	familymemberservice := familymember_service.FamilyMember_mod{
		CadreID:         form.CadreID,
		Relation:        form.Relation,
		Name:            form.Name,
		BirthDate:       form.BirthDate,       // 可选
		PoliticalStatus: form.PoliticalStatus, // 可选
		WorkUnit:        form.WorkUnit,        // 可选
	}

	if err := familymemberservice.Add_familymember_mod(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAMILYMEBER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Editfamilymember_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form EditFamilyMember_modForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	familymemberservice := familymember_service.FamilyMember_mod{
		ID:              form.ID,
		CadreID:         form.CadreID,
		Relation:        form.Relation,
		Name:            form.Name,
		BirthDate:       form.BirthDate,
		PoliticalStatus: form.PoliticalStatus,
		WorkUnit:        form.WorkUnit,
	}

	exists, err := familymemberservice.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXSIT_FAMILYMEMBER, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusNotFound, e.ERROR_NOT_EXSIT_FAMILYMEMBER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteFamilyMemberModification(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	familyMemberModificationsService := familymember_service.FamilyMemberModifications{ID: id}
	exists, err := familyMemberModificationsService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}

	err = familyMemberModificationsService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetResumeEntryModificationByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	resumeEntryModificationsService := resume_service.ResumeEntryModifications{ID: id}
	exists, err := resumeEntryModificationsService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESUME_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_RESUME_FAIL, nil)
		return
	}

	resumeEntryModification, err := resumeEntryModificationsService.GetByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESUME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, resumeEntryModification)
}

// @Summary 根据 CadreID 获取履历条目修改记录列表
// @Produce  json
// @Param cadre_id path string true "CadreID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/resume_entry_modifications/cadre/{cadre_id} [get]
func GetResumeEntryModificationsByCadreID(c *gin.Context) {
	appG := app.Gin{C: c}
	cadreID := c.Query("user_id")
	valid := validation.Validation{}
	valid.Required(cadreID, "user_id").Message("CadreID 不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	resumeEntryModificationsService := resume_service.ResumeEntryModifications{
		CadreID: cadreID,
	}
	resumeEntryModifications, err := resumeEntryModificationsService.GetByCadreID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESUME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, resumeEntryModifications)
}

// @Summary 根据 ID 删除单个履历条目修改记录
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/resume_entry_modifications/{id} [delete]
func DeleteResumeEntryModificationByID(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID 必须大于 0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	resumeEntryModificationsService := resume_service.ResumeEntryModifications{ID: id}
	exists, err := resumeEntryModificationsService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESUME_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_RESUME_FAIL, nil)
		return
	}

	err = resumeEntryModificationsService.DeleteByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESUME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
