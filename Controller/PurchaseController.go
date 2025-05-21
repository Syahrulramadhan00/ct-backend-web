package Controller

import (
	"ct-backend/Model/Dto"
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IPurchaseController interface {
		CreatePurchase(ctx *gin.Context)
		GetAllPurchase(ctx *gin.Context)
		PayDebt(ctx *gin.Context)
		DeletePurchase(ctx *gin.Context)
	}

	PurchaseController struct {
		PurchaseService Services.IPurchaseService
	}
)

func PurchaseControllerProvider(purchaseService Services.IPurchaseService) *PurchaseController {
	return &PurchaseController{
		PurchaseService: purchaseService,
	}
}

func (h *PurchaseController) CreatePurchase(ctx *gin.Context) {
	var request Dto.CreatePurchaseRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.PurchaseService.AddPurchase(&request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *PurchaseController) GetAllPurchase(ctx *gin.Context) {
	purchases, err := h.PurchaseService.GetAllPurchase()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    purchases,
	})
}

func (h *PurchaseController) PayDebt(ctx *gin.Context) {
	var id Dto.IdPurchaseRequest

	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.PurchaseService.PayDebt(id.Id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *PurchaseController) DeletePurchase(ctx *gin.Context) {
	var id Dto.IdPurchaseRequest

	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.PurchaseService.DeletePurchase(id.Id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
