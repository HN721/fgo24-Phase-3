package utils

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnect() (*pgxpool.Pool, error) {
	connection := "postgres://postgres:123@127.0.0.1:5434/inventory"

	pool, err := pgxpool.New(context.Background(), connection)
	if err != nil {
		return nil, err
	}
	return pool, nil

}
