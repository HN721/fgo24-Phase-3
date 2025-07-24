package models

import (
	"context"
	"nastha-test/utils"

	"github.com/jackc/pgx/v5"
)

type products_category struct {
	Products    string `json:"products"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type Products struct {
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Image          string  `json:"image"`
	Purchase_price float64 `json:"purchase_price"`
	Selling_price  float64 `json:"selling_price"`
	Stock          int     `json:"stock"`
}

func GetAllProductCategory() ([]products_category, error) {
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
	JOIN category c ON c.id = pc.category_id
`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	data, err := pgx.CollectRows[products_category](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetAllProducts() ([]Products, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	SELECT 
		p.name,
		c.name AS category,
		p.image_url AS image,
		p.purchase_price,
		p.selling_price,
		p.stock
	FROM products p
	JOIN products_category pc ON p.id = pc.product_id
	JOIN category c ON c.id = pc.category_id
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
func GetProductByID(id int) (Products, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Products{}, err
	}
	defer conn.Close()

	query := `
	SELECT 
		p.name,
		c.name AS category,
		p.image_url AS image,
		p.purchase_price,
		p.selling_price,
		p.stock
	FROM products p
	JOIN products_category pc ON p.id = pc.product_id
	JOIN category c ON c.id = pc.category_id
	WHERE p.id = $1
	LIMIT 1
	`

	row, err := conn.Query(context.Background(), query, id)

	data, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Products])
	if err != nil {
		return Products{}, err
	}

	return data, nil
}
