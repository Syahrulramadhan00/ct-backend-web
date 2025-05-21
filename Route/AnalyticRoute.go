package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAnalytic(c *gin.RouterGroup, db *gorm.DB) {
	r := AnalyticDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)

	c.GET("/get-revenue-stream", r.GetRevenueStream)
	c.GET("/get-stock-monitoring", r.GetStockMonitoring)
	c.GET("/get-highest-sales", r.GetHighestSales)
	c.GET("/get-expenses", r.GetExpenses)
	c.GET("/get-top-spenders", r.GetTopSpenders)
	c.GET("/get-available-months", r.GetAvailableMonths)
}
