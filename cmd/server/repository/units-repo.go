package repository

import "github.com/Pixelcutter/units_backend/cmd/server/model"

type UnitsRepository interface {
	CloseDB()
	SaveUser(model.User) (model.User, error)
	FindAllUser() ([]model.User, error)
	FindOneUser(int) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(int) error
	SaveItem(model.Item) (model.Item, error)
	FindAllItem() ([]model.Item, error)
	FindOneItem(int) (model.Item, error)
	UpdateItem(model.Item) (model.Item, error)
	DeleteItem(int) error
	// SaveCategory()
	// FindAllCategories()
	// FindOneCategory()
	// UpdateCategory()
	// DeleteCategory()
}