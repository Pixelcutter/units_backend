package service

import "github.com/Pixelcutter/units_backend/cmd/server/model"

type ItemService interface {
	SaveItem(model.Item) model.TestItem
	FetchAllItems() []model.Item
}

type itemService struct {
	items []model.Item
}

func (service *itemService) SaveItem(item model.Item) model.TestItem {
	service.items = append(service.items, item)
	newItem := model.TestItem{
		Item: item,
		R1:   "random string",
		R2:   "another random string",
	}
	return newItem
}

func (service *itemService) FetchAllItems() []model.Item {
	return service.items
}

func NewItemService() ItemService {
	return &itemService{}
}
