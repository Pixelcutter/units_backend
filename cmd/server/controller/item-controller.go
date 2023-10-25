package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
	"github.com/Pixelcutter/units_backend/cmd/server/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ItemController interface {
	FetchAllItems(ctx *gin.Context)
	SaveItem(ctx *gin.Context)
}

type item_controller struct {
	service service.ItemService
}

func (c *item_controller) FetchAllItems(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid request")
		return
	}

	items, err := c.service.FetchAllItems(userID)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, items)
	case pgx.ErrNoRows:
		ctx.String(http.StatusNotFound, "No items found")
	default:
		ctx.String(http.StatusInternalServerError, "Internal server error")
	}
}

func (c *item_controller) SaveItem(ctx *gin.Context) {
	var itemRequest model.NewItemRequest
	err := ctx.BindJSON(&itemRequest)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusBadRequest, "ERROR: Missing required fields")
		return
	}

	newItem, err := c.service.SaveItem(itemRequest)
	switch err {
	case nil:
		ctx.JSON(http.StatusCreated, newItem)
	case repository.UniqueViolation:
		ctx.String(http.StatusBadRequest, "Item already exists")
	default:
		ctx.String(http.StatusInternalServerError, "Internal server error")
	}
}

func NewItemController(service service.ItemService) ItemController {
	return &item_controller{
		service: service,
	}
}
