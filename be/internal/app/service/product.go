package service

import (
	"context"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
	"gorm.io/gorm"
)

type ProductService interface {
	GetProductList(ctx context.Context) ([]response_model.Product, error)
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

func (p productService) GetProductList(ctx context.Context) ([]response_model.Product, error) {
	res := p.productRepository.GetList(ctx, p.db)
	newRes := make([]response_model.Product, 0)
	for _, value := range res {
		newRes = append(newRes, response_model.Product{value.ProductID.String(), value.Name, value.Image, value.Stock, value.Price})
	}
	return newRes, nil
}
