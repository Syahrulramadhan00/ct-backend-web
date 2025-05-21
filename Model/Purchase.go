package Model

import "time"

type Purchase struct {
	ID         int
	ProductId  int
	SupplierId int 
	Count      int
	Price      int
	IsPaid     *bool
	ImagePath  *string
	CreatedAt  time.Time

	Product  Product  `gorm:"foreignKey:ProductId"`
}
