package service

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type UserService interface {
	SaveUser(model.User) model.User
	FetchAllUsers() []model.User
}

type userService struct {
	users []model.User
	repository repository.UnitsRepository
}

func (service *userService) SaveUser(user model.User) model.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FetchAllUsers() []model.User {
	return service.users
}

func NewUserService(repository repository.UnitsRepository) UserService {
	return &userService{
		repository: repository,
	}
}
