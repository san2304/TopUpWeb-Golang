package models

import (
	"database/sql"
	"fmt"
	"topup-game/config"
	"topup-game/entities"
)

type Invoice entities.Invoice

func GetInvoiceByID(invoiceID int) (*Invoice, error) {
	fmt.Println("Debug: InvoiceID received:", invoiceID)
	row := config.DB.QueryRow("SELECT invoice_id, data_id, product_id, payment, time, phone_number, status FROM invoices WHERE invoice_id = ?", invoiceID)

	var invoice Invoice
	err := row.Scan(&invoice.InvoiceID, &invoice.DataID, &invoice.ProductID, &invoice.Payment, &invoice.Time, &invoice.PhoneNumber, &invoice.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Handle no rows found as a special case
		}
		return nil, err
	}

	fmt.Println("Debug: Invoice retrieved from the database:", invoice)
	return &invoice, nil
}
