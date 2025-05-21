package Services

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
)

type (
	ISupplierService interface {
		GetAll() ([]Model.Supplier, error)
		Create(supplier *Dto.CreateSupplierRequest) error
		Update(supplier *Dto.UpdateSupplierRequest) error
	}

	SupplierService struct {
		repo Repository.ISupplierRepository
	}
)

func SupplierServiceProvider(repo Repository.ISupplierRepository) *SupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (h *SupplierService) GetAll() (suppliers []Model.Supplier, err error) {
	if suppliers, err = h.repo.GetAll(); err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (h *SupplierService) Create(supplier *Dto.CreateSupplierRequest) error {
	if err := h.repo.Create(supplier); err != nil {
		return err
	}

	return nil
}

func (h *SupplierService) Update(supplier *Dto.UpdateSupplierRequest) error {
	if err := h.repo.Update(supplier); err != nil {
		return err
	}

	return nil
}