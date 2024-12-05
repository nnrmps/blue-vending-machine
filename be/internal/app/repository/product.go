package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/pkg/request_model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetList(ctx context.Context, tx *gorm.DB) []persistence.Product
	GetProductByID(ctx context.Context, tx *gorm.DB, productID string) (persistence.Product, error)
	UpdateStockByProductID(ctx context.Context, tx *gorm.DB, productID string, stock int64) error
	UpdateProductByID(ctx context.Context, tx *gorm.DB, productID string, req request_model.UpdateProductByID) (persistence.Product, error)
	CreateProduct(ctx context.Context, tx *gorm.DB, req request_model.CreateProduct) (persistence.Product, error)
	DeleteProductByID(ctx context.Context, tx *gorm.DB, productID string) error
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

func (p productRepository) UpdateProductByID(ctx context.Context, tx *gorm.DB, productID string, req request_model.UpdateProductByID) (persistence.Product, error) {
	product := persistence.Product{uuid.MustParse(req.ProductId), req.Name, req.ImageUrl, req.Stock, req.Price}
	tx.Save(&product)
	return product, tx.Error
}

func (p productRepository) CreateProduct(ctx context.Context, tx *gorm.DB, req request_model.CreateProduct) (persistence.Product, error) {
	product := persistence.Product{uuid.New(), req.Name, req.ImageUrl, req.Stock, req.Price}
	tx.Create(&product)
	return product, tx.Error
}

func (p productRepository) DeleteProductByID(ctx context.Context, tx *gorm.DB, productID string) error {
	product := persistence.Product{}
	tx.Where("product_id = ?", productID).Delete(&product)
	return tx.Error
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
