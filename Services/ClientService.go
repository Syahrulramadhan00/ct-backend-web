package Services

import (
	"ct-backend/Model"
	"ct-backend/Repository"
)

type (
	IClientService interface {
		GetAll() ([]Model.Client, error)
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
