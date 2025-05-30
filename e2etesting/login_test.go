package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"cadre-management/models"
	"cadre-management/pkg/logging"
	"cadre-management/pkg/setting"
	"cadre-management/pkg/utils"
	"cadre-management/router"

	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	setting.Setup()
	models.Setup()
	logging.Setup()
	utils.Setup()
}

func TestUserLogin(t *testing.T) {
	// 注册路由并获取返回的 Gin 引擎实例
	r := router.InitRouter()

	// 准备登录数据
	loginData := map[string]string{
		"id":       "11111111",
		"password": "123456",
	}
	jsonData, err := json.Marshal(loginData)
	assert.NoError(t, err)

	// 创建一个 HTTP 请求
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
