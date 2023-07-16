package router

import (
	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

// AuthRepo struct represents the /Auth route.
type AuthRepo struct {
	Service *service.AuthService
}

func (ar *AuthRepo) addPublicRoutes(rtr fiber.Router) {
	grp := rtr.Group("/auth")
	grp.Post("/login", ar.Login)
	grp.Post("/register", ar.Register)
}

func (ar *AuthRepo) addPrivateRoutes(rtr fiber.Router) {
	grp := rtr.Group("/auth")
	grp.Get("/verify", ar.Verify)
	grp.Post("/refresh", ar.Service.Refresh)
	grp.Delete("/logout", ar.Logout)
}

func (a *AuthRepo) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return a.Service.Login(c, req.Email, req.Password)
}

func (a *AuthRepo) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return a.Service.Register(c, req.Username, req.Email, req.Password)
}

func (a *AuthRepo) Verify(c *fiber.Ctx) error {
	// If we've gotten this far, this means the request
	// has already passed through auth. Send a 200 with
	// an empty body.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (a *AuthRepo) Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
