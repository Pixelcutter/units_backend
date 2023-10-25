package controller

import (
	"net/http"
	"strconv"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
	"github.com/Pixelcutter/units_backend/cmd/server/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
		ctx.String(http.StatusBadRequest, "invalid request")
		return
	}

	// TODO: fix error handling
	user, err := c.service.FetchUser(userId)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, user)
	case pgx.ErrNoRows:
		ctx.String(http.StatusNotFound, "A user with that id does not exist")
	default:
		ctx.String(http.StatusInternalServerError, "Internal server error")
	}
}

func (c *user_controller) SaveUser(ctx *gin.Context) {
	var user model.UserDetails
	if err := ctx.BindJSON(&user); err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	newUser, err := c.service.SaveUser(user)

	switch err {
	case nil:
		ctx.JSON(http.StatusCreated, newUser)
	case repository.UniqueViolation:
		ctx.String(http.StatusBadRequest, "Username or email already exists")
	default:
		ctx.String(http.StatusInternalServerError, "Internal server error")
	}
}

func NewUserController(service service.UserService) UserController {
	return &user_controller{
		service: service,
	}
}
