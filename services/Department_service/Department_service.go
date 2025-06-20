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
	// 1. 更新用户的院系ID
	err := s.DB.Model(&models.User{}).Where("user_id = ?", userID).Update("department_id", departmentID).Error
	if err != nil {
		return err
	}
	// 2. 赋予用户department_admin角色
	return s.DB.Model(&models.User{}).Where("user_id = ?", userID).Update("role", "department_admin").Error
}

// 取消用户的院系管理员身份
func (s *DepartmentService) UnsetDepartmentAdmin(userID string, departmentID uint) error {
	// 将用户角色降级为cadre
	return s.DB.Model(&models.User{}).Where("user_id = ?", userID).Update("role", "cadre").Error
}

// 查询某院系的管理员用户
func (s *DepartmentService) GetDepartmentAdmins(departmentID uint) ([]models.User, error) {
	var users []models.User
	err := s.DB.Where("department_id = ? AND role = ?", departmentID, "department_admin").Find(&users).Error
	return users, err
}
