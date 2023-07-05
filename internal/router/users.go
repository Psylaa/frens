package router

import (
	"regexp"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

type UsersRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewUsersRepo(db *database.Database, srv *service.Service) *UsersRepo {
	return &UsersRepo{
		DB:  db,
		Srv: srv,
	}
}

func (ur *UsersRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/:userId", ur.get)
	//rtr.Get("/search", ur.search) To be implemented. This is here for now to remind me not to change the regular "get" route to have search functionality
	rtr.Post("/", ur.create)
	rtr.Patch("/:userId", ur.update)
	rtr.Delete("/", ur.delete)
}

// @Summary Get a user by ID
// @Description Fetch a specific user by their ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userId} [get]
func (ur *UsersRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "get")

	// Parse userID from path
	userID := c.Params("userId")
	if userID == "" {
		logger.Log.Info().Msg("No user ID provided")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}
	logger.DebugLogRequestUpdate("router", "users", "get", "parsed userID from path: "+userID)

	// Sanitize the input
	p := bluemonday.UGCPolicy()
	userID = p.Sanitize(userID)

	// Validate the user ID format
	if matched, _ := regexp.MatchString(`\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`, userID); !matched {
		logger.Log.Error().Msg("Invalid user ID format")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}
	logger.DebugLogRequestUpdate("router", "users", "get", "validated user ID format with regex")

	// Convert userID to UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing user ID to UUID")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}
	logger.DebugLogRequestUpdate("router", "users", "get", "converted userID to UUID")

	// Send the request to the service layer
	return ur.Srv.Users.Get(c, &userUUID)
}

// @Summary Create a user
// @Description Create a new user.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param username formData string true "Username"
// @Param email body string true "Email"
// @Param email formData string true "Email"
// @Param password body string true "Password"
// @Param password formData string true "Password"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/ [post]
func (ur *UsersRepo) create(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "create")

	// Parse the request body
	var body struct {
		Username string `form:"username" json:"username"`
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUserID))
	}

	// Sanitize the input
	p := bluemonday.UGCPolicy()
	body.Username = p.Sanitize(body.Username)
	body.Email = p.Sanitize(body.Email)
	// Don't sanitize password - it might unintentionally change it.

	// Validate email format
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, body.Email); !matched {
		logger.Log.Error().Msg("Invalid email format")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidEmail))
	}

	// Validate username and password - they should not be empty
	if body.Username == "" || body.Password == "" {
		logger.Log.Error().Msg("Username or password is empty")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	return ur.Srv.Users.Create(c, body.Username, body.Email, body.Password)
}

// @Summary Update a user
// @Description Update a users information including bio, avatar, and cover. Note that avatar and cover must first be uploaded to the server and UUIDs must be provided.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param bio body string false "Bio"
// @Param bio formData string false "Bio"
// @Param avatarId body string false "Avatar ID"
// @Param avatarId formData string false "Avatar ID"
// @Param coverId body string false "Cover ID"
// @Param coverId formData string false "Cover ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userId} [patch]
func (ur *UsersRepo) update(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "update")

	// Parse the request body
	var body struct {
		Bio      *string `form:"bio" json:"bio"`
		AvatarID *string `form:"avatarId" json:"avatarId"`
		CoverID  *string `form:"coverId" json:"coverId"`
	}

	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}
	logger.DebugLogRequestUpdate("router", "users", "update", "parsed request body")

	// Sanitize the input
	p := bluemonday.UGCPolicy()
	if body.Bio != nil {
		*body.Bio = p.Sanitize(*body.Bio)
		if len(*body.Bio) > 256 {
			logger.Log.Error().Msg("Bio is too long")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
		}
	}

	if body.AvatarID != nil {
		*body.AvatarID = p.Sanitize(*body.AvatarID)
	}

	if body.CoverID != nil {
		*body.CoverID = p.Sanitize(*body.CoverID)
	}
	logger.DebugLogRequestUpdate("router", "users", "update", "sanitized request body")

	// Convert avatarID and coverID to UUID
	var err error

	// Avatar ID
	var avatarUUID, coverUUID *uuid.UUID
	if body.AvatarID != nil {
		avatarUUID = new(uuid.UUID)
		*avatarUUID, err = uuid.Parse(*body.AvatarID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing avatar ID to UUID")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidAvatarUUID))
		}
	}
	logger.DebugLogRequestUpdate("router", "users", "update", "parsed avatar ID to UUID")

	// Cover ID
	if body.CoverID != nil {
		coverUUID = new(uuid.UUID)
		*coverUUID, err = uuid.Parse(*body.CoverID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing cover ID to UUID")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidCoverUUID))
		}
	}
	logger.DebugLogRequestUpdate("router", "users", "update", "parsed cover ID to UUID")

	return ur.Srv.Users.Update(c, body.Bio, avatarUUID, coverUUID)
}

// @Summary Delete self
// @Description Delete the user associated with the provided access token.
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users [delete]
func (ur *UsersRepo) delete(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "delete")

	return ur.Srv.Users.Delete(c)
}
