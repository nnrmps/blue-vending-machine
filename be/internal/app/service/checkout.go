package service

import (
	"context"
	"errors"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/pkg/request_model"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
	"gorm.io/gorm"
)

type CheckoutService interface {
	CheckoutProduct(ctx context.Context, productID string, total request_model.Money) (response_model.Checkout, error)
}

type checkoutService struct {
	db                 *gorm.DB
	checkoutRepository repository.ReservedMoneyRepository
	productRepository  repository.ProductRepository
}

func NewCheckoutService(db *gorm.DB, checkoutRepository repository.ReservedMoneyRepository, productRepository repository.ProductRepository) CheckoutService {
	return &checkoutService{
		db:                 db,
		checkoutRepository: checkoutRepository,
		productRepository:  productRepository,
	}
}

func updateReservedMoney(reservedMoney map[int64]int64, total request_model.Money) map[int64]int64 {
	reservedMoney[1] = reservedMoney[1] + total.Coins1
	reservedMoney[5] = reservedMoney[5] + total.Coins5
	reservedMoney[10] = reservedMoney[10] + total.Coins10
	reservedMoney[20] = reservedMoney[20] + total.Bank20
	reservedMoney[50] = reservedMoney[50] + total.Bank50
	reservedMoney[100] = reservedMoney[100] + total.Bank100
	reservedMoney[500] = reservedMoney[500] + total.Bank500
	reservedMoney[1000] = reservedMoney[1000] + total.Bank1000

	return reservedMoney
}

func (c checkoutService) CheckoutProduct(ctx context.Context, productID string, total request_model.Money) (response_model.Checkout, error) {
	productDetail, err := c.productRepository.GetProductByID(ctx, c.db, productID)
	if err != nil {
		return response_model.Checkout{}, err
	}

	if productDetail.Stock <= 0 {
		return response_model.Checkout{}, errors.New("Sorry, product out of stock")
	}

	totalDeposit := (total.Coins1 * 1) + (total.Coins5 * 5) + (total.Coins10 * 10) + (total.Bank20 * 20) + (total.Bank50 * 50) + (total.Bank100 * 100) + (total.Bank500 * 500) + (total.Bank1000 * 1000)
	if totalDeposit < productDetail.Price {
		return response_model.Checkout{}, errors.New("Sorry, not enough deposit")
	}

	reservedMoney, err := c.checkoutRepository.GetReservedMoney(ctx, c.db)
	if err != nil {
		return response_model.Checkout{}, err
	}
	reservedMoney = updateReservedMoney(reservedMoney, total)

	totalChange := totalDeposit - productDetail.Price

	moneyList := []int64{1000, 500, 100, 50, 20, 10, 5, 1}
	finalReservedMoney := map[int64]int64{
		1000: 0,
		500:  0,
		100:  0,
		50:   0,
		20:   0,
		10:   0,
		5:    0,
		1:    0,
	}
	i := 0

	for totalChange > 0 {
		if i >= len(moneyList) {
			return response_model.Checkout{}, errors.New("Sorry, don't have enough change.")
		}
		if reservedMoney[moneyList[i]] > 0 && totalChange-moneyList[i] >= 0 {
			reservedMoney[moneyList[i]] -= 1
			finalReservedMoney[moneyList[i]] += 1
			totalChange = totalChange - moneyList[i]
			continue
		}
		i++
	}

	newReservedMoney := persistence.ReservedMoney{Coins1: reservedMoney[1],
		Coins5:   reservedMoney[5],
		Coins10:  reservedMoney[10],
		Bank20:   reservedMoney[20],
		Bank50:   reservedMoney[50],
		Bank100:  reservedMoney[100],
		Bank500:  reservedMoney[500],
		Bank1000: reservedMoney[1000]}
	tx := c.db.Begin()

	err = c.checkoutRepository.UpdateReservedMoney(ctx, tx, newReservedMoney)
	if err != nil {
		tx.Rollback()
		return response_model.Checkout{}, err
	}

	err = c.productRepository.UpdateStockByProductID(ctx, tx, productID, productDetail.Stock-1)
	if err != nil {
		tx.Rollback()
		return response_model.Checkout{}, err
	}

	tx.Commit()
	newRes := response_model.Checkout{
		TotalChange: totalDeposit - productDetail.Price,
		Coins1:      finalReservedMoney[1],
		Coins5:      finalReservedMoney[5],
		Coins10:     finalReservedMoney[10],
		Bank20:      finalReservedMoney[20],
		Bank50:      finalReservedMoney[50],
		Bank100:     finalReservedMoney[100],
		Bank500:     finalReservedMoney[500],
		Bank1000:    finalReservedMoney[1000],
	}

	return newRes, nil
}
