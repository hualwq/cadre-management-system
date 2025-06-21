package v1

import (
	"cadre-management/models"
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/services/User_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	var loginForm struct {
		UserID   string `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	refreshToken := c.GetHeader("X-Refresh-Token")
	if refreshToken != "" {
		userService := User_service.User{}
		tokens, err := userService.RefreshToken(refreshToken)
		if err == nil {
			appG.Response(http.StatusOK, e.SUCCESS, gin.H{
				"access_token":  tokens.AccessToken,
				"refresh_token": tokens.RefreshToken,
				"user": gin.H{
					"user_id": loginForm.UserID,
				},
			})
			return
		}
	}

	userService := User_service.User{}
	tokens, loginResult, err := userService.Login(loginForm.UserID, loginForm.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"user_id":       loginResult.UserID,
		"role":          loginResult.Role,
		"department_id": loginResult.DepartmentID,
		"name":          loginResult.Name,
	})
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	var registerForm struct {
		UserID       string `json:"id" binding:"required"`
		Name         string `json:"name" binding:"required"`
		Password     string `json:"password" binding:"required"`
		DepartmentID uint   `json:"department_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userService := User_service.User{
		UserID:       registerForm.UserID,
		Name:         registerForm.Name,
		Password:     registerForm.Password,
		DepartmentID: registerForm.DepartmentID,
	}

	if err := userService.RegistUser(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokens, loginResult, _ := userService.Login(registerForm.UserID, registerForm.Password)
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message":       "注册成功",
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"user_id":       loginResult.UserID,
		"role":          loginResult.Role,
		"department_id": loginResult.DepartmentID,
		"name":          loginResult.Name,
	})
}

func GetUserID(c *gin.Context) {
	appG := app.Gin{C: c}

	// 从上下文获取 claims
	userID, exists := c.Get("user_id")
	if !exists {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}

	// 返回响应
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"user_id": userID,
	})
}

func RefreshToken(c *gin.Context) {
	appG := app.Gin{C: c}

	var refreshForm struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&refreshForm); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userService := User_service.User{}
	tokens, err := userService.RefreshToken(refreshForm.RefreshToken)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func GetDepartments(c *gin.Context) {
	appG := app.Gin{C: c}

	departments, err := models.GetAllDepartments()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"departments": departments,
	})
}

// GetUsers 获取用户列表，支持分页和多条件筛选
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	page := 1
	pageSize := 10
	if p := c.DefaultQuery("page", "1"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.DefaultQuery("page_size", "10"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}

	userID := c.Query("userId")
	name := c.Query("name")
	role := c.Query("role")
	departmentID := uint(0)
	if did := c.Query("departmentId"); did != "" {
		var tmp uint
		fmt.Sscanf(did, "%d", &tmp)
		departmentID = tmp
	}

	userService := User_service.User{}
	users, total, err := userService.GetUsersWithFilter(page, pageSize, userID, name, role, departmentID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERLIST_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"users": users,
		"total": total,
	})
}

func GetUserByID(c *gin.Context) {
	appG := app.Gin{C: c}
	userID := c.Param("id")
	user, err := models.GetUserByID(userID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"user": user,
	})
}

// ExistsCadreInfo 查询干部信息是否存在
func ExistsCadreInfo(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "user_id参数不能为空"})
		return
	}
	exists, err := models.ExistCadreInfo(userID)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "查询失败", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "exists": exists})
}

// ExistsResume 查询简历是否存在
func ExistsResume(c *gin.Context) {
	idStr := c.Query("id")
	var id uint
	fmt.Sscanf(idStr, "%d", &id)
	if id == 0 {
		c.JSON(400, gin.H{"code": 400, "msg": "id参数不能为空"})
		return
	}
	exists, err := models.ExistResume(id)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "查询失败", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "exists": exists})
}

// ExistsFamilyMember 查询家庭成员是否存在
func ExistsFamilyMember(c *gin.Context) {
	idStr := c.Query("id")
	var id uint
	fmt.Sscanf(idStr, "%d", &id)
	if id == 0 {
		c.JSON(400, gin.H{"code": 400, "msg": "id参数不能为空"})
		return
	}
	exists, err := models.ExistFamilyMember(id)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "查询失败", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "exists": exists})
}

// ExistsPositionHistory 查询岗位历史是否存在
func ExistsPositionHistory(c *gin.Context) {
	idStr := c.Query("id")
	var id uint
	fmt.Sscanf(idStr, "%d", &id)
	if id == 0 {
		c.JSON(400, gin.H{"code": 400, "msg": "id参数不能为空"})
		return
	}
	exists, err := models.ExistPositionHistory(id)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "查询失败", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "exists": exists})
}

// ExistsAssessment 查询考核是否存在
func ExistsAssessment(c *gin.Context) {
	idStr := c.Query("id")
	var id uint
	fmt.Sscanf(idStr, "%d", &id)
	if id == 0 {
		c.JSON(400, gin.H{"code": 400, "msg": "id参数不能为空"})
		return
	}
	exists, err := models.ExistAssessment(id)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "查询失败", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "exists": exists})
}

func GetUserDepartment(c *gin.Context) {
	appG := app.Gin{C: c}
	userID, exists := c.Get("user_id")
	if !exists {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}
	user, err := models.GetUserDepartment(userID.(string))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"user": user,
	})
}
