// middleware/jwt.go
package middleware

import (
    "net/http"
    "strings"
    
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"

    "cadre-management/pkg/utils"
    "cadre-management/pkg/e"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code = e.SUCCESS
        var data interface{}

        // 1. 从 Header 获取 Access Token
        token := ""
        authHeader := c.GetHeader("Authorization")
        if authHeader != "" {
            token = strings.TrimPrefix(authHeader, "Bearer ")
        } else {
            token = c.Query("token")
        }

        if token == "" {
            code = e.INVALID_PARAMS
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": code,
                "msg":  e.GetMsg(code),
                "data": data,
            })
            c.Abort()
            return
        }

        // 2. 解析 Access Token
        claims, err := utils.ParseToken(token)
        if err != nil {
            if ve, ok := err.(*jwt.ValidationError); ok {
                if ve.Errors&jwt.ValidationErrorExpired != 0 {
                    code = e.ERROR_USER_CHECK_TOKEN_TIMEOUT
                    c.JSON(http.StatusUnauthorized, gin.H{
                        "code": code,
                        "msg":  "AccessToken已过期，请用RefreshToken刷新",
                        "data": data,
                    })
                    c.Abort()
                    return
                } else {
                    code = e.ERROR_USER_CHECK_TOKEN_FAIL
                }
            } else {
                code = e.ERROR_USER_CHECK_TOKEN_FAIL
            }
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": code,
                "msg":  e.GetMsg(code),
                "data": data,
            })
            c.Abort()
            return
        }

        // 3. 从 claims 提取 user_id 并存入上下文
        if claims.UserID != "" {
            c.Set("user_id", claims.UserID)
        } else {
            code = e.ERROR_USER_CHECK_TOKEN_FAIL
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": code,
                "msg":  "Token 中缺少 user_id",
                "data": data,
            })
            c.Abort()
            return
        }

        // 4. 存储 claims 供其他可能用途
        c.Set("claims", claims)
        c.Next()
    }
}
