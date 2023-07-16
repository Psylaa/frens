package router

import (
	"github.com/bwoff11/frens/service/user"
	"github.com/gofiber/fiber/v2"
)

// AuthRepo struct represents the /Auth route.
type AuthRepo struct {
	Service *user.Service
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
	a.Service.Register(req.Username, req.Email, req.Password)
}
