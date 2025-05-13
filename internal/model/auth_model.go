package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UserId uint   `gorm:"not null"`
	Token  string `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Auth) TableName() string {
	return "auth.auth"
}
