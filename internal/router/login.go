package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
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
	rtr.Post("/", lr.login)
	rtr.Get("/verify", lr.verifyToken)
}

// @Summary Login
// @Description Authenticate a user and return a JWT token
// @Tags Login
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /login [post]
func (lr *LoginRepo) login(c *fiber.Ctx) error {
	logger.DebugLogRequestRecieved("router", "login", "login")
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

// @Summary Verify Token
// @Description Verify a JWT token
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401
// @Security ApiKeyAuth
// @Router /verify [get]
func (lr *LoginRepo) verifyToken(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
