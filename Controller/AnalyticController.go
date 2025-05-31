package Controller

import (
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type (
	IAnalyticController interface {
		GetRevenueStream(ctx *gin.Context)
		GetStockMonitoring(ctx *gin.Context)
		GetHighestSales(ctx *gin.Context)
		GetExpenses(ctx *gin.Context)
		GetTopSpenders(ctx *gin.Context)
		GetAvailableMonths(ctx *gin.Context)
		GetLatestBill(ctx *gin.Context)
	}

	AnalyticController struct {
		service Services.IAnalyticService
	}
)

func AnalyticControllerProvider(service Services.IAnalyticService) *AnalyticController {
	return &AnalyticController{service: service}
}

func (c *AnalyticController) GetRevenueStream(ctx *gin.Context) {
	startDateStr := ctx.Query("startDate") // e.g., "2024-01"
	endDateStr := ctx.Query("endDate")     // e.g., "2024-06"

	startDate, err := time.Parse("2006-01", startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid startDate format. Use YYYY-MM"})
		return
	}

	endDate, err := time.Parse("2006-01", endDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid endDate format. Use YYYY-MM"})
		return
	}

	data, err := c.service.GetRevenueStream(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (c *AnalyticController) GetExpenses(ctx *gin.Context) {
	startDateStr := ctx.Query("startDate") // e.g., "2024-01"
	endDateStr := ctx.Query("endDate")     // e.g., "2024-06"

	startDate, err := time.Parse("2006-01", startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid startDate format. Use YYYY-MM"})
		return
	}

	endDate, err := time.Parse("2006-01", endDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid endDate format. Use YYYY-MM"})
		return
	}

	data, err := c.service.GetExpenses(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (c *AnalyticController) GetStockMonitoring(ctx *gin.Context) {
	yearMonth := ctx.Query("yearMonth") // e.g., "2024-02"

	_, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid yearMonth format. Use YYYY-MM"})
		return
	}

	data, err := c.service.GetStockMonitoring(yearMonth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

func (c *AnalyticController) GetHighestSales(ctx *gin.Context) {
	yearMonth := ctx.Query("yearMonth") // e.g., "2024-02"

	parsedTime, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid yearMonth format. Use YYYY-MM"})
		return
	}

	startDate := parsedTime
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)

	data, err := c.service.GetHighestSales(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}


func (c *AnalyticController) GetTopSpenders(ctx *gin.Context) {
	yearMonth := ctx.Query("yearMonth") 

	_, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid yearMonth format. Use YYYY-MM"})
		return
	}

	data, err := c.service.GetTopSpenders(yearMonth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}


func (c *AnalyticController) GetAvailableMonths(ctx *gin.Context) {
	table := ctx.Query("table") 

	months, labels, err := c.service.GetAvailableMonths(table) // Updated to receive all values
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"months": months,
		"labels": labels, 
	})
}

func (c *AnalyticController) GetLatestBill(ctx *gin.Context) {
	data, err := c.service.GetLatestBill() // Call the service method
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
