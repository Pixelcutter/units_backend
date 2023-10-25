package service

import (
	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type ItemService interface {
	SaveItem(model.NewItemRequest) (*model.ItemResponse, error)
	FetchAllItems(int) ([]model.DisplayItem, error)
}

type itemService struct {
	repository repository.UnitsRepository
}

func (service *itemService) SaveItem(itemRequest model.NewItemRequest) (*model.ItemResponse, error) {
	newItem, err := service.repository.SaveItem(itemRequest.Item)
	if err != nil {
		return nil, err
	}

	componentList, err := service.repository.SaveComponents(itemRequest.Components, newItem.ID)

	return &model.ItemResponse{Item: newItem, Components: componentList}, nil
}

func (service *itemService) FetchAllItems(userID int) ([]model.DisplayItem, error) {
	items, err := service.repository.FetchAllItems(userID)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func NewItemService(repository repository.UnitsRepository) ItemService {
	return &itemService{
		repository: repository,
	}
}
