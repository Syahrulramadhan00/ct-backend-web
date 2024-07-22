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
		RequestOtp(ctx *gin.Context)
		VerifyOtp(ctx *gin.Context)
		RequestForgotPasswordOtp(ctx *gin.Context)
		VerifyForgotPasswordOtp(ctx *gin.Context)
		ChangePassword(ctx *gin.Context)
	}

	AuthController struct {
		service Services.IAuthService
	}
)

func AuthControllerProvider(service Services.IAuthService) *AuthController {
	return &AuthController{service: service}
}

func (h *AuthController) Login(ctx *gin.Context) {
	var loginRequest Dto.LoginRequest

	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, token, err := h.service.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// check verified status
	if !user.IsVerified {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not verified",
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
	var registerRequest Dto.RegisterRequest

	if err := ctx.ShouldBind(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.Register(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *AuthController) RequestOtp(ctx *gin.Context) {
	var otpRequest Dto.OtpRequest

	if err := ctx.ShouldBind(&otpRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.RequestOtp(otpRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "otp sent, please ask admin about the code",
	})
}

func (h *AuthController) VerifyOtp(ctx *gin.Context) {
	var otpVerificationRequest Dto.OtpVerificationRequest

	if err := ctx.ShouldBind(&otpVerificationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.VerifyOtp(otpVerificationRequest.Email, otpVerificationRequest.Otp)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "otp verified",
	})
}

func (h *AuthController) RequestForgotPasswordOtp(ctx *gin.Context) {
	var otpRequest Dto.OtpRequest

	if err := ctx.ShouldBind(&otpRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.RequestForgotPasswordOtp(otpRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "otp sent, please check your email",
	})
}

func (h *AuthController) VerifyForgotPasswordOtp(ctx *gin.Context) {
	var otpVerificationRequest Dto.OtpVerificationRequest

	if err := ctx.ShouldBind(&otpVerificationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.VerifyForgotPasswordOtp(otpVerificationRequest.Email, otpVerificationRequest.Otp)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "otp verified",
	})
}

func (h *AuthController) ChangePassword(ctx *gin.Context) {
	var changePasswordRequest Dto.LoginRequest

	if err := ctx.ShouldBind(&changePasswordRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.ChangePassword(changePasswordRequest.Email, changePasswordRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "password changed",
	})
}
