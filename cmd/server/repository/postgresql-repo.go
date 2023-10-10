package repository

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepo struct {
	DbPath string
	db     *pgxpool.Pool
}

func (repository *postgresRepo) CloseDB() {
	repository.db.Close()
}

func NewPostgresRepo(dbPath string) UnitsRepository {
	var (
		once       sync.Once
		pgInstance *postgresRepo
	)

	once.Do(func() {
		db, err := pgxpool.New(context.Background(), dbPath)
		if err != nil {
			log.Fatal(err)
		}

		pgInstance = &postgresRepo{
			DbPath: dbPath,
			db:     db,
		}
	})

	return pgInstance
}

func (repository *postgresRepo) SaveUser(user model.UserDetails) (model.User, error) {
	newUser := model.User{
		ID:        rand.Intn(1000000000),
		Signup:    time.Now(),
		LastLogin: time.Now(),
		Details:   user,
	}

	_, err := repository.db.Exec(context.Background(), `
		INSERT INTO "user" (id, email, pass_hash, username, signup, last_login)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, newUser.ID, user.Email, user.PassHash, user.Username, newUser.Signup, newUser.LastLogin)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
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

func (repository *postgresRepo) SaveItem(item model.NewItem) (model.Item, error) {
	return model.Item{}, nil
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
