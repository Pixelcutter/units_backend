package repository

import "github.com/Pixelcutter/units_backend/cmd/server/model"

type UnitsRepository interface {
	CloseDB()
	SaveUser(model.UserDetails) (model.User, error)
	// FindAllUser() ([]model.User, error)
	FetchUser(int) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(int) error
	SaveItem(model.NewItem) (model.Item, error)
	FetchAllItems(int) ([]model.DisplayItem, error)
	FetchOneItem(int) (model.DisplayItem, error)
	UpdateItem(model.DisplayItem) (model.DisplayItem, error)
	DeleteItem(int) error
	SaveComponents([]model.Component, int) ([]model.Component, error)
	// SaveCategory()
	// FindAllCategories()
	// FindOneCategory()
	// UpdateCategory()
	// DeleteCategory()
}
