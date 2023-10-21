package service

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type ItemService interface {
	SaveItem(model.NewItemRequest) (model.Item, error)
	FetchAllItems() []model.Item
}

type itemService struct {
	items      []model.Item
	repository repository.UnitsRepository
}

func (service *itemService) SaveItem(itemRequest model.NewItemRequest) (model.Item, error) {
	newItem, err := service.repository.SaveItem(itemRequest.Item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (service *itemService) FetchAllItems() []model.Item {
	return service.items
}

func NewItemService(repository repository.UnitsRepository) ItemService {
	return &itemService{
		repository: repository,
	}
}
