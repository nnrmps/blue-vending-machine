package repository

import (
	"context"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"gorm.io/gorm"
)

type ReservedMoneyRepository interface {
	GetReservedMoney(ctx context.Context, tx *gorm.DB) (map[int64]int64, error)
	UpdateReservedMoney(ctx context.Context, tx *gorm.DB, reservedMoney persistence.ReservedMoney) error
}

type reservedMoneyRepository struct {
}

func (c reservedMoneyRepository) GetReservedMoney(ctx context.Context, tx *gorm.DB) (map[int64]int64, error) {
	money := persistence.ReservedMoney{}
	tx.First(&money)
	newMoney := map[int64]int64{
		1:    money.Coins1,
		5:    money.Coins5,
		10:   money.Coins10,
		20:   money.Bank20,
		50:   money.Bank50,
		100:  money.Bank100,
		500:  money.Bank500,
		1000: money.Bank1000,
	}

	return newMoney, tx.Error
}

func (c reservedMoneyRepository) UpdateReservedMoney(ctx context.Context, tx *gorm.DB, reservedMoney persistence.ReservedMoney) error {
	tx.Exec("UPDATE reserved_money SET coins1 = ?,coins5 = ?,coins10 = ?,bank20 = ?,bank50 = ?,bank100 = ?,bank500 = ?,bank1000 = ?", reservedMoney.Coins1, reservedMoney.Coins5, reservedMoney.Coins10, reservedMoney.Bank20, reservedMoney.Bank50, reservedMoney.Bank100, reservedMoney.Bank500, reservedMoney.Bank1000)
	return tx.Error
}

func NewCheckoutRepository() ReservedMoneyRepository {
	return &reservedMoneyRepository{}
}
