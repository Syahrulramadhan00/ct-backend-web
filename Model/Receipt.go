package Model

import "time"

type Receipt struct {
	ID        int
	ClientId  int
	Photo     string
	Number    int
	Status    int
	Client    Client `gorm:"foreignKey:ClientId"`
	CreatedAt time.Time
}

type ReceiptInvoice struct {
	ID        int
	ReceiptId int
	InvoiceId int
	Invoice   Invoice
}
