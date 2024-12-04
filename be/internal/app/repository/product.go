package repository

import (
	"context"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetList(ctx context.Context, tx *gorm.DB) []persistence.Product
}

type productRepository struct {
}

func (p productRepository) GetList(ctx context.Context, tx *gorm.DB) []persistence.Product {
	product := make([]persistence.Product, 0)
	tx.Find(&product)

	return product
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
