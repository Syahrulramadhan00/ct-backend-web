package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"gorm.io/gorm"
)

type (
	ISupplierRepository interface {
		GetAll() ([]Model.Supplier, error)
		Create(supplier *Dto.CreateSupplierRequest) error
		Update(supplier *Dto.UpdateSupplierRequest) error
	}


	SupplierRepository struct {
		DB *gorm.DB
	}
)

func SupplierRepositoryProvider(DB *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		DB: DB,
	}
}

func (h *SupplierRepository) GetAll() (suppliers []Model.Supplier, err error) {
	if err := h.DB.Find(&suppliers).Error; err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (h *SupplierRepository) Create(supplier *Dto.CreateSupplierRequest) error {
	newSupplier := &Model.Supplier{
		Name:      supplier.Name,
		Company:   supplier.Company,
		Address:   supplier.Address,
		Telephone: supplier.Telephone,
	}

	if err := h.DB.Create(&newSupplier).Error; err != nil {
		return err
	}

	return nil
}

func (h *SupplierRepository) Update(supplier *Dto.UpdateSupplierRequest) error {
	supplierData := &Model.Supplier{
		Name:      supplier.Name,
		Company:   supplier.Company,
		Address:   supplier.Address,
		Telephone: supplier.Telephone,
	}

	if err := h.DB.Model(&Model.Supplier{}).Where("id = ?", supplier.ID).Updates(supplierData).Error; err != nil {
		return err
	}

	return nil
}