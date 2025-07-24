package models

import (
	"context"
	"nastha-test/utils"

	"github.com/jackc/pgx/v5"
)

type Products struct {
	Products  string `json:"products"`
	Kategori  string `json:"category"`
	Deskripsi string `json:"description"`
}

func GetAllProductCategory() ([]Products, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	query := `
	SELECT 
		p.name AS products, 
		c.name AS category, 
		c.description
	FROM products p
	JOIN products_category pc ON p.id = pc.product_id
	JOIN categories c ON c.id = pc.category_id
`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	data, err := pgx.CollectRows[Products](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
