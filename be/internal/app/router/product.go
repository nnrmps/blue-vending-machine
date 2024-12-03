package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (p ProductController) InitRouter(router fiber.Router) {
	router.Get("/product/:productID", p.getProductByID)
}

func (p ProductController) getProductByID(c *fiber.Ctx) error {
	productID := uuid.New()
	res, err := p.productService.GetProductByID(c.Context(), productID)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = c.JSON(res)
	return nil
}
