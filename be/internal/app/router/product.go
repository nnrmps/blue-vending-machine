package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (p ProductController) InitRouter(router fiber.Router) {
	router.Get("/products", p.getProductList)
}

func (p ProductController) getProductList(ctx *fiber.Ctx) error {
	res, err := p.productService.GetProductList(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(response_model.BaseResponse[[]response_model.Product]{
		Data: &res,
	})
	return nil
}
