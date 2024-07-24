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
		AddSaleToInvoice(ctx *gin.Context)
		UpdateSale(ctx *gin.Context)
		DeleteSale(ctx *gin.Context)
		UpdateFaktur(ctx *gin.Context)
		UpdateMainInformation(ctx *gin.Context)
		UpdateNote(ctx *gin.Context)
		UpdateStatus(ctx *gin.Context)
		GetAllSale(ctx *gin.Context)
		DeleteInvoice(ctx *gin.Context)
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

func (h *InvoiceController) AddSaleToInvoice(ctx *gin.Context) {
	var request *Dto.AddSaleRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.AddSaleToInvoice(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) UpdateSale(ctx *gin.Context) {
	var request *Dto.UpdateSaleRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.UpdateSale(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) DeleteSale(ctx *gin.Context) {
	var id Dto.IdRequest

	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.DeleteSale(id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) UpdateFaktur(ctx *gin.Context) {
	var request *Dto.UpdateFakturRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.UpdateFaktur(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) UpdateMainInformation(ctx *gin.Context) {
	var request *Dto.UpdateMainInformationRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.UpdateMainInformation(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) UpdateNote(ctx *gin.Context) {
	var request *Dto.UpdateNoteRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.UpdateNote(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) UpdateStatus(ctx *gin.Context) {
	var request *Dto.UpdateStatusRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.UpdateStatus(request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *InvoiceController) GetAllSale(ctx *gin.Context) {
	invoiceId, err := strconv.Atoi(ctx.Param("invoiceId"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invoiceId must be a number",
		})
		return
	}

	sales, err := h.InvoiceService.GetAllSale(invoiceId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    sales,
	})
}

func (h *InvoiceController) DeleteInvoice(ctx *gin.Context) {
	var id Dto.IdRequest

	if err := ctx.ShouldBind(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.InvoiceService.DeleteInvoice(id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
