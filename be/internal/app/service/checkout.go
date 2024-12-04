package service

import (
	"context"
	"errors"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
	"gorm.io/gorm"
)

type CheckoutService interface {
	CheckoutProduct(ctx context.Context, productID string) (response_model.Checkout, error)
}

type checkoutService struct {
	db                 *gorm.DB
	checkoutRepository repository.CheckoutRepository
}

func NewCheckoutService(db *gorm.DB, checkoutRepository repository.CheckoutRepository) CheckoutService {
	return &checkoutService{
		db:                 db,
		checkoutRepository: checkoutRepository,
	}
}

func (c checkoutService) CheckoutProduct(ctx context.Context, productID string) (response_model.Checkout, error) {
	//res := c.productRepository.Checkout(ctx, c.db, productID)
	if productID == "3234" {
		return response_model.Checkout{}, errors.New("error jaaa")
	}
	newRes := response_model.Checkout{
		TotalChange: 1234,
	}

	return newRes, nil
}
