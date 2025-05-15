package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Role     string `json:"role" gorm:"default:user"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
}

func (User) TableName() string {
	return "public.users"
}
