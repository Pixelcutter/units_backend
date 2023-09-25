package service

import "github.com/Pixelcutter/units_backend/cmd/server/model"

type UserService interface {
	SaveUser(model.User) model.User
	FetchAllUsers() []model.User
}

type userService struct {
	users []model.User
}

func (service *userService) SaveUser(user model.User) model.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FetchAllUsers() []model.User {
	return service.users
}

func NewUserService() UserService {
	return &userService{}
}
