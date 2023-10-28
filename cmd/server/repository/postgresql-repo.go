package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/Pixelcutter/units_backend/cmd/server/model"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBError struct {
	Code string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("Database Error Code: %v", e.Code)
}

var (
	UniqueViolation = &DBError{"23505"}
	InternalError   = &DBError{"500"}
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
	if err := dbError(err); err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

// func (repository *postgresRepo) FindAllUser() ([]model.User, error) {
// 	return []model.User{}, nil
// }

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
	query := `
			 INSERT INTO item
			 (sku, item_name, description, price, category_id, img_path, cost, for_sale, quantity, unit, created_by, updated_by)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			 RETURNING *;
			 `
	newItem := model.Item{}
	err := repository.db.QueryRow(
		context.Background(),
		query,
		item.SKU,
		item.Name,
		item.Description,
		item.Price,
		item.CategoryID,
		item.ImgPath,
		item.Cost,
		item.ForSale,
		item.Quantity,
		item.Unit,
		item.CreatedBy,
		item.UpdatedBy,
	).
		Scan(
			&newItem.ID,
			&newItem.SKU,
			&newItem.CategoryID,
			&newItem.Name,
			&newItem.Description,
			&newItem.ImgPath,
			&newItem.Quantity,
			&newItem.Unit,
			&newItem.Price,
			&newItem.Cost,
			&newItem.ForSale,
			&newItem.CreatedAt,
			&newItem.UpdatedAt,
			&newItem.CreatedBy,
			&newItem.UpdatedBy,
		)
	if err := dbError(err); err != nil {
		return model.Item{}, err
	}

	return newItem, nil
}

func (repository *postgresRepo) FetchAllItems(userID int) ([]model.DisplayItem, error) {
	query := `
	SELECT item.id, category_name, item_name, item.description, img_path, quantity, unit, price
	FROM item
	INNER JOIN category
	ON item.category_id = category.id
	WHERE item.created_by = $1 AND category.created_by = $1 AND for_sale = true
	LIMIT 10;
	`

	rows, err := repository.db.Query(context.Background(), query, userID)
	if err != nil || rows.Err() != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	items := []model.DisplayItem{}
	for rows.Next() {
		item := model.DisplayItem{}
		err := rows.Scan(
			&item.ID,
			&item.CategoryName,
			&item.Name,
			&item.Description,
			&item.ImgPath,
			&item.Quantity,
			&item.Unit,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repository *postgresRepo) FetchOneItem(id int) (model.DisplayItem, error) {
	return model.DisplayItem{}, nil
}

func (repository *postgresRepo) UpdateItem(item model.DisplayItem) (model.DisplayItem, error) {
	return item, nil
}

func (repository *postgresRepo) DeleteItem(id int) error {
	return nil
}

func (repository *postgresRepo) SaveComponents(components []model.Component, parentID int) ([]model.Component, error) {
	query := `
	INSERT INTO component (parent_id, child_id, quantity)
	VALUES ($1, $2, $3)
	RETURNING *;
	`

	for _, component := range components {
		_, err := repository.db.Exec(
			context.Background(),
			query,
			parentID,
			component.ChildId,
			component.Quantity,
		)
		if err != nil {
			return nil, err
		}
	}

	return components, nil
}

func dbError(err error) error {
	var pgerr *pgconn.PgError
	if !errors.As(err, &pgerr) {
		return nil
	}

	switch pgerr.Code {
	case UniqueViolation.Code:
		return UniqueViolation
	default:
		return InternalError
	}
}
