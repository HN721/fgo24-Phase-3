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
func CreateProduct(p Products, categoryID int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var productID int
	insertProduct := `
		INSERT INTO products (name, image_url, purchase_price, selling_price, stock)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err = tx.QueryRow(context.Background(), insertProduct,
		p.Name, p.Image, p.Purchase_price, p.Selling_price, p.Stock,
	).Scan(&productID)
	if err != nil {
		return err
	}

	insertProductCategory := `
		INSERT INTO products_category (product_id, category_id)
		VALUES ($1, $2)
	`
	_, err = tx.Exec(context.Background(), insertProductCategory, productID, categoryID)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
func UpdateProduct(id int, p Products, categoryID int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	updateProduct := `
		UPDATE products
		SET name = $1, image_url = $2, purchase_price = $3, selling_price = $4, stock = $5
		WHERE id = $6
	`
	_, err = tx.Exec(context.Background(), updateProduct,
		p.Name, p.Image, p.Purchase_price, p.Selling_price, p.Stock, id)
	if err != nil {
		return err
	}

	updateCategory := `
		UPDATE products_category
		SET category_id = $1
		WHERE product_id = $2
	`
	_, err = tx.Exec(context.Background(), updateCategory, categoryID, id)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
