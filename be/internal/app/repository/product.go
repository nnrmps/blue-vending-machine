package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(ctx context.Context, tx *gorm.DB, productID uuid.UUID) persistence.Product
}

type productRepository struct {
}

func (p productRepository) FindByID(ctx context.Context, tx *gorm.DB, productID uuid.UUID) persistence.Product {
	//TODO implement me
	return persistence.Product{
		ProductId: uuid.New(),
		Name:      "Test",
		Image:     "Test",
		Stock:     1,
		Price:     1.5,
	}
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
