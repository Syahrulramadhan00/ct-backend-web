package Model

import "time"

type DeliveryOrder struct {
	ID        int
	InvoiceId int
	SenderId  *int
	OrderCode string
	Note      string
	Place     string
	PhotoPath string
	Status    int
	CreatedAt time.Time
	Invoice   Invoice `gorm:"foreignKey:InvoiceId"`
}

type ShortDeliveryOrder struct {
	ID         int
	OrderCode  string
	ClientName string
	CreatedAt  time.Time
	Status     string
}

func (d DeliveryOrder) GetStatusName() string {
	switch d.Status {
	case 1:
		return "initialized"
	case 2:
		return "process"
	case 3:
		return "done"
	default:
		return "initialized"
	}
}
