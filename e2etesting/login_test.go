package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cadre-management/pkg/setting"
	"cadre-management/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	// 初始化系统
	setting.Setup()

	// 创建一个新的Gin引擎
	r := gin.Default()

	// 注册路由
	router.InitRouter()

	// 准备登录数据
	loginData := map[string]string{
		"id":       "11111111",
		"password": "123456",
	}
	jsonData, err := json.Marshal(loginData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// 创建一个响应记录器
	rr := httptest.NewRecorder()

	// 执行请求
	r.ServeHTTP(rr, req)

	// 断言响应状态码
	assert.Equal(t, http.StatusOK, rr.Code)
	t.Logf("响应状态码: %d", rr.Code)
	t.Logf("响应内容: %s", rr.Body.String())
}
