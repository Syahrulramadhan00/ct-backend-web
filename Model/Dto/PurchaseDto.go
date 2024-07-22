package Dto

type CreatePurchaseRequest struct {
	ProductId int     `json:"product_id" binding:"required"`
	Count     int     `json:"count" binding:"required"`
	Price     int     `json:"price" binding:"required"`
	IsPaid    *bool   `json:"is_paid"`
	ImagePath *string `json:"image_path"`
}

type IdPurchaseRequest struct {
	Id int `json:"id" binding:"required"`
}
