package Services

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
	"time"
)

type (
	IReceiptService interface {
		GetAllReceipt() (receipts []Model.Receipt, err error)
		GetReceiptById(id int) (receipt *Model.Receipt, err error)
		CreateReceipt(clientId int) (m *Model.Receipt, err error)
		GetInvoiceByReceiptId(id int) ([]Model.ReceiptInvoice, error)
		AddInvoiceReceipt(receiptInvoice *Dto.ReceiptInvoiceRequest) (*Model.ReceiptInvoice, error)
		LockReceipt(receiptId int) error
		DeleteReceiptInvoice(id int) error
	}

	ReceiptService struct {
		ReceiptRepository Repository.IReceiptRepository
	}
)

func ReceiptServiceProvider(receiptRepository Repository.IReceiptRepository) *ReceiptService {
	return &ReceiptService{
		ReceiptRepository: receiptRepository,
	}
}

func (h *ReceiptService) GetAllReceipt() (receipts []Model.Receipt, err error) {
	return h.ReceiptRepository.GetReceipts()
}

func (h *ReceiptService) GetReceiptById(id int) (receipt *Model.Receipt, err error) {
	return h.ReceiptRepository.GetReceiptById(id)
}

func (h *ReceiptService) CreateReceipt(clientId int) (m *Model.Receipt, err error) {
	lastReceipt, _ := h.ReceiptRepository.GetLast()

	receipt := &Model.Receipt{
		ClientId: clientId,
		Status:   1,
	}

	receipt.Number = 1

	if lastReceipt != nil && lastReceipt.CreatedAt.Year() == time.Now().Year() {
		receipt.Number = lastReceipt.Number + 1
	}

	return h.ReceiptRepository.CreateReceipt(receipt)
}

func (h *ReceiptService) GetInvoiceByReceiptId(id int) ([]Model.ReceiptInvoice, error) {
	return h.ReceiptRepository.GetInvoiceByReceiptId(id)
}

func (h *ReceiptService) AddInvoiceReceipt(receiptInvoice *Dto.ReceiptInvoiceRequest) (*Model.ReceiptInvoice, error) {
	return h.ReceiptRepository.AddInvoiceReceipt(receiptInvoice)
}

func (h *ReceiptService) LockReceipt(receiptId int) error {
	return h.ReceiptRepository.LockReceipt(receiptId)
}

func (h *ReceiptService) DeleteReceiptInvoice(id int) error {
	return h.ReceiptRepository.DeleteInvoiceReceipt(id)
}
