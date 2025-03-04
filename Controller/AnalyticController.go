package Controller

import (
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IAnalyticController interface {
		GetRevenueStream(ctx *gin.Context)
		GetStockMonitoring(ctx *gin.Context)
		GetHighestSales(ctx *gin.Context)
	}

	AnalyticController struct {
		service Services.IAnalyticService
	}
)

func AnalyticControllerProvider(service Services.IAnalyticService) *AnalyticController {
	return &AnalyticController{service: service}
}

func (c *AnalyticController) GetRevenueStream(ctx *gin.Context) {
	data, err := c.service.GetRevenueStream()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (c *AnalyticController) GetStockMonitoring(ctx *gin.Context) {
	data, err := c.service.GetStockMonitoring()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (c *AnalyticController) GetHighestSales(ctx *gin.Context) {
	data, err := c.service.GetHighestSales()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
