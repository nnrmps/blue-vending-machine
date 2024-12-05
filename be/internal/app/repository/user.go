package repository

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(ctx context.Context, tx *gorm.DB, userName, password string) (persistence.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u userRepository) Login(ctx context.Context, tx *gorm.DB, username string, password string) (persistence.User, error) {
	user := persistence.User{}

	pw := sha256.Sum256([]byte(password))
	hashPassword := fmt.Sprintf("%x", pw[:])

	tx = tx.Where("username=? AND password = ?", username, hashPassword).First(&user)

	return user, tx.Error
}
