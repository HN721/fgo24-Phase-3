package models

import (
	"context"
	"errors"
	"nastha-test/utils"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func SaveUser(users Users) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err = conn.Exec(context.Background(), query, users.Name, users.Email, users.Password)
	if err != nil {
		return err
	}
	return nil
}

func Login(users Users) (*Users, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user Users

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, email, password, role 
		 FROM users 
		 WHERE email = $1 AND password = $2`, users.Email, users.Password,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("email atau password salah")
		}
		return nil, err
	}

	return &user, nil
}
