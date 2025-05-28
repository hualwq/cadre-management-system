package v1

import (
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/pkg/setting"
	"cadre-management/pkg/utils"
	"cadre-management/services/admin_service"
	"cadre-management/services/assessment_mod_service"
	"cadre-management/services/cadre_service"
	"cadre-management/services/familymember_service"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

type AddPositionHistoryForm struct {
	CadreID      string `json:"user_id" binding:"required"`
	Department   string `json:"department" binding:"required"`
	PositionType string `json:"position_type" binding:"required,oneof=专职团干部 兼职团干部 教师 学生 其他"`
	StartDate    string `json:"start_date" binding:"required"`
	EndDate      string `json:"end_date,omitempty"`
	Description  string `json:"description,omitempty"`
}

type GetCadreInfoForm_mod struct {
	CadreID string `json:"user_id" binding:"required"`
}

type ComfirmcadreForm struct {
	CadreID string `json:"user_id" binding:"required"`
}

type ComfirmAssessmentForm struct {
	ID     int    `json:"id" binding:"required"`
	Result string `json:"result" binding:"required"`
}

type ComfirmPositionhistoryForm struct {
	CadreID string `json:"user_id" binding:"required"`
}

func ComfirmAssessment(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmAssessmentForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	AssessmentService := assessment_mod_service.ComfirmAssessment{
		Grade: form.Result,
		ID:    form.ID,
	}

	if err := AssessmentService.ComfirmAssessment(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ASSESSMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetCadreInfo_mod(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmcadreForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	Cadreservice := cadre_service.CadreInfo_mod{ID: form.CadreID}
	cadre, err := Cadreservice.GetCadreInfo()
	if err != nil {
		// 根据错误类型返回不同状态码
		if strings.Contains(err.Error(), "不存在") {
			appG.Response(http.StatusNotFound, e.ERROR_GET_CADREINFO_FAIL, nil)
		} else {
			appG.Response(http.StatusInternalServerError, e.ERROR_GET_CADREINFO_FAIL, nil)
		}
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, cadre)
}

func ComfirmCadreInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmcadreForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	cadreService := cadre_service.CadreInfo_mod{ID: form.CadreID}

	if err := cadreService.ComfirmCadreInfo(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CADRE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func ConfirmPositionhistory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmPositionhistoryForm
	)

	// 参数校验
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 构造服务层结构体
	confirmService := admin_service.PositionHistory{
		CadreID: form.CadreID,
	}

	// 调用确认逻辑
	if err := confirmService.ConfirmPositionHistory(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CONFIRM_POSITIONHISTORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetAssessmentsMod(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	cadreID := c.Query("cadre_id")
	department := c.Query("department")
	category := c.Query("category")
	assessDept := c.Query("assess_dept")
	yearStr := c.Query("year")
	year := 0
	if yearStr != "" {
		year = com.StrTo(yearStr).MustInt()
	}
	auditedStr := c.Query("audited")
	var audited *bool
	if auditedStr != "" {
		boolValue, err := strconv.ParseBool(auditedStr)
		if err == nil {
			audited = &boolValue
		}
	}

	assessmentService := assessment_mod_service.Assessment_mod{
		Name:       name,
		CadreID:    cadreID,
		Department: department,
		Category:   category,
		AssessDept: assessDept,
		Year:       year,
		Audited:    audited,
		PageNum:    utils.GetPage(c),
		PageSize:   setting.AppSetting.PageSize,
	}
	assessments, err := assessmentService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ASSESSMENTS_FAIL, nil)
		return
	}

	count, err := assessmentService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ASSESSMENTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": assessments,
		"total": count,
	})
}

func GetAssessmentsModByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	assesementService := assessment_mod_service.GetAssessment_modID{ID: id}
	exists, err := assesementService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ASSESEMENT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ASSESEMENT, nil)
		return
	}

	assesement, err := assesementService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ASSESSMENTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, assesement)
}

func GetPositionHistory_mod(c *gin.Context) {
	appG := app.Gin{C: c}
	cadreID := c.Param("id")
	if cadreID == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	service := admin_service.PositionHistory_mod{}
	positionHistory, err := service.GetPositionHistory_mod(cadreID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, positionHistory)
}

func GetPositionHistoriesMod(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	department := c.Query("department")
	category := c.Query("category")
	office := c.Query("office")
	academicYear := c.Query("academic_year")
	auditedStr := c.Query("audited")
	var audited *bool
	if auditedStr != "" {
		boolValue, err := strconv.ParseBool(auditedStr)
		if err == nil {
			audited = &boolValue
		}
	}

	positionHistoryService := cadre_service.GetPositionHistory_mod{
		Name:         name,
		Department:   department,
		Category:     category,
		Office:       office,
		AcademicYear: academicYear,
		Audited:      audited,
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
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POSITION_HISTORY_MOD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": positionHistories,
		"total": count,
	})
}

func GetPositionHistoryMod(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	positionHistoryModService := cadre_service.PositionHistoryModService{
		ID: id,
	}
	positionHistoryMod, err := positionHistoryModService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POSITION_HISTORY_MOD_FAIL, nil)
		return
	}

	if positionHistoryMod == nil {
		appG.Response(http.StatusNotFound, e.ERROR_GET_POSITION_HISTORY_MOD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, positionHistoryMod)
}

func GetFamilyMemberModification(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
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

	familyMemberModification, err := familyMemberModificationsService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, familyMemberModification)
}

func GetFamilyMemberModificationsByCadreID(c *gin.Context) {
	appG := app.Gin{C: c}
	cadreID := c.Query("user_id")
	valid := validation.Validation{}
	valid.Required(cadreID, "user_id").Message("user_id 不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	familyMemberModificationsService := familymember_service.FamilyMemberModifications_cadreinfo{
		CadreID: cadreID,
	}
	familyMemberModifications, err := familyMemberModificationsService.GetByCadreID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, familyMemberModifications)
}
