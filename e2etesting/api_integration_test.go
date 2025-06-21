package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"cadre-management/router"
)

var (
	testUser     = "cadre1"
	testPassword = "123456"
	server       *httptest.Server
)

func setup() {
	server = httptest.NewServer(router.InitRouter())
}

func teardown() {
	server.Close()
}

func getFreshToken(t *testing.T) string {
	loginBody := map[string]string{"id": testUser, "password": testPassword}
	body, _ := json.Marshal(loginBody)
	resp, err := http.Post(server.URL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login status: %d", resp.StatusCode)
	}

	var response struct {
		Code int                    `json:"code"`
		Msg  string                 `json:"msg"`
		Data map[string]interface{} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("decode response failed: %v", err)
	}

	if response.Code != 200 {
		t.Fatalf("unexpected response code: %d, msg: %s", response.Code, response.Msg)
	}

	data, ok := response.Data["refresh_token"].(string)
	if !ok {
		t.Fatal("refresh_token not found or not a string")
	}

	return data
}

func getSysaminToken(t *testing.T) string {
	loginBody := map[string]string{"id": "school_admin", "password": "admin"}
	body, _ := json.Marshal(loginBody)
	resp, err := http.Post(server.URL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login status: %d", resp.StatusCode)
	}

	var response struct {
		Code int                    `json:"code"`
		Msg  string                 `json:"msg"`
		Data map[string]interface{} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("decode response failed: %v", err)
	}

	if response.Code != 200 {
		t.Fatalf("unexpected response code: %d, msg: %s", response.Code, response.Msg)
	}

	data, ok := response.Data["refresh_token"].(string)
	if !ok {
		t.Fatal("refresh_token not found or not a string")
	}

	return data
}
func TestAPIRegister(t *testing.T) {

	r := router.InitRouter()
	departmentID := 5

	registerBody := map[string]interface{}{
		// "id":       "test_cadre_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"id":            "cadre1",
		"password":      "123456",
		"name":          "Test User1",
		"department_id": departmentID,
	}

	// 3. 序列化请求体
	body, err := json.Marshal(registerBody)
	if err != nil {
		t.Fatalf("Failed to marshal register body: %v", err)
	}

	// 4. 创建测试请求
	req, err := http.NewRequest(
		"POST",
		"/register",
		bytes.NewBuffer(body),
	)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 5. 创建响应记录器
	w := httptest.NewRecorder()

	// 6. 执行请求
	r.ServeHTTP(w, req)

	// 7. 验证响应
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d. Response: %s", w.Code, w.Body.String())
	}

	// 8. 解析响应
	var result struct {
		AccessToken  string `json:"access_token"`
		Message      string `json:"message"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// 9. 验证token
	if result.AccessToken == "" {
		t.Fatal("Access token is empty in response")
	}

	t.Logf("Successfully registered user %s", registerBody["id"])
}

func TestGetuserID(t *testing.T) {
	setup()
	defer teardown()

	t.Run("GetUserID", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/cadre/getuserid", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Errorf("unexpected status: %d", resp.StatusCode)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		// t.Fatalf("GetUserID response: %s", string(body))
		t.Logf("GetUserID response: %s", string(body))
	})
}

func TestPOSTCadre(t *testing.T) {
	setup()
	defer teardown()
	t.Run("AddCadreInfo", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{
		"user_id":                     "cadre",
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
		"administrative_appointment":  "任命为软件研发部经理"
		}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/cadreinfo", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("AddCadreInfo response: %s", string(bodyBytes))
	})
}

func TestAPIPOSTAssessment(t *testing.T) {
	setup()
	defer teardown()

	t.Run("AddAssessment", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre1","department":"信息工程","category":"教师","assess_dept":"教务处","year":2021,"work_summary":"表现良好", "department_id": 7}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/assessment", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("AddAssessment response: %s", string(bodyBytes))
	})
}

func TestAPIPOSTPositionhistory(t *testing.T) {
	setup()
	defer teardown()

	t.Run("AddPositionHistory", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre1","department":"信息工程","category":"教师","office":"教务处","academic_year":"2023-2024","applied_at_year":2023,"applied_at_month":9,"applied_at_day":1, "department_id": 7}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/positionhistory", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("AddPositionHistory response: %s", string(bodyBytes))
	})
}

func TestAPIPOSTYearPositionhistory(t *testing.T) {
	setup()
	defer teardown()

	t.Run("AddYearPosition", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre","year":"2023","department":"信息工程","position":"班主任", "posid": 4}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/yearposition", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("AddYearPosition response: %s", string(bodyBytes))
	})
}

func TestAPIPOSTResume(t *testing.T) {
	setup()
	defer teardown()

	t.Run("AddResume", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre","start_date":"2010.09","end_date":"2014.07","organization":"清华大学","department":"计算机","position":"学生"}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/resume", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("AddResume response: %s", string(bodyBytes))
	})
}

func TestAPIPOSTFamilymember(t *testing.T) {
	setup()
	defer teardown()

	t.Run("AddFamilyMember", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre","relation":"父亲","name":"张父","birth_date":"1960-01-01","political_status":"群众","work_unit":"工厂"}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/familymember", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("AddFamilyMember response: %s", string(bodyBytes))
	})
}

func TestAPIGetPositionhistoryBypage(t *testing.T) {
	setup()
	defer teardown()

	t.Run("GetPositionHistoryModsByPage", func(t *testing.T) {
		token := getFreshToken(t)
		url := fmt.Sprintf("%s/admin/phmodbypage?department_id=%d", server.URL, 7)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetPositionHistoryModsByPage response: %s", string(body))
	})
}

func TestAPIV1Group(t *testing.T) {
	setup()
	defer teardown()

	// GET /getasmodbypage
	t.Run("GetAssessmentByPage", func(t *testing.T) {
		token := getFreshToken(t)
		url := fmt.Sprintf("%s/admin/getasmodbypage?department_id=%d", server.URL, 7)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetAssessmentByPage response: %s", string(body))
	})

	// GET /getposexpbyposid
	t.Run("GetPosExpByPosID", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/cadre/getposexpbyposid", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetPosExpByPosID response: %s", string(body))
	})

	// PUT/DELETE接口和其它GET接口可仿照上面写法继续补充
}

func TestAPIGetDepartments(t *testing.T) {
	setup()
	defer teardown()
	t.Run("GetDepartments", func(t *testing.T) {
		token := getSysaminToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/sysadmin/departments", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetDepartments response: %s", string(body))
	})
}

func TestAPILogin(t *testing.T) {

	r := router.InitRouter()

	registerBody := map[string]string{
		// "id":       "test_cadre_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"id":       "admin",
		"password": "123456",
	}

	// 3. 序列化请求体
	body, err := json.Marshal(registerBody)
	if err != nil {
		t.Fatalf("Failed to marshal register body: %v", err)
	}

	// 4. 创建测试请求
	req, err := http.NewRequest(
		"POST",
		"/login",
		bytes.NewBuffer(body),
	)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 5. 创建响应记录器
	w := httptest.NewRecorder()

	// 6. 执行请求
	r.ServeHTTP(w, req)

	// 7. 验证响应
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d. Response: %s", w.Code, w.Body.String())
	}

	// 8. 解析响应
	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// 9. 验证token
	if result.AccessToken == "" {
		t.Fatal("Access token is empty in response")
	}

	t.Logf("Successfully registered user %s", registerBody["id"])

}

func TestAPIPostCadre(t *testing.T) {
	setup()
	defer teardown()
	t.Run("PostCadre", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre1","name":"张三","gender":"男","birth_date":"1989.1","ethnic_group":"汉族","native_place":"广东省","birth_place":"广东省","political_status":"中共党员","work_start_date":"2013.1","health_status":"良好","professional_title":"工程师","specialty":"软件开发","phone":"13800138000","current_position":"软件研发部经理","awards_and_punishments":"2018年获得公司优秀员工称号","annual_assessment":"近三年考核结果均为优秀","email":"zhangsan@example.com","filled_by":"李四","full_time_education_degree":"本科","full_time_education_school":"XX大学计算机科学与技术专业","on_the_job_education_degree":"硕士","on_the_job_education_school":"YY大学软件工程专业","reporting_unit":"ZZ公司人力资源部","approval_authority":"同意晋升","administrative_appointment":"任命为软件研发部经理"}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/cadreinfo", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("PutCadre response: %s", string(bodyBytes))
	})
}

func TestAPIPutCadre(t *testing.T) {
	setup()
	defer teardown()
	t.Run("PutCadre", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre1","name":"李四","gender":"男","birth_date":"1989.1","ethnic_group":"汉族","native_place":"广东省","birth_place":"广东省","political_status":"中共党员","work_start_date":"2013.1","health_status":"良好","professional_title":"工程师","specialty":"软件开发","phone":"13800138000","current_position":"软件研发部经理","awards_and_punishments":"2018年获得公司优秀员工称号","annual_assessment":"近三年考核结果均为优秀","email":"zhangsan@example.com","filled_by":"李四","full_time_education_degree":"本科","full_time_education_school":"XX大学计算机科学与技术专业","on_the_job_education_degree":"硕士","on_the_job_education_school":"YY大学软件工程专业","reporting_unit":"ZZ公司人力资源部","approval_authority":"同意晋升","administrative_appointment":"任命为软件研发部经理"}`
		req, _ := http.NewRequest("PUT", server.URL+"/cadre/cadreinfo", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("PutCadre response: %s", string(bodyBytes))
	})
}

