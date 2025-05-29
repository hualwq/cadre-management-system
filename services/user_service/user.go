package user_service

import (
	"cadre-management/models"
	"cadre-management/pkg/utils"
	"errors"
)

type User struct {
	UserID   string `json:"user_id"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

func (s *User) Login(userid, password string) (string, error) {
	// 1. 认证用户
	user, err := models.Authenticate(userid, password)
	if err != nil {
		return "", err // 认证失败
	}

	// 2. 生成 JWT
	token, err := utils.GenerateToken(user.UserID, user.Password, user.Role)
	if err != nil {
		return "", errors.New("生成 token 失败")
	}

	return token, nil
}

// RefreshToken 刷新 JWT（可选功能，根据需求实现）
func (s *User) RefreshToken(oldToken string) (string, error) {
	claims, err := utils.ParseToken(oldToken)
	if err != nil {
		return "", errors.New("无效的 Token")
	}

	// 检查用户是否存在（可选）
	user, err := models.GetUserByID(claims.UserID)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 生成新 Token
	newToken, err := utils.GenerateToken(user.UserID, user.Password, user.Role)
	if err != nil {
		return "", errors.New("刷新 Token 失败")
	}

	return newToken, nil
}

func (s *User) RegistUser() error {
	User := map[string]interface{}{
		"id":       s.UserID,
		"password": s.Password,
		"name":     s.Name,
	}
	if err := models.RegisterUser(User); err != nil {
		return err
	}
	return nil
}

func (s *User) GetRole(user_id string) (string, error) {
	user, err := models.GetUserByID(user_id)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}
