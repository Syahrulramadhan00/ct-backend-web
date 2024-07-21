package Repository

import (
	"ct-backend/Model"
	"gorm.io/gorm"
)

type (
	IAuthRepository interface {
		InsertUserInformation(email string, password string) (err error)
		GetUserInformation(email string) (user *Model.User, err error)
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

func (h *AuthRepository) InsertUserInformation(email string, password string) (err error) {
	if user, err := h.GetUserInformation(email); err != nil {
		return err
	} else if user != nil {
		return err
	}

	if err = h.DB.Create(&Model.User{Email: email, Password: password}).Error; err != nil {
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
