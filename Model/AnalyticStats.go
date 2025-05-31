package Model


type Revenue struct {
	Month  string  `gorm:"column:month"`
	Total  float64 `gorm:"column:total_revenue"`
}

type Stock struct {
	Name  string `gorm:"column:name"`
	Stock int    `gorm:"column:stock"`
}

type HighestSales struct {
	ProductName string `gorm:"column:product_name"`
	Total   int    `gorm:"column:total"`
}


type Expenses struct {    
	Month string  `gorm:"column:month"`
	Total float64 `gorm:"column:total_expenses"` 
}

type TopSpenders struct { 
	Name  string `gorm:"column:name"`
	Total float64 `gorm:"column:total"`
}

type LatestBill struct {
	CodeInvoices string `gorm:"column:invoice_code"`
	ClientName   string `gorm:"column:client_name"`
	ClientContact string `gorm:"column:client_contact"`
	TotalAmount  int    `gorm:"column:total_amount"`
	PaymentStatus string `gorm:"column:payment_status"`
}