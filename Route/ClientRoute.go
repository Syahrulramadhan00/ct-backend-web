package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitClient(c *gin.RouterGroup, db *gorm.DB) {
	r := ClientDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)

	c.GET("/get-all-client", r.GetAllClient)
	c.POST("/create-client", r.CreateClient)
	c.PUT("/update-client", r.UpdateClient)
}
