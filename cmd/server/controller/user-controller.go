package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FetchUser(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
}

type user_controller struct {
	service service.UserService
}

func (c *user_controller) FetchUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid user id")
	}

	user, err := c.service.FetchUser(userId)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *user_controller) SaveUser(ctx *gin.Context) {
	var user model.UserDetails
	ctx.BindJSON(&user)
	newUser, err := c.service.SaveUser(user)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Bad Request", "message": err.Error()},
		)
	} else {
		ctx.JSON(
			http.StatusCreated,
			newUser,
		)
	}
}

func NewUserController(service service.UserService) UserController {
	return &user_controller{
		service: service,
	}
}
