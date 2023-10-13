package service

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type UserService interface {
	SaveUser(model.UserDetails) (model.User, error)
	FetchUser(id int) (model.User, error)
}

type userService struct {
	users      []model.User
	repository repository.UnitsRepository
}

func (service *userService) SaveUser(user model.UserDetails) (model.User, error) {
	newUser, err := service.repository.SaveUser(user)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func (service *userService) FetchUser(id int) (model.User, error) {
	user, err := service.repository.FetchUser(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserService(repository repository.UnitsRepository) UserService {
	return &userService{
		repository: repository,
	}
}
