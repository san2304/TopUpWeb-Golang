package entities

import "time"

type Invoice struct {
	ID            int
	InvoiceNumber int
	PhoneNumber   int
	Harga         int64
	Status        StatusType
	CreatedAt     time.Time
}

type StatusType string

const (
	Sukses     StatusType = "Sukses"
	Proses     StatusType = "Proses"
	BelumBayar StatusType = "Belum di Bayar"
	Gagal      StatusType = "Gagal"
)
