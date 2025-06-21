package v1

import (
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/pkg/setting"
	"cadre-management/pkg/utils"
	"cadre-management/services/Admin_service"
	"cadre-management/services/Assessment_service"
	"cadre-management/services/Cadre_service"
	"cadre-management/services/Familymember_service"
	"cadre-management/services/Positionhistory_service"
	"cadre-management/services/Resume_service"
	"cadre-management/services/message_service"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

type GetCadreInfoForm_mod struct {
	CadreID string `json:"user_id" binding:"required"`
}

type ComfirmcadreForm struct {
	CadreID string `json:"user_id" binding:"required"`
}

type ComfirmAssessmentForm struct {
	ID    int    `json:"id" binding:"required"`
	Grade string `json:"grade" binding:"required"`
}

type ComfirmPositionhistoryForm struct {
	CadreID string `json:"user_id" binding:"required"`
}

type GetpoexpmodCadreID struct {
	CadreID string `json:"user_id" binding:"required"`
}

type ComfirmpoexpForm struct {
	CadreID string `json:"user_id" binding:"required"`
}

type ComfirmResumeForm struct {
	ID int `json:"id" binding:"required"`
}

type ComfirmfamilymemberForm struct {
	ID int `json:"id" binding:"required"`
	// 这里可以根据实际需求添加其他必要的字段
}

func ConfirmAssessment(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmAssessmentForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	AssessmentService := Assessment_service.ComfirmAssessment{
		Grade: form.Grade,
		ID:    form.ID,
	}

	if err := AssessmentService.ComfirmAssessment(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ASSESSMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetCadreInfo_mod(c *gin.Context) {
	appG := app.Gin{C: c}
	userId := c.Query("user_id")

	if userId == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	Cadreservice := Cadre_service.Cadre{ID: userId}
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
	cadreService := Cadre_service.Cadre{ID: form.CadreID}

	if err := cadreService.ComfirmCadre(); err != nil {
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
	confirmService := Positionhistory_service.Positionhistory{
		CadreID: form.CadreID,
	}

	// 调用确认逻辑
	if err := confirmService.ConfirmPositionHistory(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CONFIRM_POSITIONHISTORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetCadreInfoModByPage(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	department := c.Query("department")
	id := c.Query("user_id")
	gender := c.Query("gender")
	auditedStr := c.Query("audited")

	audited := 0
	if auditedStr != "" {
		if val, err := strconv.Atoi(auditedStr); err == nil {
			audited = val
		}
	}

	cadreInfoService := Admin_service.GetCadreInfoModByPage{
		Name:       name,
		Department: department,
		Gender:     gender,
		ID:         id,
		Audited:    audited,
		PageNum:    utils.GetPage(c),
		PageSize:   setting.AppSetting.PageSize,
	}
	cadreInfos, err := cadreInfoService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CADREINFO_FAIL, nil)
		return
	}

	count, err := cadreInfoService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CADREINFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": cadreInfos,
		"total": count,
	})
}

func GetAssessmentsMod(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	cadreID := c.Query("user_id")
	department := c.Query("department")
	category := c.Query("category")
	assessDept := c.Query("assess_dept")
	yearStr := c.Query("year")
	departmentIDStr := c.Query("department_id")
	departmentID := 0
	if departmentIDStr != "" {
		departmentID = com.StrTo(departmentIDStr).MustInt()
	}
	year := 0
	if yearStr != "" {
		year = com.StrTo(yearStr).MustInt()
	}
	auditedStr := c.Query("audited")
	audited := 0
	if auditedStr != "" {
		if val, err := strconv.Atoi(auditedStr); err == nil {
			audited = val
		}
	}

	assessmentService := Assessment_service.Assessment{
		Name:         name,
		UserID:       cadreID,
		Department:   department,
		Category:     category,
		AssessDept:   assessDept,
		Year:         year,
		Audited:      audited,
		PageNum:      utils.GetPage(c),
		PageSize:     setting.AppSetting.PageSize,
		DepartmentID: departmentID,
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

func GetAssessmentByID(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	assesementService := Assessment_service.Assessment{ID: id}
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

func GetpoexpmodByID(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	posexpModService := Positionhistory_service.Posexp{
		ID: id,
	}
	exists, err := posexpModService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POEXPMOD_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusNotFound, e.ERROR_GET_POEXPMOD_FAIL, nil)
		return
	}

	posexpMod, err := posexpModService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POEXPMOD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, posexpMod)
}

func GetPoexpModByPosID(c *gin.Context) {
	appG := app.Gin{C: c}
	userID := c.Query("user_id")

	if userID == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	poexpModService := Positionhistory_service.Posexp{
		CadreID: userID,
	}

	poexpMods, err := poexpModService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POEXPMOD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, poexpMods)
}

func Comfirmpoexp(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmpoexpForm
	)

	// 参数绑定和校验
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 调用服务层的确认逻辑
	poexpService := Admin_service.Comfirmpoexp{
		CadreID: form.CadreID,
	}
	if err := poexpService.Comfirmpoexp(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CONFIRM_POSITIONHISTORY_FAIL, nil)
		return
	}

	// 返回成功响应
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetPositionHistoriesMod(c *gin.Context) {
	appG := app.Gin{C: c}
	department := c.Query("department")
	category := c.Query("category")
	office := c.Query("office")
	academicYear := c.Query("academic_year")
	auditedStr := c.Query("audited")
	departmentIDStr := c.Query("department_id")
	departmentID := 0
	if departmentIDStr != "" {
		departmentID = com.StrTo(departmentIDStr).MustInt()
	}
	audited := 0
	if auditedStr != "" {
		if val, err := strconv.Atoi(auditedStr); err == nil {
			audited = val
		}
	}

	positionHistoryService := Positionhistory_service.Positionhistory{
		Department:   department,
		Category:     category,
		Office:       office,
		AcademicYear: academicYear,
		Audited:      audited,
		PageNum:      utils.GetPage(c),
		PageSize:     setting.AppSetting.PageSize,
		DepartmentID: departmentID,
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

	positionHistoryModService := Positionhistory_service.Positionhistory{
		ID: id,
	}
	positionHistoryMod, err := positionHistoryModService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POSITION_HISTORIES_FAIL, nil)
		return
	}

	if positionHistoryMod == nil {
		appG.Response(http.StatusNotFound, e.ERROR_GET_POSITION_HISTORIES_FAIL, nil)
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

	familyMemberModificationsService := Familymember_service.FamilyMember{ID: id}
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

	familyMemberModificationsService := Familymember_service.FamilyMember{
		CadreID: cadreID,
	}
	familyMemberModifications, err := familyMemberModificationsService.GetByCadreID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, familyMemberModifications)
}

