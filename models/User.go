package models

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID       string `gorm:"primaryKey;size:50" json:"user_id"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	Name         string `gorm:"type:varchar(50);not null" json:"name"`
	Role         string `gorm:"type:varchar(50);not null;default:'cadre'" json:"role"`
	DepartmentID *uint  `gorm:"column:department_id" json:"department_id"`

	Departments []Department `gorm:"many2many:user_departments;" json:"departments"`
}

func (User) TableName() string {
	return "cadm_users"
}

func ExistUser(userid, password string) (bool, error) {
	var usr User
	err := db.Where("user_id = ?", userid).First(&usr).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if usr.UserID != "" {
		return true, nil
	}

	return false, nil
}

func Authenticate(userid, password string) (*User, error) {
	var user User
	if err := db.Where("user_id = ?", userid).First(&user).Error; err != nil {
		return nil, err
	}

	fmt.Println(user.Password)

	// 比较用户输入的密码和数据库中存储的哈希密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	return &user, nil
}

func GetAllUser() ([]User, error) {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有找到记录，返回空切片和nil错误
			return []User{}, nil
		}
	}

	// 出于安全考虑，确保密码字段不被返回
	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func GetUserByPage(page, pageSize int) ([]User, error) {
	var users []User

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 执行分页查询
	if err := db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有找到记录，返回空切片和nil错误
			return []User{}, nil
		}
		// 其他错误直接返回
		return nil, err
	}

	// 出于安全考虑，确保密码字段不被返回
	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func (u *User) HasPermission(resource, action string) bool {
	// 简化权限检查，基于用户角色
	// 这里可以根据需要实现更复杂的权限逻辑
	switch u.Role {
	case "sysadmin":
		return true // 系统管理员拥有所有权限
	case "admin":
		return resource != "system" // 管理员不能访问系统级资源
	case "cadre":
		return resource == "cadre" && (action == "read" || action == "write")
	default:
		return false
	}
}

func RegisterUser(data map[string]interface{}) error {
	userID, ok := data["id"].(string)
	if !ok || userID == "" {
		return errors.New("无效的用户ID")
	}

	// 检查用户是否已存在
	var existing User
	if err := db.First(&existing, "user_id = ?", userID).Error; err == nil {
		return errors.New("用户已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 加密密码
	rawPassword, ok := data["password"].(string)
	if !ok || rawPassword == "" {
		return errors.New("无效的密码")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 处理院系ID
	departmentID, ok := data["department_id"].(uint)
	if !ok || departmentID == 0 {
		return errors.New("无效的院系ID")
	}

	// 验证院系是否存在
	var department Department
	if err := db.First(&department, departmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("院系不存在")
		}
		return err
	}

	// 构建新用户
	newUser := User{
		UserID:       userID,
		Name:         data["name"].(string),
		Password:     string(hashedPassword),
		Role:         "cadre", // 默认角色
		DepartmentID: &departmentID,
	}

	// 创建用户
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByID(userID string) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err // 其他数据库错误
	}
	return &user, nil
}

func ChangeUserRole(userID, newRole string) error {
	if userID == "" || newRole == "" {
		return errors.New("用户 ID 和新角色不能为空")
	}

	// 检查新角色是否合法
	validRoles := []string{"admin", "cadre", "sysadmin"}
	valid := false
	for _, role := range validRoles {
		if role == newRole {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("无效的角色")
	}

	// 更新用户角色
	var user User
	result := db.Model(&user).Where("user_id = ?", userID).Update("role", newRole)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetUsersWithFilter 支持分页和多条件筛选
func GetUsersWithFilter(page, pageSize int, userID, name, role string, departmentID uint) ([]User, int64, error) {
	var users []User
	var total int64
	dbQuery := db.Model(&User{})

	if userID != "" {
		dbQuery = dbQuery.Where("user_id LIKE ?", "%"+userID+"%")
	}
	if name != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+name+"%")
	}
	if role != "" {
		dbQuery = dbQuery.Where("role = ?", role)
	}
	if departmentID != 0 {
		dbQuery = dbQuery.Where("department_id = ?", departmentID)
	}

	dbQuery.Count(&total)

	offset := (page - 1) * pageSize
	if err := dbQuery.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	for i := range users {
		users[i].Password = ""
	}
	return users, total, nil
}

// ExistCadreInfo 查询干部信息是否存在
func ExistCadreInfo(userID string) (bool, error) {
	var count int64
	err := db.Table("cadm_cadreinfo").Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistResume 查询简历是否存在
func ExistResume(id uint) (bool, error) {
	var count int64
	err := db.Table("cadm_resume_entries").Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistFamilyMember 查询家庭成员是否存在
func ExistFamilyMember(id uint) (bool, error) {
	var count int64
	err := db.Table("cadm_family_members").Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistPositionHistory 查询岗位历史是否存在
func ExistPositionHistory(id uint) (bool, error) {
	var count int64
	err := db.Table("cadm_position_histories").Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistAssessment 查询考核是否存在
func ExistAssessment(id uint) (bool, error) {
	var count int64
	err := db.Table("cadm_assessments").Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func ChangeUserRoleByID(userID, newRole string) error {
	if userID == "" || newRole == "" {
		return errors.New("用户 ID 和新角色不能为空")
	}

	// 检查新角色是否合法
	validRoles := []string{"department_admin", "cadre", "school_admin"}
	valid := false
	for _, role := range validRoles {
		if role == newRole {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("无效的角色")
	}

	// 更新用户角色
	var user User
	result := db.Model(&user).Where("user_id = ?", userID).Update("role", newRole)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetUsersByRole 返回所有指定角色的用户
func GetUsersByRole(role string) ([]User, error) {
	var users []User
	err := db.Table("cadm_users").Where("role = ?", role).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserDepartment(userID string) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
