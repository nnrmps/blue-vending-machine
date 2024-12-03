package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
	"gorm.io/gorm"
)

type ProductService interface {
	GetProductByID(ctx context.Context, productID uuid.UUID) (response_model.Product, error)
}

type productService struct {
	db                *gorm.DB
	productRepository repository.ProductRepository
}

func NewProductService(db *gorm.DB, productRepository repository.ProductRepository) ProductService {
	return &productService{
		db:                db,
		productRepository: productRepository,
	}
}

func (p productService) GetProductByID(ctx context.Context, productID uuid.UUID) (response_model.Product, error) {
	res := p.productRepository.FindByID(ctx, p.db, productID)
	newRes := response_model.Product{
		ProductId: res.ProductId.String(),
		Name:      res.Name,
		ImageUrl:  res.Image,
		Stock:     res.Stock,
		Price:     res.Price,
	}
	return newRes, nil
}
