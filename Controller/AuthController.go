package Controller

import (
	"ct-backend/Model/Dto"
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IAuthController interface {
		Login(ctx *gin.Context)
		Register(ctx *gin.Context)
	}

	AuthController struct {
		service Services.IAuthService
	}
)

func AuthControllerProvider(service Services.IAuthService) *AuthController {
	return &AuthController{service: service}
}

func (h *AuthController) Login(ctx *gin.Context) {
	var loginData Dto.LoginDto

	if err := ctx.ShouldBind(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, token, err := h.service.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (h *AuthController) Register(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	username := ctx.PostForm("username")

	err := h.service.Register(email, password, username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
