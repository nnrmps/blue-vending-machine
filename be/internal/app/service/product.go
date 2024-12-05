package service

import (
	"context"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/pkg/request_model"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
	"gorm.io/gorm"
)

type ProductService interface {
	GetProductList(ctx context.Context) ([]response_model.Product, error)
	GetProductByID(ctx context.Context, productID string) (response_model.Product, error)
	UpdateProductByID(ctx context.Context, productID string, req request_model.UpdateProductByID) (response_model.Product, error)
	CreateProduct(ctx context.Context, req request_model.CreateProduct) (response_model.Product, error)
	DeleteProductByID(ctx context.Context, productID string) error
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

func (p productService) GetProductByID(ctx context.Context, productID string) (response_model.Product, error) {
	res, err := p.productRepository.GetProductByID(ctx, p.db, productID)
	newRes := response_model.Product{res.ProductID.String(), res.Name, res.Image, res.Stock, res.Price}
	return newRes, err
}

func (p productService) GetProductList(ctx context.Context) ([]response_model.Product, error) {
	res := p.productRepository.GetList(ctx, p.db)
	newRes := make([]response_model.Product, 0)
	for _, value := range res {
		newRes = append(newRes, response_model.Product{value.ProductID.String(), value.Name, value.Image, value.Stock, value.Price})
	}
	return newRes, nil
}

func (p productService) UpdateProductByID(ctx context.Context, productID string, req request_model.UpdateProductByID) (response_model.Product, error) {
	res, err := p.productRepository.UpdateProductByID(ctx, p.db, productID, req)
	newRes := response_model.Product{res.ProductID.String(), res.Name, res.Image, res.Stock, res.Price}
	return newRes, err
}

func (p productService) CreateProduct(ctx context.Context, req request_model.CreateProduct) (response_model.Product, error) {
	res, err := p.productRepository.CreateProduct(ctx, p.db, req)
	newRes := response_model.Product{res.ProductID.String(), res.Name, res.Image, res.Stock, res.Price}
	return newRes, err
}

func (p productService) DeleteProductByID(ctx context.Context, productID string) error {
	err := p.productRepository.DeleteProductByID(ctx, p.db, productID)
	return err
}
