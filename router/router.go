package router

import (
	"cadre-management/middleware"
	"cadre-management/pkg/upload"
	v1 "cadre-management/router/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/login", v1.Login)
	r.POST("/register", v1.Register)
	apiv1 := r.Group("/cadre")
	apiv1.Use(middleware.JWT()) //普通干部
	{
		apiv1.POST("/info", v1.AddCadreInfo_mod)
		apiv1.POST("/assessment", v1.AddAssessment_mod)
		apiv1.POST("/positionhistory", v1.AddPositionHistory_mod)
		apiv1.POST("/yearposition", v1.Addyearposition_mod)
		apiv1.POST("/resume", v1.AddResume_mod)
		apiv1.POST("/familymember", v1.Addfamilymember_mod)
		apiv1.POST("/image", v1.UploadImage)

		apiv1.PUT("/cadre", v1.EditInfo_mod)
		apiv1.PUT("/resume", v1.EditResume_Mod)
		apiv1.PUT("/positionhistory", v1.EditPh_Mod)
		apiv1.PUT("/familymember", v1.Editfamilymember_mod)
		apiv1.PUT("/assessment", v1.EditassessmentMod)

		apiv1.DELETE("/familymember", v1.DeleteFamilyMemberModification)
		apiv1.DELETE("/resume", v1.DeleteFamilyMemberModification)
		apiv1.DELETE("/positionhistory", v1.Deletephmod)
		apiv1.DELETE("/cadre", v1.Deletecadremod)
		apiv1.DELETE("/posexp", v1.DeletePosexpmodbyID)
		apiv1.DELETE("/assessment", v1.DeleteAssessmentmodbyID)

	}
	apiv2 := r.Group("/admin")
	apiv2.Use(middleware.JWT()) //管理员
	{
		apiv2.GET("/assessmentbypage", v1.GetAssessmentsMod)
		apiv2.GET("/phmodbypage", v1.GetPositionHistoriesMod)
		apiv2.GET("/phbypage", v1.GetPositionHistories)

		apiv2.GET("/positionhistory", v1.GetPositionHistory_mod)
		apiv2.GET("/assmodbyid", v1.GetAssessmentsModByID)
		apiv2.GET("/phmodbyid", v1.GetPositionHistoryMod)
		apiv2.GET("/fmmodbyid", v1.GetFamilyMemberModification)
		apiv2.GET("/resumebyid", v1.GetResumeEntryModificationByID)
		apiv2.GET("/poexpbyid", v1.GetpoexpmodByID)

		apiv2.GET("/cadreinfo", v1.GetCadreInfo_mod)
		apiv2.GET("/fammonbycadreid", v1.GetFamilyMemberModificationsByCadreID)
		apiv2.GET("/resmonbycadreid", v1.GetResumeEntryModificationsByCadreID)
		apiv2.GET("/poexpbycadreid", v1.GetPoexpModByCadreID)

		apiv2.POST("/cadreinfo", v1.ComfirmCadreInfo)
		apiv2.POST("/assessment", v1.ComfirmAssessment)
		apiv2.POST("/positionhistory", v1.ConfirmPositionhistory)
		apiv2.POST("/poexp", v1.Comfirmpoexp)

		apiv2.DELETE("/assessment", v1.DeleteAssessmentbyID)
		apiv2.DELETE("/poexp", v1.DeletePosexpByID)
	}
	apiv3 := r.Group("/sysadmin")
	apiv3.Use(middleware.JWT())
	{
		apiv3.GET("/userbypage", v1.GetUserByPage)
		apiv3.GET("/alluser", v1.GetAllUser)
	}

	return r
}
