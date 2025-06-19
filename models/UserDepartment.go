package models

type UserDepartment struct {
	UserID       string `gorm:"primaryKey;size:50" json:"user_id"`
	DepartmentID uint   `gorm:"primaryKey" json:"department_id"`

	User       User       `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Department Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
