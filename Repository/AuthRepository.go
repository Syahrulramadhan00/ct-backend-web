package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"errors"
	"gorm.io/gorm"
)

type (
	IAuthRepository interface {
		InsertUserInformation(request *Dto.RegisterRequest) (err error)
		GetUserInformation(email string) (user *Model.User, err error)
		SetOtpCode(email string, otp string) (err error)
		SetVerificationStatus(email string) (err error)
	}

	AuthRepository struct {
		DB *gorm.DB
	}
)

func AuthRepositoryProvider(DB *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: DB,
	}
}

func (h *AuthRepository) InsertUserInformation(request *Dto.RegisterRequest) (err error) {
	if user, _ := h.GetUserInformation(request.Email); user != nil {
		return errors.New("email already exist")
	}

	user := &Model.User{
		Name:     request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if err = h.DB.Create(&user).Error; err != nil {
		return err
	}

	return err
}

func (h *AuthRepository) GetUserInformation(email string) (user *Model.User, err error) {
	if err = h.DB.
		Find(&Model.User{}, "email = ?", email).
		First(&Model.User{}).
		Take(&user).Error; err != nil {
		return nil, err
	}

	return user, err
}

func (h *AuthRepository) SetOtpCode(email string, otp string) (err error) {
	if err := h.DB.Model(&Model.User{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"otp_code": otp,
		}).Error; err != nil {
		return err
	}

	return err
}

func (h *AuthRepository) SetVerificationStatus(email string) (err error) {
	if err = h.DB.Model(&Model.User{}).
		Where("email = ?", email).
		Update("is_verified", true).Error; err != nil {
		return err
	}

	return err
}
