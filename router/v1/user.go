package v1

import (
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/pkg/utils"
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
	tokens, err := userService.Login(loginForm.UserID, loginForm.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"user":          gin.H{"user_id": loginForm.UserID},
	})
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	var registerForm struct {
		UserID   string `json:"id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userService := User_service.User{
		UserID:   registerForm.UserID,
		Name:     registerForm.Name,
		Password: registerForm.Password,
	}

	if err := userService.RegistUser(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokens, _ := userService.Login(registerForm.UserID, registerForm.Password)
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "注册成功",
		"user": gin.H{
			"id":   registerForm.UserID,
			"name": registerForm.Name,
			"role": "cadre",
		},
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func GetUserID(c *gin.Context) {
	appG := app.Gin{C: c}

	// 从上下文获取 claims
	claims, exists := c.Get("claims")
	if !exists {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}

	// 断言 claims 类型
	jwtClaims, ok := claims.(*utils.Claims)
	if !ok {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}

	// 获取 user_id
	userID := jwtClaims.UserID

	// 返回响应
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"user_id": userID,
	})
}

func GetUserRole(c *gin.Context) {
	appG := app.Gin{C: c}

	// 从上下文获取 claims
	claims, exists := c.Get("claims")
	if !exists {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}

	// 断言 claims 类型
	jwtClaims, ok := claims.(*utils.Claims)
	if !ok {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_CHECK_TOKEN_FAIL, nil)
		return
	}

	// 获取 role
	role := jwtClaims.Role

	// 返回响应
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"role": role,
	})
}
