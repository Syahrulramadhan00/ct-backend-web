package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitInvoice(c *gin.RouterGroup, db *gorm.DB) {
	r := InvoiceDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)
	c.POST("/add-invoice", r.AddInvoice)
	c.GET("/get-all-invoice", r.GetAllInvoice)
	c.GET("/get-invoice/:id", r.GetInvoiceById)
	c.POST("/lock-invoice", r.LockInvoice)
	c.POST("/add-sale-to-invoice", r.AddSaleToInvoice)
}
