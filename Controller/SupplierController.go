package Controller

import (
	"ct-backend/Model/Dto"
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	ISupplierController interface {
		GetAllSupplier(ctx *gin.Context)
		CreateSupplier(ctx *gin.Context)
		UpdateSupplier(ctx *gin.Context)
	}

	SupplierController struct {
		SupplierService Services.ISupplierService
	}
)

func SupplierControllerProvider(supplierService Services.ISupplierService) *SupplierController {
	return &SupplierController{
		SupplierService: supplierService,
	}
}

func (h *SupplierController) GetAllSupplier(ctx *gin.Context) {
	suppliers, err := h.SupplierService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    suppliers,
	})
}

func (h *SupplierController) CreateSupplier(ctx *gin.Context) {
	var request Dto.CreateSupplierRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.SupplierService.Create(&request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (h *SupplierController) UpdateSupplier(ctx *gin.Context) {
	var request Dto.UpdateSupplierRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.SupplierService.Update(&request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}