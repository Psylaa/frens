package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

// LoginRepo struct represents the /login route.
type LoginRepo struct {
	DB  *database.Database
	Srv *service.Service
}

// NewLoginRepo creates a new LoginRepo instance.
func NewLoginRepo(db *database.Database, srv *service.Service) *LoginRepo {
	return &LoginRepo{
		DB:  db,
		Srv: srv,
	}
}

// ConfigureRoutes configures the routes associated with login functionality.
func (lr *LoginRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Post("/login", lr.login)
	rtr.Get("/verify", lr.verifyToken)
}

// @route POST /v1/login/login
// @description Authenticate a user with their username and password.
// @tags login
// @produce json
// @success 200
// @failure 400
func (lr *LoginRepo) login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	if body.Username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	return lr.Srv.Login.Login(c, &body.Username, &body.Password)
}

// @route GET /v1/login/verify
// @description Verify the JWT token.
// @tags login
// @produce json
// @success 200
// @failure 400
func (lr *LoginRepo) verifyToken(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
