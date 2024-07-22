package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProduct(c *gin.RouterGroup, db *gorm.DB) {
	r := ProductDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)
	c.POST("/add-product", r.AddProduct)
	c.GET("/get-all-product", r.GetAllProduct)
	c.POST("/edit-name-product", r.EditNameProduct)
}
