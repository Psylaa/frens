package router

import (
	"github.com/microcosm-cc/bluemonday"

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
// @Accept  json,xml,x-www-form-urlencoded,multipart/form-data
// @Produce  json
// @Param username body string true "Username"
// @Param username formData string true "Username"
// @Param password body string true "Password"
// @Param password formData string true "Password"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /login [post]
func (lr *LoginRepo) login(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "login", "login")

	// Parse body
	var body struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	// Sanitize input to prevent XSS attacks
	p := bluemonday.UGCPolicy()
	body.Username = p.Sanitize(body.Username)
	body.Password = p.Sanitize(body.Password)

	// Validate body
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
// @Router /login/verify [get]
func (lr *LoginRepo) verifyToken(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "login", "verifyToken")

	// If we've gotten this far, the token has already passed through the middleware and is valid
	return c.SendStatus(fiber.StatusOK)
}
