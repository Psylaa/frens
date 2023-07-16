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
	grp.Post("/refresh", ar.Service.Refresh)
	//rtr.Post("/logout", ar.Logout)
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
