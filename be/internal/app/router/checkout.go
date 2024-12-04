package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/pkg/request_model"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
)

type CheckoutController struct {
	checkoutService service.CheckoutService
}

func NewCheckoutController(checkoutService service.CheckoutService) *CheckoutController {
	return &CheckoutController{checkoutService: checkoutService}
}

func (c CheckoutController) InitRouter(router fiber.Router) {
	router.Post("/checkout", c.checkoutByProductID)
}

func (c CheckoutController) checkoutByProductID(ctx *fiber.Ctx) error {
	body := new(request_model.Checkout)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	res, err := c.checkoutService.CheckoutProduct(ctx.Context(), body.ProductId, body.Total)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(response_model.BaseResponse[response_model.Checkout]{
		Data: &res,
	})
	return nil
}
