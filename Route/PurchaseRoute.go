package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPurchase(c *gin.RouterGroup, db *gorm.DB) {
	r := PurchaseDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)
	c.GET("/get-all-purchase", r.GetAllPurchase)
	c.POST("/add-purchase", r.CreatePurchase)
	c.POST("/pay-debt", r.PayDebt)
	c.DELETE("/delete-purchase", r.DeletePurchase)
}
