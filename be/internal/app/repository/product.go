package repository

import (
	"context"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetList(ctx context.Context, tx *gorm.DB) []persistence.Product
	GetProductByID(ctx context.Context, tx *gorm.DB, productID string) (persistence.Product, error)
	UpdateStockByProductID(ctx context.Context, tx *gorm.DB, productID string, stock int64) error
}

type productRepository struct {
}

func (p productRepository) GetList(ctx context.Context, tx *gorm.DB) []persistence.Product {
	product := make([]persistence.Product, 0)
	tx.Find(&product)

	return product
}

func (p productRepository) GetProductByID(ctx context.Context, tx *gorm.DB, productID string) (persistence.Product, error) {
	product := persistence.Product{}
	tx.Where("product_id=?", productID).First(&product)

	return product, tx.Error
}

func (p productRepository) UpdateStockByProductID(ctx context.Context, tx *gorm.DB, productID string, stock int64) error {
	product := persistence.Product{}
	tx.Model(&product).Where("product_id = ?", productID).Update("stock", stock)
	return tx.Error
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
