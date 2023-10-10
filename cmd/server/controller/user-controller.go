package controller

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FetchAllUsers() []model.User
	SaveUser(ctx *gin.Context) model.User
}

type user_controller struct {
	service service.UserService
}

func (c *user_controller) FetchAllUsers() []model.User {
	return c.service.FetchAllUsers()
}

func (c *user_controller) SaveUser(ctx *gin.Context) model.User {
	var user model.UserDetails
	ctx.BindJSON(&user)
	newUser := c.service.SaveUser(user)
	return newUser
}

func NewUserController(service service.UserService) UserController {
	return &user_controller{
		service: service,
	}
}
