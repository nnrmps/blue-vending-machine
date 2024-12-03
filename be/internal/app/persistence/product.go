package persistence

import "github.com/google/uuid"

type Product struct {
	ProductId uuid.UUID
	Name      string
	Image     string
	Stock     int64
	Price     float64
}
