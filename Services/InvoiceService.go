package Services

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
	"errors"
)

type (
	IInvoiceService interface {
		AddInvoice(request *Dto.CreateInvoiceRequest) error
		GetAllInvoice() ([]Model.ShortInvoice, error)
		GetInvoiceById(id int) (Model.Invoice, error)
		LockInvoice(request *Dto.IdRequest) error
		AddSaleToInvoice(request *Dto.AddSaleRequest) error
		UpdateSale(request *Dto.UpdateSaleRequest) error
		DeleteSale(request Dto.IdRequest) error
		GetAllSale(invoiceId int) ([]Model.Sale, error)
		UpdateFaktur(request *Dto.UpdateFakturRequest) error
		UpdateMainInformation(request *Dto.UpdateMainInformationRequest) error
		UpdateNote(request *Dto.UpdateNoteRequest) error
		UpdateStatus(request *Dto.UpdateStatusRequest) error
		DeleteInvoice(request Dto.IdRequest) error
	}

	InvoiceService struct {
		InvoiceRepository Repository.IInvoiceRepository
		ProductRepository Repository.IProductRepository
	}
)

func InvoiceServiceProvider(invoiceRepository Repository.IInvoiceRepository, ProductRepository Repository.IProductRepository) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository: invoiceRepository,
		ProductRepository: ProductRepository,
	}
}

func (h *InvoiceService) AddInvoice(request *Dto.CreateInvoiceRequest) error {
	// TODO : Add Random Invoice Code

	return h.InvoiceRepository.Create(request)
}

func (h *InvoiceService) GetAllInvoice() ([]Model.ShortInvoice, error) {
	invoices, err := h.InvoiceRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var shortInvoices []Model.ShortInvoice
	for _, invoice := range invoices {
		shortInvoices = append(shortInvoices, Model.ShortInvoice{
			ID:          invoice.ID,
			InvoiceCode: invoice.InvoiceCode,
			ClientName:  invoice.Client.Name,
			CreatedAt:   invoice.CreatedAt,
			Status:      invoice.GetStatusName(),
		})
	}

	return shortInvoices, nil
}

func (h *InvoiceService) GetInvoiceById(id int) (Model.Invoice, error) {
	return h.InvoiceRepository.GetById(id)
}

func (h *InvoiceService) LockInvoice(request *Dto.IdRequest) error {
	return h.InvoiceRepository.UpdateStatus(&Dto.UpdateStatusRequest{InvoiceId: request.Id, InvoiceStatusId: 2})
}

func (h *InvoiceService) AddSaleToInvoice(request *Dto.AddSaleRequest) error {
	err := h.ProductRepository.SumStockProduct(request.ProductId, request.Count*-1)
	if err != nil {
		return err
	}

	return h.InvoiceRepository.AddSale(request)
}

func (h *InvoiceService) UpdateSale(request *Dto.UpdateSaleRequest) error {
	count := request.Count - request.CurrentCount
	err := h.ProductRepository.SumStockProduct(request.ProductId, count*-1)
	if err != nil {
		return err
	}

	return h.InvoiceRepository.UpdateSale(request)
}

func (h *InvoiceService) DeleteSale(request Dto.IdRequest) error {
	return h.InvoiceRepository.DeleteSale(request)
}

func (h *InvoiceService) UpdateFaktur(request *Dto.UpdateFakturRequest) error {
	return h.InvoiceRepository.UpdateFaktur(request)
}

func (h *InvoiceService) UpdateMainInformation(request *Dto.UpdateMainInformationRequest) error {
	return h.InvoiceRepository.UpdateMainInformation(request)
}

func (h *InvoiceService) UpdateNote(request *Dto.UpdateNoteRequest) error {
	return h.InvoiceRepository.UpdateNote(request)
}

func (h *InvoiceService) UpdateStatus(request *Dto.UpdateStatusRequest) error {
	return h.InvoiceRepository.UpdateStatus(request)
}

func (h *InvoiceService) DeleteInvoice(request Dto.IdRequest) error {
	products, err := h.InvoiceRepository.GetAllSale(request.Id)
	if err != nil {
		return err
	}

	if len(products) > 0 {
		return errors.New("Invoice masih terdapat produk")
	}

	return h.InvoiceRepository.Delete(request)
}

func (h *InvoiceService) GetAllSale(invoiceId int) ([]Model.Sale, error) {
	return h.InvoiceRepository.GetAllSale(invoiceId)
}
