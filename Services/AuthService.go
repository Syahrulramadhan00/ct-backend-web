package Services

import (
	"ct-backend/Model"
	"ct-backend/Repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type (
	IAuthService interface {
		Login(email string, password string) (user *Model.User, token string, err error)
		Register(email string, password string, username string) (err error)
	}

	AuthService struct {
		repo Repository.IAuthRepository
	}
)

func AuthServiceProvider(repo Repository.IAuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (h *AuthService) Login(email string, password string) (user *Model.User, token string, err error) {

	// get user information
	if user, err = h.repo.GetUserInformation(email); err != nil {
		return nil, "", err
	} else if user == nil {
		return nil, "", errors.New("user not found")
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", err
	}

	// generate token
	expirationTime := time.Now().Add(3 * 30 * 24 * time.Hour)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub": user.ID,
		"exp": expirationTime.Unix(),
	})

	if token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET"))); err != nil {
		return nil, "", err
	}

	return user, token, err
}

func (h *AuthService) Register(email string, password string, username string) (err error) {
	var (
		hashedPassword []byte
	)

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err = h.repo.InsertUserInformation(email, string(hashedPassword)); err != nil {
		return err
	}

	return err
}
