package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u UserController) InitRouter(router fiber.Router) {
	router.Post("/login", u.login)
}

func (u UserController) login(ctx *fiber.Ctx) error {
	req := persistence.User{}

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	token, err := u.userService.Login(ctx.Context(), req.Username, req.Password)
	if err != nil {
		return err
	}
	_ = ctx.Status(200)
	_ = ctx.JSON(token)
	return nil
}
