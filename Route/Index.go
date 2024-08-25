package Route

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(c *gin.Engine, db *gorm.DB, svc *s3.Client) {
	prefix := c.Group("/api/")
	InitAuth(prefix, db)
	InitProduct(prefix, db)
	InitPurchase(prefix, db)
	InitInvoice(prefix, db, svc)
	InitClient(prefix, db)
	InitDelivery(prefix, db)
	InitUser(prefix, db)
	InitReceipt(prefix, db)
}
