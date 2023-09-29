package service

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type ItemService interface {
	SaveItem(model.Item) model.TestItem
	FetchAllItems() []model.Item
}

type itemService struct {
	items []model.Item
	repository repository.UnitsRepository
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

func NewItemService(repository repository.UnitsRepository) ItemService {
	return &itemService{
		repository: repository,
	}
}
