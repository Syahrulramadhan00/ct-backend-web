package Dto

type ChartData struct {
	Labels   []string   `json:"labels"`  
	Datasets []Dataset `json:"datasets"` 
}

type Dataset struct {
	Label string  `json:"label"` 
	Data  []float64   `json:"data"` 
}

type LatestBillDTO struct {
	InvoiceCode   string `json:"invoice_code"`
	ClientName    string `json:"client_name"`
	ClientContact string `json:"client_contact"`
	TotalAmount   int    `json:"total_amount"`
	PaymentStatus string `json:"payment_status"`
}