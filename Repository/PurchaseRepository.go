package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"gorm.io/gorm"
)

type (
	IPurchaseRepository interface {
		AddPurchase(request *Dto.CreatePurchaseRequest) error
		GetAllPurchase() (purchases []Model.Purchase, err error)
		ChangeIsPaidStatus(id int) (err error)
		DeletePurchase(id int) (err error)
		GetPurchaseById(id int) (purchase *Model.Purchase, err error)
	}

	PurchaseRepository struct {
		DB *gorm.DB
	}
)

func PurchaseRepositoryProvider(DB *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{
		DB: DB,
	}
}

func (h *PurchaseRepository) AddPurchase(request *Dto.CreatePurchaseRequest) (err error) {
	purchase := &Model.Purchase{
		ProductId: request.ProductId,
		SupplierId: request.SupplierId,
		Count:     request.Count,
		Price:     request.Price,
		IsPaid:    request.IsPaid,
		ImagePath: request.ImagePath,
	}

	if err := h.DB.Create(&purchase).Error; err != nil {
		return err
	}

	return nil
}

func (h *PurchaseRepository) GetAllPurchase() (purchases []Model.Purchase, err error) {
	if err := h.DB.Preload("Product").Preload("Supplier").Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, err
}

func (h *PurchaseRepository) ChangeIsPaidStatus(id int) (err error) {
	if err := h.DB.Model(&Model.Purchase{}).Where("id = ?", id).Update("is_paid", true).Error; err != nil {
		return err
	}

	return nil
}

func (h *PurchaseRepository) DeletePurchase(id int) (err error) {
	if err := h.DB.Where("id = ?", id).Delete(&Model.Purchase{}).Error; err != nil {
		return err
	}

	return nil
}

func (h *PurchaseRepository) GetPurchaseById(id int) (purchase *Model.Purchase, err error) {
	if err := h.DB.Where("id = ?", id).First(&purchase).Error; err != nil {
		return nil, err
	}

	return purchase, nil
}
