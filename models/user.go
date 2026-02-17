package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=3,max=50"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(255);not null" json:"password" validate:"required,min=6"`
	Role     string `gorm:"type:varchar(50);default:'user'" json:"role" validate:"oneof=admin user"`
}
