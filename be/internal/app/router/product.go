package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/pkg/request_model"
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
	router.Get("/products/:productID", p.getProductByID)
}

func (p ProductController) InitAdminRouter(router fiber.Router) {
	router.Get("/products", p.getProductList)
	router.Get("/products/:productID", p.getProductByID)
	router.Post("/products", p.createProduct)
	router.Put("/products/:productID", p.updateProductByID)
	router.Delete("/products/:productID", p.deleteProductByID)
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

func (p ProductController) getProductByID(ctx *fiber.Ctx) error {
	res, err := p.productService.GetProductByID(ctx.Context(), ctx.Params("productID"))
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(response_model.BaseResponse[response_model.Product]{
		Data: &res,
	})
	return nil
}

func (p ProductController) updateProductByID(ctx *fiber.Ctx) error {
	req := request_model.UpdateProductByID{}

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	res, err := p.productService.UpdateProductByID(ctx.Context(), ctx.Params("productID"), req)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(response_model.BaseResponse[response_model.Product]{
		Data: &res,
	})
	return nil
}

func (p ProductController) createProduct(ctx *fiber.Ctx) error {
	req := request_model.CreateProduct{}

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	res, err := p.productService.CreateProduct(ctx.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(response_model.BaseResponse[response_model.Product]{
		Data: &res,
	})
	return nil
}

func (p ProductController) deleteProductByID(ctx *fiber.Ctx) error {
	err := p.productService.DeleteProductByID(ctx.Context(), ctx.Params("productID"))
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON("Delete Product success")
	return nil
}
