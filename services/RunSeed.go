package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunSeed(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Gagal acquire koneksi:", err)
		return
	}
	defer conn.Release()

	SeedCategory(conn)
	SeedProduct(conn)
	SeedProducts_Category(conn)
}
