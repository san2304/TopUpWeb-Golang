package models

import (
	"database/sql"
	"fmt"
	"log" // Import log package
	"time"
	"topup-game/config"
)

type InvoiceWithProduct struct {
	ProductID   int
	ProductName string
	Value       string
	Harga       int64
	Stock       int
	InvoiceID   int
	DataID      int
	Payment     string
	PhoneNumber int64
	Status      StatusType
	Time        time.Time
}

type StatusType string

const (
	Sukses     StatusType = "Sukses"
	Proses     StatusType = "Proses"
	BelumBayar StatusType = "Belum di Bayar"
	Gagal      StatusType = "Gagal"
)

func GetInvoiceWithProductByID(invoiceID int) (*InvoiceWithProduct, error) {
	fmt.Println("Debug: InvoiceID received:", invoiceID)

	query := `
		SELECT i.invoice_id, i.data_id, i.product_id, i.payment, i.time, i.phone_number, i.status,
			   f.product_name, f.value AS product_value, f.harga AS product_harga, f.stock AS product_stock
		FROM invoices i
		INNER JOIN freefire f ON i.product_id = f.product_id
		WHERE i.invoice_id = ?`

	row := config.DB.QueryRow(query, invoiceID)

	var invoiceWithProduct InvoiceWithProduct
	err := row.Scan(
		&invoiceWithProduct.InvoiceID,
		&invoiceWithProduct.DataID,
		&invoiceWithProduct.ProductID,
		&invoiceWithProduct.Payment,
		&invoiceWithProduct.Time,
		&invoiceWithProduct.PhoneNumber,
		&invoiceWithProduct.Status,
		&invoiceWithProduct.ProductName,
		&invoiceWithProduct.Value,
		&invoiceWithProduct.Harga,
		&invoiceWithProduct.Stock,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Handle no rows found as a special case
		}

		log.Println("Error: Unable to scan row:", err) // Log the error
		return nil, err
	}

	log.Println("Debug: InvoiceWithProduct retrieved from the database:", invoiceWithProduct)
	return &invoiceWithProduct, nil
}

func GetLatestInvoiceWithProduct() (*InvoiceWithProduct, error) {
	query := `
        SELECT i.invoice_id, i.data_id, i.product_id, i.payment, i.time, i.phone_number, i.status,
               f.product_name, f.value AS product_value, f.harga AS product_harga, f.stock AS product_stock
        FROM invoices i
        INNER JOIN freefire f ON i.product_id = f.product_id
        ORDER BY i.time DESC
        LIMIT 1`

	row := config.DB.QueryRow(query)

	var invoiceWithProduct InvoiceWithProduct
	err := row.Scan(
		&invoiceWithProduct.InvoiceID,
		&invoiceWithProduct.DataID,
		&invoiceWithProduct.ProductID,
		&invoiceWithProduct.Payment,
		&invoiceWithProduct.Time,
		&invoiceWithProduct.PhoneNumber,
		&invoiceWithProduct.Status,
		&invoiceWithProduct.ProductName,
		&invoiceWithProduct.Value,
		&invoiceWithProduct.Harga,
		&invoiceWithProduct.Stock,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		log.Println("Error: Unable to scan row:", err)
		return nil, err
	}

	log.Println("Debug: Latest InvoiceWithProduct retrieved from the database:", invoiceWithProduct)
	return &invoiceWithProduct, nil
}
