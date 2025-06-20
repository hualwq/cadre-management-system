package models

import (
	"errors"
)

type Department struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Users []User `gorm:"many2many:user_departments;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users"`
}

func (Department) TableName() string {
	return "cadm_departments"
}

// GetAllDepartments 获取所有院系
func GetAllDepartments() ([]Department, error) {
	var departments []Department
	err := db.Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// CreateDepartment 创建新院系
func CreateDepartment(department *Department) error {
	return db.Create(department).Error
}

// UpdateDepartment 更新院系信息
func UpdateDepartment(id uint, department *Department) error {
	result := db.Model(&Department{}).Where("id = ?", id).Updates(department)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("院系不存在")
	}
	return nil
}

// DeleteDepartment 删除院系
func DeleteDepartment(id uint) error {
	result := db.Delete(&Department{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("院系不存在")
	}
	return nil
}

// GetDepartmentByID 根据ID获取院系
func GetDepartmentByID(id uint) (*Department, error) {
	var department Department
	err := db.First(&department, id).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

// GetDepartmentByName 根据名称获取院系
func GetDepartmentByName(name string) (*Department, error) {
	var department Department
	err := db.Where("name = ?", name).First(&department).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}
