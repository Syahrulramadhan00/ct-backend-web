package Controller

import (
	"ct-backend/Model/Dto"
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type (
	IInvoiceController interface {
		AddInvoice(ctx *gin.Context)
		GetAllInvoice(ctx *gin.Context)
		GetInvoiceById(ctx *gin.Context)
		LockInvoice(ctx *gin.Context)
	}

	InvoiceController struct {
		InvoiceService Services.IInvoiceService
	}
)

func InvoiceControllerProvider(invoiceService Services.IInvoiceService) *InvoiceController {
	return &InvoiceController{
		InvoiceService: invoiceService,
	}
}

func (h *InvoiceController) AddInvoice(ctx *gin.Context) {
	var request *Dto.CreateInvoiceRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := h.InvoiceService.AddInvoice(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) GetAllInvoice(ctx *gin.Context) {
	invoices, err := h.InvoiceService.GetAllInvoice()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    invoices,
	})
}

func (h *InvoiceController) GetInvoiceById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be a number",
		})
		return
	}

	invoice, err := h.InvoiceService.GetInvoiceById(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    invoice,
	})
}

func (h *InvoiceController) LockInvoice(ctx *gin.Context) {
	var id Dto.IdRequest

	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.LockInvoice(&id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
