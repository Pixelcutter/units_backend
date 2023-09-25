package controller

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/service"
	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FetchAllItems() []model.Item
	SaveItem(ctx *gin.Context) model.TestItem
}

type item_controller struct {
	service service.ItemService
}

func (c *item_controller) FetchAllItems() []model.Item {
	return c.service.FetchAllItems()
}

func (c *item_controller) SaveItem(ctx *gin.Context) model.TestItem {
	var item model.Item
	ctx.BindJSON(&item)
	newItem := c.service.SaveItem(item)
	return newItem
}

func NewItemController(service service.ItemService) ItemController {
	return &item_controller{
		service: service,
	}
}
