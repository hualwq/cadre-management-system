package e2e

import (
	"bytes"
	"cadre-management/models"
	"cadre-management/pkg/logging"
	"cadre-management/pkg/setting"
	"cadre-management/pkg/utils"
	"cadre-management/router"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	baseURL      = "http://localhost:8088"
	testUser     = "cadre"
	testPassword = "123456"
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

// 每次调用前都刷新token
func getFreshToken(t *testing.T) string {
	loginBody := map[string]string{"id": "", "password": testPassword}
	body, _ := json.Marshal(loginBody)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("login status: %d", resp.StatusCode)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["access_token"].(string)
}

func TestAPIRegister(t *testing.T) string {

	r := router.InitRouter()

	registerBody := map[string]string{
		// "id":       "test_cadre_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		"id":       testUser,
		"password": "123",
		"name":     "Test User",
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
		AccessToken string `json:"access_token"`
		Message     string `json:"message"`
		Success     bool   `json:"success"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// 9. 验证token
	if result.AccessToken == "" {
		t.Fatal("Access token is empty in response")
	}

	t.Logf("Successfully registered user %s", registerBody["id"])
	return result.AccessToken
}

func TestAPIIntegration(t *testing.T) {
	t.Run("GetUserID", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", baseURL+"/cadre/getuserid", nil)
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
		t.Logf("GetUserID response: %s", string(body))
	})

	t.Run("GetUserRole", func(t *testing.T) {
		token := getFreshToken(t)
		req, _ := http.NewRequest("GET", baseURL+"/getuserrole", nil)
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
		t.Logf("GetUserRole response: %s", string(body))
	})

	// 可继续添加更多接口测试，每次都用getFreshToken(t)获取新token
}
