package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedUsers(conn *pgxpool.Conn) {
	users := []struct {
		Name     string
		Email    string
		Password string
		Role     string
	}{
		{"Admin", "admin123@gmail.com", "admin123", "admin"},
		{"bayu", "bayu@gmail.com", "bayu123", "user"},
		{"hosea", "hosea@gmail.com", "hosea123", "user"},
	}

	for _, u := range users {
		_, err := conn.Exec(context.Background(),
			`INSERT INTO users (name, email, password, role)
			 VALUES ($1, $2, $3, $4)`,
			u.Name, u.Email, u.Password, u.Role,
		)
		if err != nil {
			log.Printf("Gagal insert user %s: %v\n", u.Name, err)
		} else {
			fmt.Printf("User %s berhasil ditambahkan\n", u.Name)
		}
	}
}
