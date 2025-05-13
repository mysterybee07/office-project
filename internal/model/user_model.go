package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Role     string `json:"role"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (User) TableName() string {
	return "public.users"
}
