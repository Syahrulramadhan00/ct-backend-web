package Services

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
)

type (
	IClientService interface {
		GetAll() ([]Model.Client, error)
		Create(client *Dto.CreateClientRequest) error
		Update(client *Dto.UpdateClientRequest) error
	}

	ClientService struct {
		repo Repository.IClientRepository
	}
)

func ClientServiceProvider(repo Repository.IClientRepository) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

func (h *ClientService) GetAll() (clients []Model.Client, err error) {
	if clients, err = h.repo.GetAll(); err != nil {
		return nil, err
	}

	return clients, nil
}

func (h *ClientService) Create(client *Dto.CreateClientRequest) error {
	if err := h.repo.Create(client); err != nil {
		return err
	}

	return nil
}

func (h *ClientService) Update(client *Dto.UpdateClientRequest) error {
	if err := h.repo.Update(client); err != nil {
		return err
	}

	return nil
}
