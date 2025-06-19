package User_service

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

func (s *User) Login(userid, password string) (utils.TokenPair, error) {
	user, err := models.Authenticate(userid, password)
	if err != nil {
		return utils.TokenPair{}, err
	}
	role := "cadre" // 可根据业务查role
	return utils.GenerateTokenPair(user.UserID, role)
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
		"id":       s.UserID,
		"password": s.Password,
		"name":     s.Name,
	}
	if err := models.RegisterUser(User); err != nil {
		return err
	}
	return nil
}
