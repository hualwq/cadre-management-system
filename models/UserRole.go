package models

type UserRole struct {
	UserID string `gorm:"primaryKey;size:50" json:"user_id"`
	Role   string `gorm:"primaryKey;size:50" json:"role"` // e.g., system_admin, school_admin, department_admin, cadre

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
