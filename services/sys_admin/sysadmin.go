package sys_admin

import (
	"cadre-management/models"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type ChangeUserRole struct {
	CadreID string `json:"user_id"`
	Role    string `json:"role"`
}

type GetUser struct {
	PageNum  int
	PageSize int
}

func (u *User) GetAllUser() ([]User, error) {
	// 调用model层获取原始数据
	dbUsers, err := models.GetAllUser()
	if err != nil {
		return nil, err
	}

	// 转换为service层的User结构
	var serviceUsers []User
	for _, dbUser := range dbUsers {
		serviceUsers = append(serviceUsers, User{
			Name: dbUser.Name,
			Role: dbUser.Role,
		})
	}

	return serviceUsers, nil
}

func (u *GetUser) GetUserByPage(page, pageSize int) ([]User, error) {
	// 调用model层获取原始数据
	dbUsers, err := models.GetUserByPage(page, pageSize)
	if err != nil {
		return nil, err
	}

	// 转换为service层的User结构
	var serviceUsers []User
	for _, dbUser := range dbUsers {
		serviceUsers = append(serviceUsers, User{
			Name: dbUser.Name,
			Role: dbUser.Role,
		})
	}

	return serviceUsers, nil
}

func (s *ChangeUserRole) ChangeUserRole(userID, newRole string) error {
	return models.ChangeUserRole(userID, newRole)
}
