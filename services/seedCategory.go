package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SeedCategory(conn *pgxpool.Conn) {
	categories := []struct {
		Name        string
		Description string
	}{
		{"Elektronik", "Produk-produk elektronik seperti handphone, laptop, dll"},
		{"Pakaian", "Berbagai macam pakaian pria dan wanita"},
		{"Makanan", "Makanan ringan dan berat"},
		{"Minuman", "Minuman dingin dan panas"},
	}

	for _, cat := range categories {
		_, err := conn.Exec(context.Background(),
			`INSERT INTO category (name, description) VALUES ($1, $2) 
			 `,
			cat.Name, cat.Description,
		)
		if err != nil {
			log.Printf("Gagal seeding kategori: %v\n", err)
		} else {
			fmt.Printf("Kategori %s berhasil disimpan\n", cat.Name)
		}
	}
}
