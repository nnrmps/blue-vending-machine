package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
)

type HealthController struct {
	healthService service.HealthService
}

func NewHealthController(healthService service.HealthService) *HealthController {
	return &HealthController{healthService: healthService}
}

func (h HealthController) InitRouters(router fiber.Router) {
	router.Get("/health", h.healthHandlers)
}

func (h HealthController) healthHandlers(ctx *fiber.Ctx) error {
	err := h.healthService.HealthCheck()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	_ = ctx.JSON(fiber.Map{"success": true})
	return nil
}
