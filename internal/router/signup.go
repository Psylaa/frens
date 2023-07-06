package router

import (
	"regexp"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

// SignupRepo struct represents the /signup route.
type SignupRepo struct {
	DB  *database.Database
	Srv *service.Service
}

// NewSignupRepo creates a new SignupRepo instance.
func NewSignupRepo(db *database.Database, srv *service.Service) *SignupRepo {
	return &SignupRepo{
		DB:  db,
		Srv: srv,
	}
}

// ConfigureRoutes configures the routes associated with signup functionality.
func (sr *SignupRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Post("/", sr.signup)
}

// CreateUserRequest represents the request body for creating a new user account.
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// @Summary Signup
// @Description Create a new user account.
// @Tags Signup
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "The user account to create"
// @Param user formData CreateUserRequest true "The user account to create"
// @Failure 400
// @Failure 500
// @Router /signup [post]
func (sr *SignupRepo) signup(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "signup", "signup")

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
