package service

import (
	"math/rand"
	"time"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
)

type ItemService interface {
	SaveItem(model.NewItem) model.Item
	FetchAllItems() []model.Item
}

type itemService struct {
	items      []model.Item
	repository repository.UnitsRepository
}

func (service *itemService) SaveItem(item model.NewItem) model.Item {
	newItem := model.Item{
		ID:          rand.Intn(1000000000),
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
