package persistence

import "github.com/google/uuid"

type Product struct {
	ProductID uuid.UUID `gorm:"primaryKey"`
	Name      string
	Image     string
	Stock     int64
	Price     int64
}

func (Product) TableName() string {
	return "product"
}