func TestAPIGetAssessmentByPage(t *testing.T) {
	setup()
	defer teardown()

	url := fmt.Sprintf("%s/admin/assessmentbypage?page=%d&user_id=%s&department_id=%d", server.URL, 1, "cadre1", 5)
	t.Run("GetAssessmentByPage", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetAssessmentByPage response: %s", string(body))
	})
}

func TestAPIGetAssessmentByID(t *testing.T) {
	setup()
	defer teardown()

	url := fmt.Sprintf("%s/admin/assmodbyid?id=%d", server.URL, 1)

	t.Run("GetAssessmentByID", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetAssessmentByID response: %s", string(body))
	})
}

func TestAPIPostPositionhistory(t *testing.T) {
	setup()
	defer teardown()
	t.Run("PostPositionhistory", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre1","department":"信息工程","category":"教师","office":"教务处","academic_year":"2023-2024","applied_at_year":2023,"applied_at_month":9,"applied_at_day":1}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/positionhistory", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("PostPositionhistory response: %s", string(bodyBytes))
	})
}

func TestAPIPutPositionhistory(t *testing.T) {
	setup()
	defer teardown()
	t.Run("PutPositionhistory", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"id":5, "user_id":"cadre1","department":"信息","category":"教师","office":"教务处","academic_year":"2023-2024","applied_at_year":2023,"applied_at_month":9,"applied_at_day":1}`
		req, _ := http.NewRequest("PUT", server.URL+"/cadre/positionhistory", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("PutPositionhistory response: %s", string(bodyBytes))
	})
}

func TestAPIPostPosexp(t *testing.T) {
	setup()
	defer teardown()
	t.Run("PostPosexp", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"user_id":"cadre1","year":"2023","department":"信息工程","position":"班主任", "posid": 6}`
		req, _ := http.NewRequest("POST", server.URL+"/cadre/yearposition", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("PostPosexp response: %s", string(bodyBytes))
	})
}