func DeletePosexpmodbyID(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	posexpService := Positionhistory_service.Posexp{
		ID: id,
	}
	if err := posexpService.DeleteByID(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_POSEXP_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeletePosexpByID(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	posexpService := Positionhistory_service.Posexp{
		ID: id,
	}
	if err := posexpService.DeleteByID(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_POSEXP_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteAssessment(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	assessmentService := Assessment_service.Assessment{
		ID: id,
	}
	if err := assessmentService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ASSESSMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteAssessmentbyID(c *gin.Context) {
	appG := app.Gin{C: c}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	assessmentService := Assessment_service.Assessment{
		ID: id,
	}
	if err := assessmentService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ASSESSMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteFamilyMemberByID(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID 必须大于 0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	familyMemberService := Familymember_service.FamilyMember{ID: id}
	err := familyMemberService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_FAMILYMEMBERIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeletePositionhistorybyID(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID 必须大于 0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	positionHistoryService := Admin_service.PositionHistoryDelete{ID: id}
	err := positionHistoryService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_POSITION_HISTORIES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteCadreInfoByID(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := c.Param("id")
	valid.Required(id, "id").Message("ID 不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	cadreInfoService := Admin_service.CadreInfoDelete{ID: id}
	err := cadreInfoService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CADREINFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteResumeEntryByID(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID 必须大于 0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	resumeEntryService := Resume_service.ResumeEntry{ID: id}
	err := resumeEntryService.DeleteByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESUME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func ComfirmResume(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmResumeForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	ResumeService := Resume_service.ComfirmResume{
		ID: form.ID,
	}

	if err := ResumeService.ComfirmResume(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RESUME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Comfirmfamilymember(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ComfirmfamilymemberForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	FamilyMemberService := Familymember_service.FamilyMember{
		ID: form.ID,
	}

	if err := FamilyMemberService.Comfirmfamilymember(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAMILYMEBER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func SendMessage(c *gin.Context) {
	appG := app.Gin{C: c}

	// 绑定请求参数
	var form struct {
		RecipientID string `json:"recipient_id" binding:"required"`
		Message     string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 发送消息
	err := message_service.SendMessage(form.RecipientID, form.Message)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 发送成功响应
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "消息发送成功",
	})
}
