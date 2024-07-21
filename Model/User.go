package Model

import "time"

type User struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	IsVerified bool
	UpdatedAt  time.Time
	CreatedAt  time.Time
	OtpCode    string
}
