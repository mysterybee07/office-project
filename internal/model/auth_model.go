package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UserId uint   `gorm:"not null"`
	Token  string `gorm:"not null;type:text"`
	User   User   `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Auth) TableName() string {
	return "auth.auth"
}
