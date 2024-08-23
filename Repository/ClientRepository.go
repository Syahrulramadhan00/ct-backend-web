package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"gorm.io/gorm"
)

type (
	IClientRepository interface {
		GetAll() ([]Model.Client, error)
		Create(client *Dto.CreateClientRequest) error
		Update(client *Dto.UpdateClientRequest) error
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

func (h *ClientRepository) Create(client *Dto.CreateClientRequest) error {
	newClient := &Model.Client{
		Name:      client.Name,
		Place:     "",
		Telephone: "",
	}

	if err := h.DB.Create(&newClient).Error; err != nil {
		return err
	}

	return nil
}

func (h *ClientRepository) Update(client *Dto.UpdateClientRequest) error {
	clientData := &Model.Client{
		Name:      client.Name,
		Place:     "",
		Telephone: "",
	}

	if err := h.DB.Model(&Model.Client{}).Where("id = ?", client.ID).Updates(clientData).Error; err != nil {
		return err
	}

	return nil
}
