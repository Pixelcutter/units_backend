package repository

import (
	"context"
	"log"
	"sync"

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

// func CheckForUser(db *pgxpool.Conn, email string, username string) (bool, error) {
// 	var id int
// 	query := `SELECT id from units_user WHERE email = $1 OR username = $2`
// 	err := db.QueryRow(context.Background(), query, email, username).Scan(&id)
// 	if err == nil {
// 		return true, nil
// 	}

// 	return false, err
// }

func (repository *postgresRepo) SaveUser(user model.UserDetails) (model.User, error) {
	// Insert into db and return new user
	query := `
			 INSERT INTO units_user (email, pass_hash, username)
			 VALUES ($1, $2, $3)
			 RETURNING id, email, username, signup, last_login
			 `
	newUser := model.User{}
	err := repository.db.QueryRow(
		context.Background(), query, user.Email, user.PassHash, user.Username).
		Scan(
			&newUser.ID,
			&newUser.Email,
			&newUser.Username,
			&newUser.Signup,
			&newUser.LastLogin,
		)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func (repository *postgresRepo) FindAllUser() ([]model.User, error) {
	return []model.User{}, nil
}

func (repository *postgresRepo) FetchUser(id int) (model.User, error) {
	query := `SELECT id, signup, last_login, email, username FROM units_user WHERE id = $1`
	var user model.User
	err := repository.db.QueryRow(
		context.Background(), query, id).
		Scan(
			&user.ID,
			&user.Signup,
			&user.LastLogin,
			&user.Email,
			&user.Username,
		)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
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
