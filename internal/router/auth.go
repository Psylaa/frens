package router

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

// AuthRepo struct represents the /Auth route.
type AuthRepo struct {
	Service *service.Service
}

func (ar *AuthRepo) ConfigurePublicRoutes(rtr fiber.Router) {
	rtr.Post("/login", ar.login)
	rtr.Post("/register", ar.register)
}

func (ar *AuthRepo) ConfigureProtectedRoutes(rtr fiber.Router) {
	rtr.Get("/verify", ar.verify)
	rtr.Post("/logout", ar.logout)
}

// @Summary Authenticate User
// @Description Authenticate a user with the given credentials and return a JWT token.
// @Tags Auth
// @Accept  json,xml,x-www-form-urlencoded,multipart/form-data
// @Produce  json
// @Param email body string true "Email"
// @Param email formData string true "Email"
// @Param password body string true "Password"
// @Param password formData string true "Password"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /auth/login [post]
func (ar *AuthRepo) login(c *fiber.Ctx) error {
	logger.Debug(logger.LogMessage{
		Package:  "router",
		Function: "login",
		Message:  "Logging in user",
	})

	req := new(models.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	return ar.Service.Users.Login(c, req)
}

// @Summary Logout User
// @Description Logs out the user associated with the provided authentication token. The token will no longer be valid.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401
// @Security ApiKeyAuth
// @Router /auth/logout [post]
func (ar *AuthRepo) logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

// @Summary Verify Authentication Token
// @Description Verifies the authenticity of the provided authentication token.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401
// @Security ApiKeyAuth
// @Router /auth/verify [get]
func (ar *AuthRepo) verify(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

// @Summary Register New User
// @Description Creates a new user account and returns a confirmation.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "The user account to create"
// @Param user formData models.RegisterRequest true "The user account to create"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /auth/register [post]
func (ar *AuthRepo) register(c *fiber.Ctx) error {
	logger.Debug(logger.LogMessage{
		Package:  "router",
		Function: "register",
		Message:  "Registering new user",
	})

	req := new(models.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	return ar.Service.Users.Create(c, req)
}
