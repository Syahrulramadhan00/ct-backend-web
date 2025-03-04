// package Services

// import (
// 	"fmt"
// 	"github.com/golang-jwt/jwt/v5"
// 	"os"
// 	"time"
// )

// type (
// 	IJwtService interface {
// 		GenerateToken(userId int) (token string, err error)
// 		ParseToken(token string) (claims jwt.MapClaims, err error)
// 	}

// 	JwtService struct {
// 	}
// )

// func JwtServiceProvider() *JwtService {
// 	return &JwtService{}
// }

// func (h *JwtService) GenerateToken(userId int) (token string, err error) {
// 	expirationTime := time.Now().Add(3 * 30 * 24 * time.Hour)

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
// 		"sub": userId,
// 		"exp": expirationTime.Unix(),
// 	})

// 	return jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
// }

// func (h *JwtService) ParseToken(authHeader string) (claims jwt.MapClaims, err error) {
// 	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})
// 	if err != nil {
// 		fmt.Println("error parse token : ", err)
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		return claims, nil
// 	} else {
// 		return nil, err
// 	}
// }


package Services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

type (
	IJwtService interface {
		GenerateToken(userId int) (token string, err error)
		ParseToken(authHeader string) (claims jwt.MapClaims, err error)
	}

	JwtService struct{}
)

func JwtServiceProvider() *JwtService {
	return &JwtService{}
}

// GenerateToken creates a JWT token with userId and expiration time
func (h *JwtService) GenerateToken(userId int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	expirationTime := time.Now().Add(3 * 30 * 24 * time.Hour) // 3 months

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": expirationTime.Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

// ParseToken validates and extracts claims from the JWT token
func (h *JwtService) ParseToken(authHeader string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not set in environment variables")
	}

	// Check if token starts with "Bearer "
	tokenString := strings.TrimSpace(authHeader)
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	// Handle parsing errors
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