func TestAPIGetUserByID(t *testing.T) {
	setup()
	defer teardown()
	t.Run("GetUserByID", func(t *testing.T) {
		token := getSysaminToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/sysadmin/user/cadre1", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetUserByID response: %s", string(bodyBytes))
	})
}

func TestAPIGetDepartmentByID(t *testing.T) {
	setup()
	defer teardown()
	t.Run("GetDepartmentByID", func(t *testing.T) {
		token := getSysaminToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/sysadmin/department/1", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetDepartmentByID response: %s", string(bodyBytes))
	})
}

func TestAPIGetUserRoleList(t *testing.T) {
	setup()
	defer teardown()
	t.Run("GetUserRoleList", func(t *testing.T) {
		token := getSysaminToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/sysadmin/user/role?role=department_admin", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetUserRoleList response: %s", string(bodyBytes))
	})
}

func TestAPILoginAdmin(t *testing.T) {
	setup()
	defer teardown()
	t.Run("LoginAdmin", func(t *testing.T) {
		body := `{"id": "admin", "password": "123456"}`
		req, _ := http.NewRequest("POST", server.URL+"/login", bytes.NewBuffer([]byte(body)))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("LoginAdmin response: %s", string(bodyBytes))
	})
}

func TestAPIGetUserDepartment(t *testing.T) {
	setup()
	defer teardown()
	t.Run("GetUserDepartment", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", server.URL+"/sysadmin/user/department", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("GetUserDepartment response: %s", string(bodyBytes))
	})
}

func TestAPIPostAssessment(t *testing.T) {
	setup()
	defer teardown()
	t.Run("PostAssessment", func(t *testing.T) {
		token := getFreshToken(t)
		body := `{"id": 1, "grade": "优秀"}`
		req, _ := http.NewRequest("POST", server.URL+"/admin/assessment", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		t.Logf("PostAssessment response: %s", string(bodyBytes))
	})
}
