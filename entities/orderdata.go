package entities

import "time"

type Invoice struct {
	InvoiceID   int
	DataID      int
	ProductID   int
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
