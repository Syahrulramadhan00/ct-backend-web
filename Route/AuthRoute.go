package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuth(c *gin.RouterGroup, db *gorm.DB) {
	r := AuthDI(db)

	c.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api ok",
		})
	})
	c.POST("/login", r.Login)
	c.POST("/register", r.Register)
	c.POST("/request-otp", r.RequestOtp)
	c.POST("/verify-otp", r.VerifyOtp)
	c.POST("/request-forgot-password-otp", r.RequestForgotPasswordOtp)
	c.POST("/verify-forgot-password-otp", r.VerifyForgotPasswordOtp)
	c.POST("/change-password", r.ChangePassword)
}
