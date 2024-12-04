package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckoutRepository interface {
	CheckoutProduct(ctx context.Context, tx *gorm.DB, productID uuid.UUID)
}

type checkoutRepository struct {
}

func (c checkoutRepository) CheckoutProduct(ctx context.Context, tx *gorm.DB, productID uuid.UUID) {
	//TODO implement me
	panic("implement me")
}

func NewCheckoutRepository() CheckoutRepository {
	return &checkoutRepository{}
}
