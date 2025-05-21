package Dto

type (
	CreateSupplierRequest struct {
		Name      string `json:"name" binding:"required"`
		Company   string `json:"company" binding:"required"`
		Address   string `json:"address" binding:"required"`
		Telephone string `json:"telephone" binding:"required"`
	}

	UpdateSupplierRequest struct {
		ID        int    `json:"id" binding:"required"`
		Name      string `json:"name" binding:"required"`	
		Company   string `json:"company" binding:"required"`
		Address   string `json:"address" binding:"required"`
		Telephone string `json:"telephone" binding:"required"`
	}
)