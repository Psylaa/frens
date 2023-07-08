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
	rtr.Get("/self", ur.getSelf)
	rtr.Delete("/self", ur.delete)
	rtr.Get("/:userID", ur.get)
	//rtr.Get("/search", ur.search) To be implemented. This is here for now to remind me not to change the regular "get" route to have search functionality
	rtr.Patch("/:userID", ur.update)
	rtr.Post("/:userID/blocks", ur.block)
	rtr.Delete("/:userID/blocks", ur.unblock)
}

// @Summary Get information about the authenticated user
// @Description Fetch information about the user making the request
// @Tags Users
// @Accept  json
// @Produce  json
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self [get]
func (ur *UsersRepo) getSelf(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "getSelf")

	// Get the userID from the token. This could vary depending on your authentication method.
	// For example, if you are using JWT for authentication, you could retrieve the userID from the payload.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// If the user ID is not provided or invalid, return an error
	if requestorID == nil {
		logger.Log.Info().Msg("No valid user ID provided in the token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}
	logger.DebugLogRequestUpdate("router", "users", "getSelf", "parsed userID from token: "+requestorID.String())

	// Send the request to the service layer
	return ur.Srv.Users.Get(c, requestorID)
}

// @Summary Get a user by ID
// @Description Fetch a specific user by their ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID} [get]
func (ur *UsersRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "get")

	// Parse userID from path
	userID := c.Params("userID")
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
	if userUUID == uuid.Nil {
		logger.Log.Error().Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	// Send the request to the service layer
	return ur.Srv.Users.Get(c, &userUUID)
}

// @Summary Update a user
// @Description Update a users information including bio, avatar, and cover. Note that avatar and cover must first be uploaded to the server and UUIDs must be provided.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param bio body string false "Bio"
// @Param bio formData string false "Bio"
// @Param avatarID body string false "Avatar ID"
// @Param avatarID formData string false "Avatar ID"
// @Param coverID body string false "Cover ID"
// @Param coverID formData string false "Cover ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID} [patch]
func (ur *UsersRepo) update(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "update")

	// Parse the request body
	var body struct {
		Bio      *string `form:"bio" json:"bio"`
		AvatarID *string `form:"avatarID" json:"avatarID"`
		CoverID  *string `form:"coverID" json:"coverID"`
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
// @Router /users/self [delete]
func (ur *UsersRepo) delete(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "delete")

	return ur.Srv.Users.Delete(c)
}

// @Summary Block a user
// @Description Block a user by their ID.
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/blocks [post]
func (ur *UsersRepo) block(c *fiber.Ctx) error {
	return nil
}

// @Summary Unblock a user
// @Description Unblock a user by their ID.
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/blocks [delete]
func (ur *UsersRepo) unblock(c *fiber.Ctx) error {
	return nil
}

// @Summary Get a block by user ID
// @Description Get a block by user ID
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/blocks [get]
func (ur *UsersRepo) getBlock(c *fiber.Ctx) error {
	return nil
}

// @Summary Get a list of all users that are following a user by user ID
// @Description Get a list of all users that are following a user by user ID
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [get]
func (ur *UsersRepo) GetFollowersByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Follow a user by user ID
// @Description Follow a user by user ID
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [post]
func (ur *UsersRepo) FollowUserByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Unfollow a user by user ID
// @Description Unfollow a user by user ID
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [delete]
func (ur *UsersRepo) UnfollowUserByUserID(c *fiber.Ctx) error {
	return nil
}
