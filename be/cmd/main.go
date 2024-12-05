package main

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/router"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {
	setting.InitConfiguration()
}

func main() {
	app := fiber.New()

	groupApi := app.Group("/api")
	adminGroupApi := app.Group("/admin-api")
	adminGroupApi.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(setting.AppConfig.SecretKey)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Customize the error handling
			log.Println("JWT Error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
		},
	},
	))

	//init DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		setting.AppConfig.Database.Host,
		setting.AppConfig.Database.Username,
		setting.AppConfig.Database.Password,
		setting.AppConfig.Database.Name,
		setting.AppConfig.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(persistence.Product{})
	db.AutoMigrate(persistence.ReservedMoney{})
	db.AutoMigrate(persistence.User{})
	if err != nil {
		panic("failed to connect database")
	}

	//health
	healthService := service.NewHealthService()
	healthController := router.NewHealthController(healthService)
	healthController.InitRouters(groupApi)

	//product
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(db, productRepository)
	productController := router.NewProductController(productService)
	productController.InitRouter(groupApi)
	productController.InitAdminRouter(adminGroupApi)

	//checkout
	checkoutRepository := repository.NewCheckoutRepository()
	checkoutService := service.NewCheckoutService(db, checkoutRepository, productRepository)
	checkoutController := router.NewCheckoutController(checkoutService)
	checkoutController.InitRouter(groupApi)

	//reserve-money
	reservedRepository := repository.NewCheckoutRepository()
	reservedService := service.NewReservedMoneyService(db, reservedRepository)
	reservedController := router.NewReservedMoneyController(reservedService)
	reservedController.InitRouter(adminGroupApi)

	//user
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository)
	userController := router.NewUserController(userService)
	userController.InitRouter(groupApi)

	_ = app.Listen(":8080")
}
