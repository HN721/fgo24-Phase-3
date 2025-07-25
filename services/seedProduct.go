package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedProduct(conn *pgxpool.Conn) {
	products := []struct {
		Kode          string
		Name          string
		ImageURL      string
		PurchasePrice float64
		SellingPrice  float64
		Stock         int
	}{
		{"P001", "Laptop Asus ROG", "https://example.com/rog.jpg", 15000000, 17000000, 10},
		{"P002", "Kaos Polos", "https://example.com/kaos.jpg", 30000, 50000, 100},
		{"P003", "Keripik Kentang", "https://example.com/snack.jpg", 5000, 8000, 200},
		{"P004", "Air Mineral", "https://example.com/air.jpg", 2000, 3500, 300},
	}

	for _, p := range products {
		_, err := conn.Exec(context.Background(),
			`INSERT INTO products (kode, name, image_url, purchase_price, selling_price, stock)
			 VALUES ($1, $2, $3, $4, $5, $6)
			`,
			p.Kode, p.Name, p.ImageURL, p.PurchasePrice, p.SellingPrice, p.Stock,
		)
		if err != nil {
			log.Printf("Gagal insert produk %s: %v\n", p.Name, err)
		} else {
			fmt.Printf("Produk %s berhasil ditambahkan\n", p.Name)
		}
	}
}
