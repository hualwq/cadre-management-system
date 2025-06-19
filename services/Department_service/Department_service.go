package Department_service

import (
	"cadre-management/models"

	"gorm.io/gorm"
)

type DepartmentService struct {
	DB *gorm.DB
}

// 新增院系
func (s *DepartmentService) CreateDepartment(name, description string) error {
	dept := models.Department{Name: name, Description: description}
	return s.DB.Create(&dept).Error
}

// 删除院系
func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.DB.Delete(&models.Department{}, id).Error
}

// 修改院系
func (s *DepartmentService) UpdateDepartment(id uint, name, description string) error {
	return s.DB.Model(&models.Department{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        name,
		"description": description,
	}).Error
}

// 查询所有院系
func (s *DepartmentService) ListDepartments() ([]models.Department, error) {
	var depts []models.Department
	err := s.DB.Find(&depts).Error
	return depts, err
}

// 设置用户为院系管理员
func (s *DepartmentService) SetDepartmentAdmin(userID string, departmentID uint) error {
	// 1. 绑定用户和院系
	ud := models.UserDepartment{UserID: userID, DepartmentID: departmentID}
	err := s.DB.FirstOrCreate(&ud, ud).Error
	if err != nil {
		return err
	}
	// 2. 赋予用户department_admin角色
	ur := models.UserRole{UserID: userID, Role: "department_admin"}
	return s.DB.FirstOrCreate(&ur, ur).Error
}

// 取消用户的院系管理员身份
func (s *DepartmentService) UnsetDepartmentAdmin(userID string, departmentID uint) error {
	// 1. 解绑用户和院系（可选，通常不解绑，只降级角色）
	// 2. 移除department_admin角色
	return s.DB.Delete(&models.UserRole{}, "user_id = ? AND role = ?", userID, "department_admin").Error
}

// 查询某院系的管理员用户
func (s *DepartmentService) GetDepartmentAdmins(departmentID uint) ([]models.User, error) {
	var users []models.User
	err := s.DB.Joins("JOIN user_departments ON user_departments.user_id = users.user_id").
		Joins("JOIN user_roles ON user_roles.user_id = users.user_id").
		Where("user_departments.department_id = ? AND user_roles.role = ?", departmentID, "department_admin").
		Find(&users).Error
	return users, err
}
