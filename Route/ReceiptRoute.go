package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitReceipt(c *gin.RouterGroup, db *gorm.DB) {
	r := ReceiptDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)
	c.GET("/get-all-receipt", r.GetReceipts)
	c.POST("/add-receipt", r.CreateReceipt)
	c.GET("/get-receipt/:id", r.GetReceiptById)
	c.GET("/get-invoice-by-receipt/:id", r.GetReceiptByInvoiceId)
	c.POST("/add-invoice-to-receipt", r.AddInvoiceToReceipt)
	c.POST("/lock-receipt", r.LockReceipt)
	c.DELETE("/delete-invoice-receipt/:id", r.DeleteReceiptInvoice)
	c.GET("/get-receipt-invoices/:id", r.GetAvailableInvoices)
	c.GET("/get-client-receipts", r.GetClientReceipts)
}
