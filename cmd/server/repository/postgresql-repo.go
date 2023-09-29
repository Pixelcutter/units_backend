package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/jackc/pgx/v5"
)

type postgresRepo struct {
	DbPath string
	conn *pgx.Conn
}

func (repository *postgresRepo) SaveUser(user model.User) (model.User, error) {
	return user, nil
}

func (repository *postgresRepo) FindAllUser() ([]model.User, error) {
	return []model.User{}, nil
}

func (repository *postgresRepo) FindOneUser(id int) (model.User, error) {
	return model.User{}, nil
}

func (repository *postgresRepo) UpdateUser(user model.User) (model.User, error) {
	return user, nil
}

func (repository *postgresRepo) DeleteUser(id int) error {
	return nil
}

func (repository *postgresRepo) SaveItem(item model.Item) (model.Item, error) {
	return item, nil
}

func (repository *postgresRepo) FindAllItem() ([]model.Item, error) {
	return []model.Item{}, nil
}

func (repository *postgresRepo) FindOneItem(id int) (model.Item, error) {
	return model.Item{}, nil
}

func (repository *postgresRepo) UpdateItem(item model.Item) (model.Item, error) {
	return item, nil
}

func (repository *postgresRepo) DeleteItem(id int) error { 
	return nil
}

func (repository *postgresRepo) CloseDB() {
	repository.conn.Close(context.Background())
}

func NewPostgresRepo(dbPath string) UnitsRepository {
	conn, err := pgx.Connect(context.Background(), dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &postgresRepo{
		DbPath: dbPath,
		conn: conn,
	}
}