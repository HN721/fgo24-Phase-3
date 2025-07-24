package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedProducts_Category(conn *pgxpool.Conn) {

	mappings := []struct {
		ProductID  int
		CategoryID int
	}{
		{5, 1},
		{6, 2},
		{7, 3},
		{8, 4},
	}

	for _, m := range mappings {
		_, err := conn.Exec(context.Background(),
			`INSERT INTO products_category (product_id, category_id)
			 VALUES ($1, $2)
			`,
			m.ProductID, m.CategoryID,
		)
		if err != nil {
			log.Printf("Gagal insert relasi produk %d dan kategori %d: %v\n", m.ProductID, m.CategoryID, err)
		} else {
			fmt.Printf("Relasi produk %d dengan kategori %d berhasil ditambahkan\n", m.ProductID, m.CategoryID)
		}
	}
}
