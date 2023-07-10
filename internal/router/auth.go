package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

// AuthRepo struct represents the /Auth route.
type AuthRepo struct {
	DB  *database.Database
	Srv *service.Service
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=1,max=24"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required"`
}

// NewAuthRepo creates a new AuthRepo instance.
func NewAuthRepo(db *database.Database, srv *service.Service) *AuthRepo {
	return &AuthRepo{
		DB:  db,
		Srv: srv,
	}
}

func (lr *AuthRepo) ConfigurePublicRoutes(rtr fiber.Router) {
	rtr.Post("/login", lr.login)
	rtr.Post("/register", lr.register)
}

func (lr *AuthRepo) ConfigureProtectedRoutes(rtr fiber.Router) {
	rtr.Get("/verify", lr.verify)
	rtr.Post("/logout", lr.logout)
}

// @Summary Authenticate User
// @Description Authenticate a user with the given credentials and return a JWT token.
// @Tags Auth
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
// @Router /auth/login [post]
func (lr *AuthRepo) login(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "Auth", "Auth")

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

	return lr.Srv.Auth.Login(c, body.Username, body.Password)
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
func (lr *AuthRepo) logout(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "Auth", "logout")

	return c.SendStatus(fiber.StatusOK)
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
func (lr *AuthRepo) verify(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "Auth", "verifyToken")

	// If we've gotten this far, the token has already passed through the middleware and is valid
	return c.SendStatus(fiber.StatusOK)
}

// @Summary Register New User
// @Description Creates a new user account and returns a confirmation.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "The user account to create"
// @Param user formData RegisterRequest true "The user account to create"
// @Failure 400
// @Failure 500
// @Router /auth/register [post]
func (sr *AuthRepo) register(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "auth", "register")

	// Parse the request body
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUserID))
	}

	// Sanitize the input
	p := bluemonday.UGCPolicy()
	req.Username = p.Sanitize(req.Username)
	req.Email = p.Sanitize(req.Email)
	// Don't sanitize password - it might unintentionally change it.

	// Validate using validator package
	v := validator.New()
	if jsonErrs, err := validateRequest(v, req); err != nil {
		logger.Log.Error().Err(err).Msg("Validation error")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"errors": jsonErrs})
	}

	return sr.Srv.Users.Create(c, req.Username, req.Email, req.Password)
}
