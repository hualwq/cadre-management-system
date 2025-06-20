package v1

import (
	"cadre-management/models"
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/services/User_service"
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
				"user":          gin.H{"user_id": loginForm.UserID},
			})
			return
		}
	}

	userService := User_service.User{}
	tokens, role, err := userService.Login(loginForm.UserID, loginForm.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"role":          role,
		"user":          gin.H{"user_id": loginForm.UserID},
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

	tokens, role, _ := userService.Login(registerForm.UserID, registerForm.Password)
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "注册成功",
		"user": gin.H{
			"id":            registerForm.UserID,
			"name":          registerForm.Name,
			"department_id": registerForm.DepartmentID,
		},
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"role":          role,
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
