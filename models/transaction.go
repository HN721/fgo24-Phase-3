package models

import (
	"context"
	"fmt"
	"nastha-test/utils"
	"time"

	"github.com/jackc/pgx/v5"
)

type Transactions struct {
	TransactionID   int       `json:"transaction_id"`
	ProductID       int       `json:"product_id"`
	UsersID         int       `json:"users_id"`
	Type            string    `json:"type"`
	Quantity        int       `json:"quantity"`
	TotalPrice      float64   `json:"total_price"`
	TransactionDate time.Time `json:"transaction_date"`
}
type TransactionHistory struct {
	KodeBarang       string    `json:"kode_barang"`
	NamaBarang       string    `json:"nama_barang"`
	KategoriBarang   string    `json:"kategori_barang"`
	BarangMasuk      int       `json:"barang_masuk"`
	BarangKeluar     int       `json:"barang_keluar"`
	HargaBeliSatuan  float64   `json:"harga_beli_satuan"`
	HargaJualSatuan  float64   `json:"harga_jual_satuan"`
	TotalPembelian   float64   `json:"total_pembelian"`
	TotalPenjualan   float64   `json:"total_penjualan"`
	StokTersedia     int       `json:"stok_tersedia"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
}

func CreateTransaction(trx Transactions) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	product, err := GetProductByID(trx.ProductID)
	if err != nil {
		return err
	}

	trx.TotalPrice = float64(trx.Quantity) * product.Selling_price

	query := `
	INSERT INTO transactions (
		product_id, users_id, type, quantity, total_price, transaction_date
	) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = conn.Exec(
		context.Background(),
		query,
		trx.ProductID,
		trx.UsersID,
		trx.Type,
		trx.Quantity,
		trx.TotalPrice,
		time.Now(),
	)
	if err != nil {
		return err
	}

	var stockUpdateQuery string
	if trx.Type == "out" {
		if product.Stock < trx.Quantity {
			return fmt.Errorf("stok tidak mencukupi")
		}
		stockUpdateQuery = `UPDATE products SET stock = stock - $1 WHERE id = $2`
	} else if trx.Type == "in" {
		stockUpdateQuery = `UPDATE products SET stock = stock + $1 WHERE id = $2`
	} else {
		return fmt.Errorf("tipe transaksi tidak valid: gunakan 'in' atau 'out'")
	}

	_, err = conn.Exec(context.Background(), stockUpdateQuery, trx.Quantity, trx.ProductID)
	if err != nil {
		return err
	}

	return nil
}
func GetTransactionHistory() ([]TransactionHistory, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
	SELECT 
		p.kode AS kode_barang,
		p.name AS nama_barang,
		c.name AS kategori_barang,
		CASE WHEN t.type = 'in' THEN t.quantity ELSE 0 END AS barang_masuk,
		CASE WHEN t.type = 'out' THEN t.quantity ELSE 0 END AS barang_keluar,
		p.purchase_price AS harga_beli_satuan,
		p.selling_price AS harga_jual_satuan,
		CASE WHEN t.type = 'in' THEN t.total_price ELSE 0 END AS total_pembelian,
		CASE WHEN t.type = 'out' THEN t.total_price ELSE 0 END AS total_penjualan,
		p.stock AS stok_tersedia,
		t.transaction_date AS tanggal_transaksi
	FROM transactions t
	JOIN products p ON p.id = t.product_id
	JOIN products_category pc ON p.id = pc.product_id
	JOIN category c ON c.id = pc.category_id
	ORDER BY t.transaction_date DESC
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows[TransactionHistory](rows, pgx.RowToStructByName)
}
