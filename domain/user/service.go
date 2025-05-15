package user

import "github.com/mysterybee07/office-project-setup/domain/model"

type UserService struct {
	userRepo *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepo.CreateUser(user)
}
