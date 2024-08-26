package Dto

type (
	CreateClientRequest struct {
		Name      string `json:"name" binding:"required"`
		Address   string `json:"address" binding:"required"`
		Telephone string `json:"telephone" binding:"required"`
	}

	UpdateClientRequest struct {
		ID        int    `json:"id" binding:"required"`
		Name      string `json:"name" binding:"required"`
		Address   string `json:"address" binding:"required"`
		Telephone string `json:"telephone" binding:"required"`
	}
)
