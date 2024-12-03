package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/repository"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/router"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	groupApi := app.Group("/api")
	//db := initGorm()
	dsn := "host=localhost user=root password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	_ = app.Listen(":8080")
}
