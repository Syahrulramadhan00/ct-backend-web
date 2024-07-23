package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		GetAll() ([]Model.Invoice, error)
		GetById(id int) (Model.Invoice, error)
		Create(request *Dto.CreateInvoiceRequest) (err error)
		Delete(request Dto.IdRequest) (err error)
		UpdateDocument(request Dto.UpdateDocumentRequest) (err error)
		AddSale(request *Dto.AddSaleRequest) (err error)
		UpdateSale(request Dto.AddSaleRequest) (err error)
		DeleteSale(request Dto.IdRequest) (err error)
		UpdateFaktur(request Dto.UpdateFakturRequest) (err error)
		UpdateMainInformation(request Dto.UpdateMainInformationRequest) (err error)
		UpdateNote(request Dto.UpdateNoteRequest) (err error)
		UpdateStatus(request Dto.UpdateStatusRequest) (err error)
	}

	InvoiceRepository struct {
		DB *gorm.DB
	}
)

func InvoiceRepositoryProvider(DB *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{
		DB: DB,
	}
}

func (h *InvoiceRepository) GetAll() (invoices []Model.Invoice, err error) {
	if err := h.DB.Preload("Client").Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (h *InvoiceRepository) GetById(id int) (invoice Model.Invoice, err error) {
	if err := h.DB.Preload("Client").Where("id = ?", id).First(&invoice).Error; err != nil {
		return Model.Invoice{}, err
	}

	return invoice, nil
}

func (h *InvoiceRepository) Create(request *Dto.CreateInvoiceRequest) (err error) {
	invoice := &Model.Invoice{
		ClientId:        request.ClientId,
		InvoiceCode:     request.InvoiceCode,
		InvoiceStatusId: 1,
	}

	if err := h.DB.Create(&invoice).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) Delete(request Dto.IdRequest) (err error) {
	if err := h.DB.Where("id = ?", request.Id).Delete(&Model.Invoice{}).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateDocument(request Dto.UpdateDocumentRequest) (err error) {
	documentType := "po_path"

	if request.DocumentType == "faktur" {
		documentType = "facture_path"
	}

	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update(documentType, request.DocumentPath).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) AddSale(request *Dto.AddSaleRequest) (err error) {
	sale := &Model.Sale{
		InvoiceId:  request.InvoiceId,
		ProductId:  request.ProductId,
		Quantity:   request.Count,
		Price:      request.Price,
		SendStatus: false,
	}

	if err := h.DB.Create(&sale).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateSale(request Dto.AddSaleRequest) (err error) {
	if err := h.DB.
		Model(&Model.Sale{}).
		Where("id = ?", request.Id).
		Update("quantity", request.Count).
		Update("price", request.Price).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) DeleteSale(request Dto.IdRequest) (err error) {
	if err := h.DB.Where("id = ?", request.Id).Delete(&Model.Sale{}).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateFaktur(request Dto.UpdateFakturRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("discount", request.Discount).
		Update("is_taxable", request.IsTaxable).
		Update("deadline", request.Deadline).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateMainInformation(request Dto.UpdateMainInformationRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("note", request.Note).
		Update("po_code", request.PoCode).
		Update("seller", request.Seller).
		Update("platform", request.Platform).
		Update("payment_method", request.PaymentMethod).
		Update("platform_description", request.PlatformDescription).
		Update("platform_number", request.PlatformNumber).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateNote(request Dto.UpdateNoteRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("note", request.Note).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateStatus(request Dto.UpdateStatusRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("invoice_status_id", request.InvoiceStatusId).Error; err != nil {
		return err
	}

	return nil
}
