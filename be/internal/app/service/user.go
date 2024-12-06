package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/setting"
	"gorm.io/gorm"
	"time"
)

type UserService interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type userService struct {
	db             *gorm.DB
	userRepository repository.UserRepository
}

func NewUserService(db *gorm.DB, userRepository repository.UserRepository) UserService {
	return &userService{
		db:             db,
		userRepository: userRepository,
	}
}

func (u userService) Login(ctx context.Context, username string, password string) (string, error) {
	res, err := u.userRepository.Login(ctx, u.db, username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": res.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(setting.AppConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
