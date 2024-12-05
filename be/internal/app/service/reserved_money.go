package service

import (
	"context"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"gorm.io/gorm"
)

type ReserveMoneyService interface {
	GetReservedMoney(ctx context.Context) (map[string]int64, error)
	UpdateReservedMoney(ctx context.Context, req persistence.ReservedMoney) error
}

type reservedMoneyService struct {
	db                      *gorm.DB
	reservedMoneyRepository repository.ReservedMoneyRepository
}

func NewReservedMoneyService(db *gorm.DB, reservedMoneyRepository repository.ReservedMoneyRepository) ReserveMoneyService {
	return &reservedMoneyService{
		db:                      db,
		reservedMoneyRepository: reservedMoneyRepository,
	}
}

func (r reservedMoneyService) GetReservedMoney(ctx context.Context) (map[string]int64, error) {
	res, err := r.reservedMoneyRepository.GetReservedMoney(ctx, r.db)

	newRes := map[string]int64{
		"coins1":   res[1],
		"coins5":   res[5],
		"coins10":  res[10],
		"bank20":   res[20],
		"bank50":   res[50],
		"bank100":  res[100],
		"bank500":  res[500],
		"bank1000": res[1000],
	}

	return newRes, err
}

func (r reservedMoneyService) UpdateReservedMoney(ctx context.Context, req persistence.ReservedMoney) error {
	err := r.reservedMoneyRepository.UpdateReservedMoney(ctx, r.db, req)

	return err
}
