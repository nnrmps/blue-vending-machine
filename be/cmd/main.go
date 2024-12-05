package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/router"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	setting.InitConfiguration()
}

func main() {
	app := fiber.New()

	groupApi := app.Group("/api")
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

	//checkout
	checkoutRepository := repository.NewCheckoutRepository()
	checkoutService := service.NewCheckoutService(db, checkoutRepository, productRepository)
	checkoutController := router.NewCheckoutController(checkoutService)
	checkoutController.InitRouter(groupApi)

	//reserve-money
	reservedRepository := repository.NewCheckoutRepository()
	reservedService := service.NewReservedMoneyService(db, reservedRepository)
	reservedController := router.NewReservedMoneyController(reservedService)
	reservedController.InitRouter(groupApi)

	_ = app.Listen(":8080")
}
