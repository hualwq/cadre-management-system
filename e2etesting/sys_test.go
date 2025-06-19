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
		"id":       "wangwu",
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

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTExMTExMTExIiwicGFzc3dvcmQiOiI5OTYwZGQ4ODRmMDZlODM5NjllNzM2YjNlYWUwYTUwYiIsInJvbGUiOiJjYWRyZSIsImV4cCI6MTc1MDc0NzYzNn0.-CmBS4fHP6bk0N-hlmIjZUy1WEYiSI585_S99WcrKNI"

	CadreData := map[string]string{
		"user_id":                     "111111111",
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

func TestPUTCadreInfo(t *testing.T) {
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

func TestPOSTAssessment(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWJiZDg4NjQ2MDgyNzAxNWU1ZDYwNWVkNDQyNTIyNTEiLCJwYXNzd29yZCI6ImY0ZGEzNTE0M2I1YWM5NDQyMGRiMzZjM2VmOTdiNzY4Iiwicm9sZSI6ImNhZHJlIiwiZXhwIjoxNzQ5MTg1MjQ0fQ.cnurwurqLHeScJjpBJrEad4Giu873JmIe5j9xF_x6KA"

	CadreData := map[string]interface{}{
		"user_id":      "11111111",
		"department":   "软件学院",
		"category":     "专职团干部",
		"assess_dept":  "校团委,院系党委",
		"work_summary": "本年度我主要负责学院团委的日常工作，包括组织学生活动、开展思想教育工作等。具体工作包括：1. 策划并执行了5场大型学生活动；2. 组织了3次团干部培训；3. 完成了学院团委的日常管理工作。在工作中我注重团队协作，积极创新工作方法，取得了显著成效。",
		"year":         2020,
	}
	jsonData, err := json.Marshal(CadreData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/cadre/assessment", bytes.NewBuffer(jsonData))
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

func TestPOSTPosexp(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWJiZDg4NjQ2MDgyNzAxNWU1ZDYwNWVkNDQyNTIyNTEiLCJwYXNzd29yZCI6ImY0ZGEzNTE0M2I1YWM5NDQyMGRiMzZjM2VmOTdiNzY4Iiwicm9sZSI6ImNhZHJlIiwiZXhwIjoxNzQ5MTg1MjQ0fQ.cnurwurqLHeScJjpBJrEad4Giu873JmIe5j9xF_x6KA"

	CadreData := map[string]interface{}{
		"user_id":    "11111111",
		"year":       "2021",
		"department": "计算机学院",
		"position":   "未知职位",
	}
	jsonData, err := json.Marshal(CadreData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/cadre/yearposition", bytes.NewBuffer(jsonData))
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

func TestPOSTRESUME(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWJiZDg4NjQ2MDgyNzAxNWU1ZDYwNWVkNDQyNTIyNTEiLCJwYXNzd29yZCI6ImY0ZGEzNTE0M2I1YWM5NDQyMGRiMzZjM2VmOTdiNzY4Iiwicm9sZSI6ImNhZHJlIiwiZXhwIjoxNzQ5MTg1MjQ0fQ.cnurwurqLHeScJjpBJrEad4Giu873JmIe5j9xF_x6KA"

	CadreData := map[string]interface{}{
		"user_id":      "11111111",
		"start_date":   "2015.09",
		"end_date":     "2018.01",
		"organization": "××大学",
		"department":   "××学院××专业",
		"position":     "本科生",
	}
	jsonData, err := json.Marshal(CadreData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/cadre/resume", bytes.NewBuffer(jsonData))
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

func TestPOSTFAMILYMEMBER(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWJiZDg4NjQ2MDgyNzAxNWU1ZDYwNWVkNDQyNTIyNTEiLCJwYXNzd29yZCI6ImY0ZGEzNTE0M2I1YWM5NDQyMGRiMzZjM2VmOTdiNzY4Iiwicm9sZSI6ImNhZHJlIiwiZXhwIjoxNzQ5MTg1MjQ0fQ.cnurwurqLHeScJjpBJrEad4Giu873JmIe5j9xF_x6KA"

	CadreData := map[string]interface{}{
		"user_id":          "11111111",
		"relation":         "母亲",
		"name":             "张母",
		"birth_date":       "1962.1",
		"political_status": "群众",
		"work_unit":        "XX单位退休职工",
	}
	jsonData, err := json.Marshal(CadreData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/cadre/familymember", bytes.NewBuffer(jsonData))
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

func TestPOSTPositionhistory(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMWJiZDg4NjQ2MDgyNzAxNWU1ZDYwNWVkNDQyNTIyNTEiLCJwYXNzd29yZCI6ImY0ZGEzNTE0M2I1YWM5NDQyMGRiMzZjM2VmOTdiNzY4Iiwicm9sZSI6ImNhZHJlIiwiZXhwIjoxNzQ5MTg1MjQ0fQ.cnurwurqLHeScJjpBJrEad4Giu873JmIe5j9xF_x6KA"

	CadreData := map[string]interface{}{
		"user_id":          "11111111",
		"department":       "计算机学院",
		"category":         "专职团干部",
		"office":           "学生会",
		"academic_year":    "2023-2024第一学期",
		"applied_at_year":  2024,
		"applied_at_month": 6,
		"applied_at_day":   1,
	}
	jsonData, err := json.Marshal(CadreData)
	assert.NoError(t, err)

	// 创建一个HTTP请求
	req, err := http.NewRequest("POST", "/cadre/positionhistory", bytes.NewBuffer(jsonData))
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

func TestGETCadreinfoByPage(t *testing.T) { //测试第二页的时候有问题
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzUwNjY4MDUxfQ.GR4QGGXnTz9H6KzH0pnVicRauwGcuf0huvTbOOvCVsI"

	page := 2
	baseurl := "/admin/cadrebypage?"
	queryParams := fmt.Sprintf("page=%d", page)
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

func TestGETCadreinfoByID(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzUwNjY4MDUxfQ.GR4QGGXnTz9H6KzH0pnVicRauwGcuf0huvTbOOvCVsI"

	user_id := "111111111"
	baseurl := "/admin/cadreinfo?"
	queryParams := fmt.Sprintf("user_id=%s", user_id)
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

func TestGETPositionhistoryByPage(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzUwNjY4MDUxfQ.GR4QGGXnTz9H6KzH0pnVicRauwGcuf0huvTbOOvCVsI"

	page := 1
	baseurl := "/admin/phmodbypage?"
	queryParams := fmt.Sprintf("page=%d", page)
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

func TestGETAssessmentByPage(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzUwNjY4MDUxfQ.GR4QGGXnTz9H6KzH0pnVicRauwGcuf0huvTbOOvCVsI"

	page := 1
	baseurl := "/admin/assessmentbypage?"
	queryParams := fmt.Sprintf("page=%d", page)
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

func TestGETAssessmentByID(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzUwNjY4MDUxfQ.GR4QGGXnTz9H6KzH0pnVicRauwGcuf0huvTbOOvCVsI"

	id := 2
	baseurl := "/admin/assmodbyid?"
	queryParams := fmt.Sprintf("id=%d", id)
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

func TestGETFamilymemberByID(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzQ5MTg4NjI5fQ.rB9K1sLRhvDiqU_nMo9ynBe5n1y5mxPHYExJuWQpqdA"

	id := 4
	baseurl := "/admin/fmmodbyid?"
	queryParams := fmt.Sprintf("id=%d", id)
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

func TestGETResumeByID(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzQ5MTg4NjI5fQ.rB9K1sLRhvDiqU_nMo9ynBe5n1y5mxPHYExJuWQpqdA"

	id := 1
	baseurl := "/admin/resumebyid?"
	queryParams := fmt.Sprintf("id=%d", id)
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

func TestGETPoexpByID(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOWYwMDFlNDE2NmNmMjZiZmJkZDNiNGY2N2Q5ZWY2MTciLCJwYXNzd29yZCI6ImM4MTU4MDkzMjMyOWIwMWJjMmZjMTMwYjBjNTIxZDk2Iiwicm9sZSI6ImFkbWluIiwiZXhwIjoxNzUwNjY4MDUxfQ.GR4QGGXnTz9H6KzH0pnVicRauwGcuf0huvTbOOvCVsI"

	id := 3
	baseurl := "/admin/poexpbyid?"
	queryParams := fmt.Sprintf("id=%d", id)
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

func TestGetPositionhistoryByPage(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTExMTExMTExIiwicGFzc3dvcmQiOiI5OTYwZGQ4ODRmMDZlODM5NjllNzM2YjNlYWUwYTUwYiIsInJvbGUiOiJjYWRyZSIsImV4cCI6MTc1MDc3MDQzM30.nP-EW90FIw99oarnnEnjL5MktPn9Rf-FKsiTZST4oWw"

	page := 1
	pageSize := 10
	baseurl := "/cadre/getphmodbypage?"
	queryParams := fmt.Sprintf("page=%d&pagesize=%d", page, pageSize)
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

func TestGetAssesementByPage(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTExMTExMTExIiwicGFzc3dvcmQiOiI5OTYwZGQ4ODRmMDZlODM5NjllNzM2YjNlYWUwYTUwYiIsInJvbGUiOiJjYWRyZSIsImV4cCI6MTc1MDc3MDQzM30.nP-EW90FIw99oarnnEnjL5MktPn9Rf-FKsiTZST4oWw"

	page := 1
	pageSize := 10
	baseurl := "/cadre/getasmodbypage?"
	queryParams := fmt.Sprintf("page=%d&pagesize=%d", page, pageSize)
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

func TestGetPoexpbyPosid(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTExMTExMTExIiwicGFzc3dvcmQiOiI5OTYwZGQ4ODRmMDZlODM5NjllNzM2YjNlYWUwYTUwYiIsInJvbGUiOiJjYWRyZSIsImV4cCI6MTc1MDc3MDQzM30.nP-EW90FIw99oarnnEnjL5MktPn9Rf-FKsiTZST4oWw"

	posid := 5
	baseurl := "/cadre/getposexpbyposid?"
	queryParams := fmt.Sprintf("id=%d", posid)
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

func TestGetPositionhistoryByID(t *testing.T) {
	r := router.InitRouter()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTExMTExMTExIiwicGFzc3dvcmQiOiI5OTYwZGQ4ODRmMDZlODM5NjllNzM2YjNlYWUwYTUwYiIsInJvbGUiOiJjYWRyZSIsImV4cCI6MTc1MDc3MDQzM30.nP-EW90FIw99oarnnEnjL5MktPn9Rf-FKsiTZST4oWw"

	id := 5
	baseurl := "/admin/phmodbyid?"
	queryParams := fmt.Sprintf("id=%d", id)
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
