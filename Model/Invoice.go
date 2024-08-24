package Model

import (
	"time"
)

type Invoice struct {
	ID                  int
	InvoiceStatusId     int
	ClientId            int
	InvoiceCode         string
	Description         string
	PaymentMethod       string
	PoCode              string
	Seller              string
	PlatformNumber      string
	Platform            string
	PlatformDescription string
	Tax                 int
	Note                string
	FacturePath         string
	PoPath              string
	Discount            int
	IsTaxable           bool
	PaymentTerm         int
	TotalPrice          int
	ProjectName         string
	Date                time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Client              Client `gorm:"foreignKey:ClientId"`
}

type ShortInvoice struct {
	ID          int
	InvoiceCode string
	ClientName  string
	CreatedAt   time.Time
	Status      string
	StatusId    int
}

func (i Invoice) GetStatusName() string {
	switch i.InvoiceStatusId {
	case 1:
		return "initialized"
	case 2:
		return "finalized"
	case 3:
		return "sended"
	case 4:
		return "preorder_created"
	case 5:
		return "faktur_created"
	case 6:
		return "inserted_into_receipt"
	case 7:
		return "paid"
	case 8:
		return "done"
	default:
		return "initialized"
	}
}
