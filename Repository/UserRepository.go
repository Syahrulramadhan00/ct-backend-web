package Repository

import (
	"ct-backend/Model"
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		GetAllVerified() (users []Model.User, err error)
	}

	UserRepository struct {
		DB *gorm.DB
	}
)

func UserRepositoryProvider(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (h *UserRepository) GetAllVerified() (users []Model.User, err error) {
	if err := h.DB.Where("is_verified = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
