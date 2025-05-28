package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID   string `gorm:"primaryKey;type:varchar(50);column:user_id" json:"id"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Role     string `gorm:"type:ENUM('admin','cadre','sysadmin');default:'cadre';not null" json:"role"`
}

func (User) TableName() string {
	return "cadm_users"
}

func ExistUser(userid, password string) (bool, error) {
	var usr User
	err := db.Select("id").Where(User{UserID: userid, Password: password}).First(&usr).Error
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

	// 比较用户输入的密码和数据库中存储的哈希密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
	var count int64
	db.Model(u).
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Joins("JOIN role_permissions ON role_permissions.role_id = roles.id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("users.id = ? AND permissions.resource = ? AND permissions.action = ?",
			u.UserID, resource, action).
		Count(&count)
	return count > 0
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

	// 构建新用户
	newUser := User{
		UserID:   userID,
		Name:     data["name"].(string),
		Password: string(hashedPassword),
		Role:     "cadre", // 默认角色
	}

	return db.Create(&newUser).Error
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
