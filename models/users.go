package models

import (
	"context"
	"nastha-test/utils"
)

type Users struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SaveUser(users Users) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `INSERT INTO users (name,email,password)VALUES($1,$2,$3)`
	_, err = conn.Query(context.Background(), query, users.Name, users.Email, users.Password)
	if err != nil {
		return err
	}
	return nil
}
