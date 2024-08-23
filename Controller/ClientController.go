package Controller

import (
	"ct-backend/Model/Dto"
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IClientController interface {
		GetAllClient(ctx *gin.Context)
		CreateClient(ctx *gin.Context)
		UpdateClient(ctx *gin.Context)
	}

	ClientController struct {
		ClientService Services.IClientService
	}
)

func ClientControllerProvider(clientService Services.IClientService) *ClientController {
	return &ClientController{
		ClientService: clientService,
	}
}

func (h *ClientController) GetAllClient(ctx *gin.Context) {
	clients, err := h.ClientService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    clients,
	})
}

func (h *ClientController) CreateClient(ctx *gin.Context) {
	var request Dto.CreateClientRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.ClientService.Create(&request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (h *ClientController) UpdateClient(ctx *gin.Context) {
	var request Dto.UpdateClientRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.ClientService.Update(&request); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
