package models

type Department struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Users  []User  `gorm:"many2many:user_departments;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users"`
	Cadres []Cadre `gorm:"foreignKey:DepartmentID" json:"cadres"`
}
