package Services

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
)

type (
	IInvoiceService interface {
		AddInvoice(request *Dto.CreateInvoiceRequest) error
		GetAllInvoice() ([]Model.ShortInvoice, error)
		GetInvoiceById(id int) (Model.Invoice, error)
		LockInvoice(request *Dto.IdRequest) error
	}

	InvoiceService struct {
		InvoiceRepository Repository.IInvoiceRepository
	}
)

func InvoiceServiceProvider(invoiceRepository Repository.IInvoiceRepository) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository: invoiceRepository,
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
	return h.InvoiceRepository.UpdateStatus(Dto.UpdateStatusRequest{InvoiceId: request.Id, InvoiceStatusId: 2})
}
