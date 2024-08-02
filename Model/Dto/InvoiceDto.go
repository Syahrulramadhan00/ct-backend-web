package Dto

type CreateInvoiceRequest struct {
	ClientId    int    `json:"client_id" binding:"required"`
	InvoiceCode string `json:"invoice_code"`
}

type UpdateDocumentRequest struct {
	InvoiceId    int    `json:"invoice_id" binding:"required"`
	DocumentType string `json:"document_type" binding:"required"`
	DocumentPath string `json:"document_path" binding:"required"`
}

type AddSaleRequest struct {
	Id        int `json:"id"`
	InvoiceId int `json:"invoice_id" binding:"required"`
	ProductId int `json:"product_id" binding:"required"`
	Count     int `json:"count" binding:"required"`
	Price     int `json:"price" binding:"required"`
}

type UpdateSaleRequest struct {
	Id           int `json:"id"`
	ProductId    int `json:"product_id" binding:"required"`
	CurrentCount int `json:"current_count" binding:"required"`
	Count        int `json:"count" binding:"required"`
	Price        int `json:"price" binding:"required"`
}

type UpdateFakturRequest struct {
	InvoiceId   int  `json:"invoice_id" binding:"required"`
	Discount    int  `json:"discount"`
	PaymentTerm int  `json:"payment_term" binding:"required"`
	IsTaxable   bool `json:"is_taxable" binding:"required"`
}

type UpdateMainInformationRequest struct {
	InvoiceId           int    `json:"invoice_id" binding:"required"`
	PoCode              string `json:"po_code" binding:"required"`
	Note                string `json:"note" binding:"required"`
	Seller              string `json:"seller" binding:"required"`
	Platform            string `json:"platform" binding:"required"`
	PaymentMethod       string `json:"payment_method" binding:"required"`
	PlatformDescription string `json:"platform_description" binding:"required"`
	PlatformNumber      string `json:"platform_number" binding:"required"`
}

type UpdateNoteRequest struct {
	InvoiceId int    `json:"invoice_id" binding:"required"`
	Note      string `json:"note" binding:"required"`
}

type UpdateStatusRequest struct {
	InvoiceId       int `json:"invoice_id" binding:"required"`
	InvoiceStatusId int `json:"invoice_status_id" binding:"required"`
}
