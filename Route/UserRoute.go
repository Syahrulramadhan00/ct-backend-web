package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUser(c *gin.RouterGroup, db *gorm.DB) {
	r := UserDI(db)
	m := CommonMiddlewareDI()

	c.Use(m.Authentication)
	c.GET("/get-all-verified", r.GetAllVerified)
}
