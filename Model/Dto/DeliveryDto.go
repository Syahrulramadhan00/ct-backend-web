package Dto

type (
	CreateDeliveryProductRequest struct {
		DeliveryID int `json:"delivery_id" required:"true"`
		SalesID    int `json:"sales_id" required:"true"`
		Quantity   int `json:"quantity" required:"true"`
	}

	UpdateDeliveryProductRequest struct {
		ID              int `json:"id" required:"true"`
		Quantity        int `json:"quantity" required:"true"`
		CurrentQuantity int `json:"current_quantity" required:"true"`
	}

	DeleteDeliveryProductRequest struct {
		ID       int `json:"id" required:"true"`
		Quantity int `json:"quantity" required:"true"`
		SaleId   int `json:"sale_id" required:"true"`
	}

	UpdateSenderRequest struct {
		ID       int `json:"id" required:"true"`
		SenderId int `json:"sender" required:"true"`
	}

	UpdateDeliveryInformationRequest struct {
		DeliveryId int    `json:"delivery_id" required:"true"`
		Note       string `json:"note" required:"true"`
		Place      string `json:"place" required:"true"`
	}

	LockDeliveryOrderRequest struct {
		DeliveryId int `json:"delivery_id" required:"true"`
		InvoiceId  int `json:"invoice_id" required:"true"`
	}

	UpdateDeliveryStatusRequest struct {
		DeliveryId int `json:"delivery_id" required:"true"`
		Status     int `json:"status" required:"true"`
	}
)
