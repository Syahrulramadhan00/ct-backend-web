package Dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OtpRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type OtpVerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	Otp   string `json:"otp" binding:"required"`
}
