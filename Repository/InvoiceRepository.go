package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"errors"
	"gorm.io/gorm"
)

type (
	IInvoiceRepository interface {
		GetAll() ([]Model.Invoice, error)
		GetById(id int) (Model.Invoice, error)
		GetLast() (*Model.Invoice, error)
		Create(request *Dto.CreateInvoiceRequest) (err error)
		Delete(request Dto.IdRequest) (err error)
		UpdateDocument(request Dto.UpdateDocumentRequest) (err error)
		AddSale(request *Dto.AddSaleRequest) (err error)
		GetAllSale(invoiceId int) ([]Model.Sale, error)
		UpdateSale(request *Dto.UpdateSaleRequest) (err error)
		DeleteSale(request Dto.IdRequest) (err error)
		UpdateFaktur(request *Dto.UpdateFakturRequest) (err error)
		UpdateMainInformation(request *Dto.UpdateMainInformationRequest) (err error)
		UpdateNote(request *Dto.UpdateNoteRequest) (err error)
		UpdateStatus(request *Dto.UpdateStatusRequest) (err error)
		UpdateNotSentSale(request *Dto.UpdateNotSentSaleRequest) (err error)
		GetSalesByInvoiceId(invoiceId int) ([]Model.Sale, error)
		GetAllForDelivery() ([]Model.Invoice, error)
		GetAllForReceipt() ([]Model.Invoice, error)
		UpdateInvoiceTotalPrice(invoiceID int) error
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

func (h *InvoiceRepository) GetLast() (invoice *Model.Invoice, err error) {
	if err := h.DB.Last(&invoice).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return invoice, nil
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
		InvoiceId:    request.InvoiceId,
		ProductId:    request.ProductId,
		Quantity:     request.Count,
		Price:        request.Price,
		NotSentCount: request.Count,
		SendStatus:   false,
	}

	if err := h.DB.Create(&sale).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateSale(request *Dto.UpdateSaleRequest) (err error) {
	if err := h.DB.
		Model(&Model.Sale{}).
		Where("id = ?", request.Id).
		Update("quantity", request.Count).
		Update("not_sent_count", request.Count).
		Update("price", request.Price).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) DeleteSale(request Dto.IdRequest) (err error) {
	tx := h.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var sale Model.Sale
	if err := tx.First(&sale, request.Id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&Model.Product{}).Where("id = ?", sale.ProductId).Update("stock", gorm.Expr("stock + ?", sale.Quantity)).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&sale).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (h *InvoiceRepository) UpdateFaktur(request *Dto.UpdateFakturRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("discount", request.Discount).
		Update("is_taxable", request.IsTaxable).
		Update("payment_term", request.PaymentTerm).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateMainInformation(request *Dto.UpdateMainInformationRequest) (err error) {
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

func (h *InvoiceRepository) UpdateNote(request *Dto.UpdateNoteRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("note", request.Note).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) UpdateStatus(request *Dto.UpdateStatusRequest) (err error) {
	if err := h.DB.
		Model(&Model.Invoice{}).
		Where("id = ?", request.InvoiceId).
		Update("invoice_status_id", request.InvoiceStatusId).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) GetAllSale(invoiceId int) (sales []Model.Sale, err error) {
	if err := h.DB.Preload("Product").Where("invoice_id = ?", invoiceId).Find(&sales).Error; err != nil {
		return nil, err
	}

	return sales, nil
}

func (h *InvoiceRepository) UpdateNotSentSale(request *Dto.UpdateNotSentSaleRequest) (err error) {
	if err := h.DB.
		Model(&Model.Sale{}).
		Where("id = ?", request.SaleId).
		Update("not_sent_count", gorm.Expr("not_sent_count - ?", request.Count)).Error; err != nil {
		return err
	}

	return nil
}

func (h *InvoiceRepository) GetSalesByInvoiceId(invoiceId int) (sales []Model.Sale, err error) {
	if err := h.DB.Preload("Product").Where("invoice_id = ?", invoiceId).Find(&sales).Error; err != nil {
		return nil, err
	}

	return sales, nil
}

func (h *InvoiceRepository) GetAllForDelivery() (invoices []Model.Invoice, err error) {
	if err := h.DB.Preload("Client").
		Joins("JOIN sales ON sales.invoice_id = invoices.id").
		Where("invoice_status_id = 2 AND sales.not_sent_count > 0").
		Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (h *InvoiceRepository) GetAllForReceipt() (invoices []Model.Invoice, err error) {
	if err := h.DB.Preload("Client").Where("invoice_status_id > ? AND invoice_status_id < ?", 2, 6).Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (h *InvoiceRepository) UpdateInvoiceTotalPrice(invoiceID int) error {
	var total float64

	err := h.DB.Model(&Model.Sale{}).
		Where("invoice_id = ?", invoiceID).
		Select("SUM(price * quantity) as total").
		Scan(&total).Error
	if err != nil {
		return err
	}

	err = h.DB.Model(&Model.Invoice{}).
		Where("id = ?", invoiceID).
		Update("total_price", total).Error
	if err != nil {
		return err
	}

	return nil
}
