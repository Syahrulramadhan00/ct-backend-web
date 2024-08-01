package Repository

import (
	"ct-backend/Model"
	"gorm.io/gorm"
)

type (
	IClientRepository interface {
		GetAll() ([]Model.Client, error)
	}

	ClientRepository struct {
		DB *gorm.DB
	}
)

func ClientRepositoryProvider(DB *gorm.DB) *ClientRepository {
	return &ClientRepository{
		DB: DB,
	}
}

func (h *ClientRepository) GetAll() (clients []Model.Client, err error) {
	if err := h.DB.Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}
