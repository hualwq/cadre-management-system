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
	r.GET("/getuserrole", v1.GetUserRole)
	apiv1 := r.Group("/cadre")
	apiv1.Use(middleware.JWT()) //普通干部
	{

		apiv1.GET("/getuserid", v1.GetUserID)
		apiv1.POST("/cadreinfo", middleware.RoleMiddleware("cadre"), v1.AddCadreInfo)
		apiv1.POST("/assessment", middleware.RoleMiddleware("cadre"), v1.AddAssessment)
		apiv1.POST("/positionhistory", middleware.RoleMiddleware("cadre"), v1.AddPositionHistory)
		apiv1.POST("/yearposition", middleware.RoleMiddleware("cadre"), v1.Addyearposition)
		apiv1.POST("/resume", middleware.RoleMiddleware("cadre"), v1.AddResume)
		apiv1.POST("/familymember", middleware.RoleMiddleware("cadre"), v1.Addfamilymember)
		apiv1.POST("/image", middleware.RoleMiddleware("cadre"), v1.UploadImage)

		apiv1.GET("/getphmodbypage", middleware.RoleMiddleware("cadre"), v1.GetPositionHistoryModsByPage)
		apiv1.GET("/getasmodbypage", middleware.RoleMiddleware("cadre"), v1.GetAssessmentByPage)
		apiv1.GET("/getposexpbyposid", middleware.RoleMiddleware("cadre"), v1.GetPosExpByPosID)

		apiv1.PUT("/cadreinfo", middleware.RoleMiddleware("cadre"), v1.EditInfo)
		apiv1.PUT("/resume", middleware.RoleMiddleware("cadre"), v1.EditResume)
		apiv1.PUT("/positionhistory", middleware.RoleMiddleware("cadre"), v1.EditPh)
		apiv1.PUT("/familymember", middleware.RoleMiddleware("cadre"), v1.Editfamilymember)
		apiv1.PUT("/assessment", middleware.RoleMiddleware("cadre"), v1.EditAssessment)

		apiv1.DELETE("/familymember", middleware.RoleMiddleware("cadre"), v1.DeleteFamilyMemberModification)
		apiv1.DELETE("/resume", middleware.RoleMiddleware("cadre"), v1.DeleteFamilyMemberModification)
		apiv1.DELETE("/positionhistory", middleware.RoleMiddleware("cadre"), v1.Deletephmod)
		apiv1.DELETE("/cadre", middleware.RoleMiddleware("cadre"), v1.Deletecadremod)
		apiv1.DELETE("/posexp", middleware.RoleMiddleware("cadre"), v1.DeletePosexpmodbyID)
		apiv1.DELETE("/assessment", middleware.RoleMiddleware("cadre"), v1.DeleteAssessment)

		apiv1.GET("/getmessage", middleware.RoleMiddleware("cadre"), v1.GetCadreMessages)
	}
	apiv2 := r.Group("/admin")
	apiv2.Use(middleware.JWT()) // 管理员
	{
		apiv2.GET("/assessmentbypage", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetAssessmentsMod)
		apiv2.GET("/phmodbypage", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetPositionHistoriesMod)
		apiv2.GET("/phbypage", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetPositionHistories)
		apiv2.GET("/cadrebypage", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetCadreInfoModByPage)

		apiv2.GET("/assmodbyid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetAssessmentByID)
		apiv2.GET("/phmodbyid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetPositionHistoryMod)
		apiv2.GET("/fmmodbyid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetFamilyMemberModification)
		apiv2.GET("/resumebyid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetResumeEntryModificationByID)
		apiv2.GET("/poexpbyid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetpoexpmodByID)

		apiv2.GET("/cadreinfo", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetCadreInfo_mod)
		apiv2.GET("/fammonbycadreid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetFamilyMemberModificationsByCadreID)
		apiv2.GET("/resmonbycadreid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetResumeEntryModificationsByCadreID)
		apiv2.GET("/poexpbycadreid", middleware.RoleMiddleware("admin", "sysadmin", "cadre"), v1.GetPoexpModByPosID)

		apiv2.POST("/cadreinfo", middleware.RoleMiddleware("admin", "sysadmin"), v1.ComfirmCadreInfo)
		apiv2.POST("/assessment", middleware.RoleMiddleware("admin", "sysadmin"), v1.ConfirmAssessment)
		apiv2.POST("/positionhistory", middleware.RoleMiddleware("admin", "sysadmin"), v1.ConfirmPositionhistory)
		apiv2.POST("/poexp", middleware.RoleMiddleware("admin", "sysadmin"), v1.Comfirmpoexp)
		apiv2.POST("/resume", middleware.RoleMiddleware("admin", "sysadmin"), v1.ComfirmResume)
		apiv2.POST("/familymember", middleware.RoleMiddleware("admin", "sysadmin"), v1.Comfirmfamilymember)
		apiv2.POST("/sendmessage", middleware.RoleMiddleware("admin", "sysadmin"), v1.SendMessage)

		apiv2.DELETE("/assessment", middleware.RoleMiddleware("admin", "sysadmin"), v1.DeleteAssessmentbyID)
		apiv2.DELETE("/poexp", middleware.RoleMiddleware("admin", "sysadmin"), v1.DeletePosexpByID)
		apiv2.DELETE("/familymember", middleware.RoleMiddleware("admin", "sysadmin"), v1.DeleteFamilyMemberByID)
		apiv2.DELETE("/positionhistory", middleware.RoleMiddleware("admin", "sysadmin"), v1.DeletePositionhistorybyID)
		apiv2.DELETE("/cadreinfo", middleware.RoleMiddleware("admin", "sysadmin"), v1.DeleteCadreInfoByID)
		apiv2.DELETE("/resume", middleware.RoleMiddleware("admin", "sysadmin"), v1.DeleteResumeEntryByID)
	}
	apiv3 := r.Group("/sysadmin")
	apiv3.Use(middleware.JWT())
	{
		apiv3.GET("/userbypage", middleware.RoleMiddleware("sysadmin"), v1.GetUserByPage)
		apiv3.GET("/alluser", middleware.RoleMiddleware("sysadmin"), v1.GetAllUser)

		apiv3.POST("/changerole", middleware.RoleMiddleware("sysadmin"), v1.ChangeUserRole)

		// department management (school_admin)
		apiv3.POST("/department", middleware.RoleMiddleware("school_admin"), v1.CreateDepartment)
		apiv3.DELETE("/department/:id", middleware.RoleMiddleware("school_admin"), v1.DeleteDepartment)
		apiv3.PUT("/department/:id", middleware.RoleMiddleware("school_admin"), v1.UpdateDepartment)
		apiv3.GET("/departments", middleware.RoleMiddleware("school_admin"), v1.ListDepartments)
		apiv3.POST("/department/admin", middleware.RoleMiddleware("school_admin"), v1.SetDepartmentAdmin)
		apiv3.POST("/department/admin/unset", middleware.RoleMiddleware("school_admin"), v1.UnsetDepartmentAdmin)
		apiv3.GET("/department/admins", middleware.RoleMiddleware("school_admin"), v1.GetDepartmentAdmins)
	}

	return r
}
