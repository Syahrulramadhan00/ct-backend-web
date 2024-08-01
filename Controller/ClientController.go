package Controller

import (
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IClientController interface {
		GetAllClient(ctx *gin.Context)
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
