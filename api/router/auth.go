package router

import (
	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

// AuthRepo struct represents the /Auth route.
type AuthRepo struct {
	Service *service.AuthService
}

func (a *AuthRepo) Login(c *fiber.Ctx) error {
	return nil
}

func (a *AuthRepo) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	validate.Struct(req)
	return a.Service.Register(req.Username, req.Email, req.Password)
}
