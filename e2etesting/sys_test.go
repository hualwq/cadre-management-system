package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestUserRegister(t *testing.T) {
	r := router.InitRouter()

	// 准备注册数据
	registerData := map[string]string{
		"id":       "wangwu",
		"name":     "王五",
		"password": "123456",
	}
	jsonData, err := json.Marshal(registerData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
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

func TestChangeUserRole(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYzA2M2E5YTQwNDBmNzQxMmVmODgxOTc1YmRlMjA4MWMiLCJwYXNzd29yZCI6IjdhNDM0ZjFlZjA2MDFiYWZmOGUyNTM4YmFmNzdlZmNhIiwicm9sZSI6InN5c2FkbWluIiwiZXhwIjoxNzQ5MTc4ODI3fQ.CYt6PsRZ8rOHHecMUYusnqOFqv4PVnwzprCBusyBZ4s"

	// 准备注册数据
	registerData := map[string]string{
		"user_id": "wangwu",
		"role":    "admin",
	}
	jsonData, err := json.Marshal(registerData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/sysadmin/changerole", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
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

func TestGetAllUser(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYzA2M2E5YTQwNDBmNzQxMmVmODgxOTc1YmRlMjA4MWMiLCJwYXNzd29yZCI6IjdhNDM0ZjFlZjA2MDFiYWZmOGUyNTM4YmFmNzdlZmNhIiwicm9sZSI6InN5c2FkbWluIiwiZXhwIjoxNzQ5MTc4ODI3fQ.CYt6PsRZ8rOHHecMUYusnqOFqv4PVnwzprCBusyBZ4s"

	// 创建一个HTTP请求
	req, err := http.NewRequest("GET", "/sysadmin/alluser", nil)
	req.Header.Set("Authorization", "Bearer "+token)
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

func TestGetUserByPage(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYzA2M2E5YTQwNDBmNzQxMmVmODgxOTc1YmRlMjA4MWMiLCJwYXNzd29yZCI6IjdhNDM0ZjFlZjA2MDFiYWZmOGUyNTM4YmFmNzdlZmNhIiwicm9sZSI6InN5c2FkbWluIiwiZXhwIjoxNzQ5MTc4ODI3fQ.CYt6PsRZ8rOHHecMUYusnqOFqv4PVnwzprCBusyBZ4s"

	page := 1
	pagesize := 3
	baseurl := "/sysadmin/userbypage?"
	queryParams := fmt.Sprintf("page=%d&pagesize=%d", page, pagesize)
	fullurl := baseurl + queryParams
	// 创建一个HTTP请求
	req, err := http.NewRequest("GET", fullurl, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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

func TestPOSTCadreInfo(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWJiZDg4NjQ2MDgyNzAxNWU1ZDYwNWVkNDQyNTIyNTEiLCJwYXNzd29yZCI6ImY0ZGEzNTE0M2I1YWM5NDQyMGRiMzZjM2VmOTdiNzY4Iiwicm9sZSI6ImNhZHJlIiwiZXhwIjoxNzQ5MTg1MjQ0fQ.cnurwurqLHeScJjpBJrEad4Giu873JmIe5j9xF_x6KA"

	CadreData := map[string]string{
		"user_id":                     "11111111",
		"name":                        "张三",
		"gender":                      "男",
		"birth_date":                  "1989.2",
		"ethnic_group":                "汉族",
		"native_place":                "广东省",
		"birth_place":                 "广东省",
		"political_status":            "中共党员",
		"work_start_date":             "2013.7",
		"health_status":               "良好",
		"professional_title":          "工程师",
		"specialty":                   "软件开发",
		"phone":                       "13800138000",
		"current_position":            "软件研发部经理",
		"awards_and_punishments":      "2018年获得公司优秀员工称号",
		"annual_assessment":           "近三年考核结果均为优秀",
		"email":                       "zhangsan@example.com",
		"filled_by":                   "李四",
		"full_time_education_degree":  "本科",
		"full_time_education_school":  "XX大学计算机科学与技术专业",
		"on_the_job_education_degree": "硕士",
		"on_the_job_education_school": "YY大学软件工程专业",
		"reporting_unit":              "ZZ公司人力资源部",
		"approval_authority":          "同意晋升",
		"administrative_appointment":  "任命为软件研发部经理",
	}
	jsonData, err := json.Marshal(CadreData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/cadre/cadreinfo", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
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
