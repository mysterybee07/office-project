package user

import (
	"github.com/mysterybee07/office-project-setup/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
