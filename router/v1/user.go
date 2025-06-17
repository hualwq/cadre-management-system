package v1

import (
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/pkg/utils"
	"cadre-management/services/user_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	// 1. 绑定请求参数
	var loginForm struct {
		UserID   string `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 2. 检查 Authorization 头
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.ParseToken(tokenString)

		// 如果 token 解析成功并且 user_id 匹配
		if err == nil && claims.UserID == loginForm.UserID {
			// 获取角色
			userService := user_service.User{}
			userrole, roleErr := userService.GetRole(loginForm.UserID)
			if roleErr != nil {
				appG.Response(http.StatusBadRequest, e.ERROR_GET_ROLE, nil)
				return
			}

			appG.Response(http.StatusOK, e.SUCCESS, gin.H{
				"token": tokenString,
				"user":  gin.H{"user_id": loginForm.UserID},
				"role":  userrole,
			})
			return
		}

		// Token 无效或过期，尝试刷新
		userService := user_service.User{}
		newToken, refreshErr := userService.RefreshToken(tokenString)
		if refreshErr == nil {
			userrole, roleErr := userService.GetRole(loginForm.UserID)
			if roleErr != nil {
				appG.Response(http.StatusBadRequest, e.ERROR_GET_ROLE, nil)
				return
			}
			appG.Response(http.StatusOK, e.SUCCESS, gin.H{
				"token": newToken,
				"user":  gin.H{"user_id": loginForm.UserID},
				"role":  userrole,
			})
			return
		}
	}

	// 5. 正常登录流程
	userService := user_service.User{}
	token, err := userService.Login(loginForm.UserID, loginForm.Password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	// 获取角色
	userrole, roleErr := userService.GetRole(loginForm.UserID)
	if roleErr != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_GET_ROLE, nil)
		return
	}

	// 6. 返回新 Token
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"token": token,
		"user":  gin.H{"user_id": loginForm.UserID},
		"role":  userrole,
	})
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	// 定义接收前端数据的结构体
	var registerForm struct {
		UserID   string `json:"id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 绑定并验证参数
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 调用 service 层注册函数
	userService := user_service.User{
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

	// 注册成功响应
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "注册成功",
		"user": gin.H{
			"id":   registerForm.UserID,
			"name": registerForm.Name,
			"role": "cadre", // 默认角色
		},
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
