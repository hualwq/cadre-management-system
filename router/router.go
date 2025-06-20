package router

import (
	"cadre-management/middleware"
	"cadre-management/pkg/upload"
	v1 "cadre-management/router/v1"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	v1.InitDepartmentService()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // 你前端的地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 如果你设置了 cookie 或需要身份凭证
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/login", v1.Login)
	r.POST("/register", v1.Register)
	r.POST("/refresh-token", v1.RefreshToken)
	r.GET("/departments", v1.GetDepartments)
	apiv1 := r.Group("/cadre")
	apiv1.Use(middleware.JWT()) //普通干部
	{
		apiv1.GET("/getuserid", v1.GetUserID)
		apiv1.POST("/cadreinfo", v1.AddCadreInfo)
		apiv1.POST("/assessment", v1.AddAssessment)
		apiv1.POST("/positionhistory", v1.AddPositionHistory)
		apiv1.POST("/yearposition", v1.Addyearposition)
		apiv1.POST("/resume", v1.AddResume)
		apiv1.POST("/familymember", v1.Addfamilymember)
		apiv1.POST("/image", v1.UploadImage)

		apiv1.GET("/getphmodbypage", v1.GetPositionHistoryModsByPage)
		apiv1.GET("/getasmodbypage", v1.GetAssessmentByPage)
		apiv1.GET("/getposexpbyposid", v1.GetPosExpByPosID)

		apiv1.PUT("/cadreinfo", v1.EditInfo)
		apiv1.PUT("/resume", v1.EditResume)
		apiv1.PUT("/positionhistory", v1.EditPh)
		apiv1.PUT("/familymember", v1.Editfamilymember)
		apiv1.PUT("/assessment", v1.EditAssessment)

		apiv1.DELETE("/familymember", v1.DeleteFamilyMemberModification)
		apiv1.DELETE("/resume", v1.DeleteFamilyMemberModification)
		apiv1.DELETE("/positionhistory", v1.Deletephmod)
		apiv1.DELETE("/cadre", v1.Deletecadremod)
		apiv1.DELETE("/posexp", v1.DeletePosexpmodbyID)
		apiv1.DELETE("/assessment", v1.DeleteAssessment)

		apiv1.GET("/getmessage", v1.GetCadreMessages)
	}
	apiv2 := r.Group("/admin")
	apiv2.Use(middleware.JWT()) // 管理员
	{
		apiv2.GET("/assessmentbypage", v1.GetAssessmentsMod)
		apiv2.GET("/phmodbypage", v1.GetPositionHistoriesMod)
		apiv2.GET("/phbypage", v1.GetPositionHistories)
		apiv2.GET("/cadrebypage", v1.GetCadreInfoModByPage)

		apiv2.GET("/assmodbyid", v1.GetAssessmentByID)
		apiv2.GET("/phmodbyid", v1.GetPositionHistoryMod)
		apiv2.GET("/fmmodbyid", v1.GetFamilyMemberModification)
		apiv2.GET("/resumebyid", v1.GetResumeEntryModificationByID)
		apiv2.GET("/poexpbyid", v1.GetpoexpmodByID)

		apiv2.GET("/cadreinfo", v1.GetCadreInfo_mod)
		apiv2.GET("/fammonbycadreid", v1.GetFamilyMemberModificationsByCadreID)
		apiv2.GET("/resmonbycadreid", v1.GetResumeEntryModificationsByCadreID)
		apiv2.GET("/poexpbycadreid", v1.GetPoexpModByPosID)

		apiv2.POST("/cadreinfo", v1.ComfirmCadreInfo)
		apiv2.POST("/assessment", v1.ConfirmAssessment)
		apiv2.POST("/positionhistory", v1.ConfirmPositionhistory)
		apiv2.POST("/poexp", v1.Comfirmpoexp)
		apiv2.POST("/resume", v1.ComfirmResume)
		apiv2.POST("/familymember", v1.Comfirmfamilymember)
		apiv2.POST("/sendmessage", v1.SendMessage)

		apiv2.DELETE("/assessment", v1.DeleteAssessmentbyID)
		apiv2.DELETE("/poexp", v1.DeletePosexpByID)
		apiv2.DELETE("/familymember", v1.DeleteFamilyMemberByID)
		apiv2.DELETE("/positionhistory", v1.DeletePositionhistorybyID)
		apiv2.DELETE("/cadreinfo", v1.DeleteCadreInfoByID)
		apiv2.DELETE("/resume", v1.DeleteResumeEntryByID)
	}
	apiv3 := r.Group("/sysadmin")
	apiv3.Use(middleware.JWT())
	{
		apiv3.GET("/userbypage", v1.GetUserByPage)
		apiv3.GET("/alluser", v1.GetAllUser)

		apiv3.POST("/changerole", v1.ChangeUserRole)

		// department management (school_admin)
		apiv3.POST("/department", v1.CreateDepartment)
		apiv3.DELETE("/department/:id", v1.DeleteDepartment)
		apiv3.PUT("/department/:id", v1.UpdateDepartment)
		apiv3.GET("/departments", v1.ListDepartments)
		apiv3.GET("/department/:id", v1.GetDepartmentByID)
		apiv3.POST("/department/admin", v1.SetDepartmentAdmin)
		apiv3.POST("/department/admin/unset", v1.UnsetDepartmentAdmin)
		apiv3.GET("/department/admins", v1.GetDepartmentAdmins)
	}

	return r
}
