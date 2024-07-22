package Middleware

import (
	"ct-backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type (
	ICommonMiddleware interface {
		AuthMiddleware(ctx *gin.Context)
	}

	CommonMiddleware struct {
		jwtService Services.IJwtService
	}
)

func CommonMiddlewareProvider(jwtService Services.IJwtService) *CommonMiddleware {
	return &CommonMiddleware{
		jwtService: jwtService,
	}
}

func (h *CommonMiddleware) Authentication(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token not found",
		})
		ctx.Abort()
		return
	}

	claims, err := h.jwtService.ParseToken(authHeader)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token expired",
		})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", claims["sub"])
	ctx.Next()
}
