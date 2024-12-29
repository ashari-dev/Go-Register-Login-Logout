package repository

import (
	"context"
	"trial/lib"
	"trial/models"

	"github.com/jackc/pgx/v5"
)

func CreateUser(data models.Users) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	sql := `INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.FullName, data.Email, data.Password)
	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Users])

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users WHERE email = $1`

	rows, err := db.Query(context.Background(), sql, email)
	if err != nil {
		return models.Users{}, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func GetProfileByUserId(id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users WHERE id = $1`

	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Users{}, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}
