package service

import (
	"log"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type UserService interface {
	SaveUser(model.UserDetails) model.User
	FetchAllUsers() []model.User
}

type userService struct {
	users      []model.User
	repository repository.UnitsRepository
}

func (service *userService) SaveUser(user model.UserDetails) model.User {
	newUser, err := service.repository.SaveUser(user)
	if err != nil {
		log.Fatal(err)
	}

	return newUser
}

func (service *userService) FetchAllUsers() []model.User {
	return service.users
}

func NewUserService(repository repository.UnitsRepository) UserService {
	return &userService{
		repository: repository,
	}
}
