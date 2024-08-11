package Controller

import (
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
)

type (
	IUserController interface {
		GetAllVerified(ctx *gin.Context)
	}

	UserController struct {
		service Services.IUserService
	}
)

func UserControllerProvider(service Services.IUserService) *UserController {
	return &UserController{service: service}
}

func (h *UserController) GetAllVerified(ctx *gin.Context) {
	users, err := h.service.GetAllVerified()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    users,
	})
}
