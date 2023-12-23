package models

import (
	"database/sql"
	"topup-game/config"
	"topup-game/entities"
)

type Invoice entities.Invoice

func GetInvoiceByNumber(invoiceNumber int) (*Invoice, error) {
	row := config.DB.QueryRow("SELECT ID, InvoiceNumber, PhoneNumber, Harga, Status, CreatedAt FROM invoices WHERE InvoiceNumber = ?", invoiceNumber)

	var invoice Invoice
	err := row.Scan(&invoice.ID, &invoice.InvoiceNumber, &invoice.PhoneNumber, &invoice.Harga, &invoice.Status, &invoice.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Handle no rows found as a special case
		}
		return nil, err
	}

	return &invoice, nil
}
