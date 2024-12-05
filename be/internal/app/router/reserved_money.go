package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/pkg/response_model"
)

type ReservedMoneyController struct {
	reservedMoneyService service.ReserveMoneyService
}

func NewReservedMoneyController(reservedMoneyService service.ReserveMoneyService) *ReservedMoneyController {
	return &ReservedMoneyController{reservedMoneyService: reservedMoneyService}
}

func (r ReservedMoneyController) InitRouter(router fiber.Router) {
	router.Get("/reserved-money", r.getReservedMoney)
	router.Put("/reserved-money", r.updateReservedMoney)
}

func (r ReservedMoneyController) getReservedMoney(ctx *fiber.Ctx) error {
	res, err := r.reservedMoneyService.GetReservedMoney(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(response_model.BaseResponse[map[string]int64]{
		Data: &res,
	})
	return nil
}

func (r ReservedMoneyController) updateReservedMoney(ctx *fiber.Ctx) error {
	req := persistence.ReservedMoney{}

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := r.reservedMoneyService.UpdateReservedMoney(ctx.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.Status(200)
	_ = ctx.JSON("Update success!")
	return nil
}
