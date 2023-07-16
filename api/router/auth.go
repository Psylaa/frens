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
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	validate.Struct(req)
	return a.Service.Login(c, req.Email, req.Password)
}

func (a *AuthRepo) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	validate.Struct(req)
	return a.Service.Register(c, req.Username, req.Email, req.Password)
}
