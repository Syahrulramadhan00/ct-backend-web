package Dto

type (
	CreateClientRequest struct {
		Name string `json:"name" binding:"required"`
	}

	UpdateClientRequest struct {
		ID   int    `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}
)
