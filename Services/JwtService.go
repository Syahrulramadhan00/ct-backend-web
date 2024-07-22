package Services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type (
	IJwtService interface {
		GenerateToken(userId int) (token string, err error)
		ParseToken(token string) (claims jwt.MapClaims, err error)
	}

	JwtService struct {
	}
)

func JwtServiceProvider() *JwtService {
	return &JwtService{}
}

func (h *JwtService) GenerateToken(userId int) (token string, err error) {
	expirationTime := time.Now().Add(3 * 30 * 24 * time.Hour)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub": userId,
		"exp": expirationTime.Unix(),
	})

	return jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (h *JwtService) ParseToken(authHeader string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println("error parse token : ", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("expired date : ", claims["exp"])
		return claims, nil
	} else {
		return nil, err
	}
}
