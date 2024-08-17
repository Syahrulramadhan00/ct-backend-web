package Dto

type ReceiptInvoiceRequest struct {
	ReceiptId int `json:"receipt_id" binding:"required"`
	InvoiceId int `json:"invoice_id" binding:"required"`
}
