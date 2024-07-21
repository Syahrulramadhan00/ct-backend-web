package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuth(c *gin.RouterGroup, db *gorm.DB) {
	r := AuthDI(db)

	//c.GET("/health", func(c *gin.Context) error {
	//	return c.SendString("api ok")
	//})

	c.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api ok",
		})
	})
	c.POST("/login", r.Login)
	c.POST("/register", r.Register)
}
