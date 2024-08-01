package Services

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
	"ct-backend/Utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

type (
	IAuthService interface {
		Login(email string, password string) (user *Model.User, token string, err error)
		Register(request *Dto.RegisterRequest) (err error)
		RequestOtp(email string) (err error)
		VerifyOtp(email string, otp string) (token string, err error)
		RequestForgotPasswordOtp(email string) (err error)
		VerifyForgotPasswordOtp(email string, otp string) (err error)
		ChangePassword(email string, password string) (err error)
	}

	AuthService struct {
		repo       Repository.IAuthRepository
		jwtService IJwtService
	}
)

func AuthServiceProvider(repo Repository.IAuthRepository, jwtService IJwtService) *AuthService {
	return &AuthService{repo: repo, jwtService: jwtService}
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
		return nil, "", errors.New("invalid password")
	}

	if token, err = h.jwtService.GenerateToken(int(user.ID)); err != nil {
		return nil, "", err
	}

	return user, token, err
}

func (h *AuthService) Register(request *Dto.RegisterRequest) (err error) {
	var (
		hashedPassword []byte
	)

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	request.Password = string(hashedPassword)

	if err = h.repo.InsertUserInformation(request); err != nil {
		return err
	}

	return err
}

func (h *AuthService) RequestOtp(email string) (err error) {
	randomNumber := rand.Intn(9000) + 1000
	if err = h.repo.SetOtpCode(email, strconv.Itoa(randomNumber)); err != nil {
		return err
	}

	if err = Utils.SendEmailToAdmin(
		"OTP Code for "+email,
		"Your OTP Code is: "+strconv.Itoa(randomNumber),
	); err != nil {
		return err
	}

	return err
}

func (h *AuthService) VerifyOtp(email string, otp string) (token string, err error) {
	var (
		user *Model.User
	)

	if user, err = h.repo.GetUserInformation(email); err != nil {
		return "", err
	} else if user == nil {
		return "", errors.New("user not found")
	}

	if user.OtpCode != otp {
		return "", errors.New("invalid otp")
	} else {

		if user.UpdatedAt.Before(time.Now().Add(-5 * time.Minute)) {
			return "", errors.New("otp expired")
		}

		if err = h.repo.SetVerificationStatus(email); err != nil {
			return "", err
		}
	}

	if token, err = h.jwtService.GenerateToken(int(user.ID)); err != nil {
		return "", err
	}

	return token, err
}

func (h *AuthService) RequestForgotPasswordOtp(email string) (err error) {
	randomNumber := rand.Intn(9000) + 1000
	if err = h.repo.SetOtpCode(email, strconv.Itoa(randomNumber)); err != nil {
		return err
	}

	if err = Utils.SendEmail(
		email,
		"FORGOT PASSWORD OTP for "+email,
		"Your OTP Code is: "+strconv.Itoa(randomNumber),
	); err != nil {
		return err
	}

	return err
}

func (h *AuthService) VerifyForgotPasswordOtp(email string, otp string) (err error) {
	var (
		user *Model.User
	)

	if user, err = h.repo.GetUserInformation(email); err != nil {
		return err
	} else if user == nil {
		return errors.New("user not found")
	}

	if user.OtpCode != otp {
		return errors.New("invalid otp")
	} else {
		if user.UpdatedAt.Before(time.Now().Add(-5 * time.Minute)) {
			return errors.New("otp expired")
		}
	}

	return err
}

func (h *AuthService) ChangePassword(email string, password string) (err error) {
	var (
		hashedPassword []byte
	)

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err = h.repo.ChangePassword(email, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}
