package Dto

type LoginDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
