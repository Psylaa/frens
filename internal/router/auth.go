package router

import (
	"regexp"

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

// NewAuthRepo creates a new AuthRepo instance.
func NewAuthRepo(db *database.Database, srv *service.Service) *AuthRepo {
	return &AuthRepo{
		DB:  db,
		Srv: srv,
	}
}

// ConfigureRoutes configures the routes associated with Auth functionality.
func (lr *AuthRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Post("/", lr.login)
	rtr.Get("/verify", lr.verify)
	rtr.Post("/register", lr.register)
}

// @Summary User Login
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
// @Router /login [post]
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

// @Summary Verify Token
// @Description Verify the validity of the provided JWT token.
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

// CreateUserRequest represents the request body for creating a new user account.
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// @Summary User Registration
// @Description Register a new user with the provided account details.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "The user account to create"
// @Param user formData CreateUserRequest true "The user account to create"
// @Failure 400
// @Failure 500
// @Router /auth/register [post]
func (sr *AuthRepo) register(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "auth", "register")

	// Parse the request body
	req := new(CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUserID))
	}

	// Sanitize the input
	p := bluemonday.UGCPolicy()
	req.Username = p.Sanitize(req.Username)
	req.Email = p.Sanitize(req.Email)
	req.Phone = p.Sanitize(req.Phone)
	// Don't sanitize password - it might unintentionally change it.

	// Validate email or phone format
	if req.Email == "" && req.Phone == "" {
		logger.Log.Error().Msg("Email or phone must be provided")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}
	if req.Email != "" {
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, req.Email); !matched {
			logger.Log.Error().Msg("Invalid email format")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidEmail))
		}
	}
	if req.Phone != "" {
		if matched, _ := regexp.MatchString(`^[0-9]{10}$`, req.Phone); !matched {
			logger.Log.Error().Msg("Invalid phone format")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidPhone))
		}
	}

	// Validate username and password - they should not be empty
	if req.Username == "" || req.Password == "" {
		logger.Log.Error().Msg("Username or password is empty")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	return sr.Srv.Users.Create(c, req.Username, req.Email, req.Phone, req.Password)
}
