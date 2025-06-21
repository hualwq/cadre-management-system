package User_service

import (
	"cadre-management/models"
	"cadre-management/pkg/utils"
	"errors"
)

type User struct {
	UserID       string `json:"user_id"`
	Password     string `json:"-"`
	Role         string `json:"role"`
	Name         string `json:"name"`
	DepartmentID uint   `json:"department_id"`
}

// LoginResult 登录结果结构体
type LoginResult struct {
	UserID       string `json:"user_id"`
	Role         string `json:"role"`
	Name         string `json:"name"`
	DepartmentID uint   `json:"department_id"`
}

func (s *User) Login(userid, password string) (utils.TokenPair, *LoginResult, error) {
	user, err := models.Authenticate(userid, password)
	if err != nil {
		return utils.TokenPair{}, nil, err
	}

	// 获取用户角色
	role := user.Role

	tokens, err := utils.GenerateTokenPair(user.UserID, role)
	if err != nil {
		return utils.TokenPair{}, nil, err
	}

	// 构建登录结果
	loginResult := &LoginResult{
		UserID:       user.UserID,
		Role:         user.Role,
		Name:         user.Name,
		DepartmentID: *user.DepartmentID,
	}

	return tokens, loginResult, nil
}

func (s *User) RefreshToken(refreshToken string) (utils.TokenPair, error) {
	claims, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return utils.TokenPair{}, errors.New("无效的 RefreshToken")
	}
	return utils.GenerateTokenPair(claims.UserID, claims.Role)
}

func (s *User) RegistUser() error {
	User := map[string]interface{}{
		"id":            s.UserID,
		"password":      s.Password,
		"name":          s.Name,
		"department_id": s.DepartmentID,
	}
	if err := models.RegisterUser(User); err != nil {
		return err
	}
	return nil
}

func ParaseDepartmentName(departmentName string) (uint, error) {
	department, err := models.GetDepartmentByName(departmentName)
	if err != nil {
		return 0, err
	}
	return department.ID, nil
}

// GetUsersWithFilterService 分页和多条件筛选用户
func (s *User) GetUsersWithFilter(page, pageSize int, userID, name, role string, departmentID uint) ([]models.User, int64, error) {
	return models.GetUsersWithFilter(page, pageSize, userID, name, role, departmentID)
}
