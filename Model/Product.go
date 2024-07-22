package Model

import "time"

type Product struct {
	ID        uint
	Name      string
	Stock     int
	UpdatedAt time.Time
	CreatedAt time.Time
}
