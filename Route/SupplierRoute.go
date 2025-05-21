package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSupplier(c *gin.RouterGroup, db *gorm.DB) {
	r := SupplierDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)

	c.GET("/get-all-supplier", r.GetAllSupplier)
	c.POST("/create-supplier", r.CreateSupplier)
	c.PUT("/update-supplier", r.UpdateSupplier)
}