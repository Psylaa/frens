package router

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	rtr.Get("/", ur.search)
	rtr.Get("/self", ur.search)
	rtr.Get("/:userID", ur.getByID)
	rtr.Get("/:userID/posts", ur.getPosts)
}

// @Summary Search Users
// @Description Search for users with query parameters.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userID query string false "User ID"
// @Param username query string false "Username"
// @Param count query string false "The number of users to return."
// @Param offset query string false "The number of users to offset the returned users by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users [get]
func (ur *UsersRepo) search(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "search")
	return nil
}

// @Summary Retrieve User by ID
// @Description Retrieves a user by ID.
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
func (ur *UsersRepo) getByID(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "getByID")

	// Get target user by ID from params
	userID := c.Params("userID")

	// Parse as UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
	}

	// Send to service layer
	return ur.Srv.Users.GetByID(c, &userUUID)
}

// @Summary Update User
// @Description Update the authenticated user's profile.
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self [put]
func (ur *UsersRepo) update(c *fiber.Ctx) error {
	return nil
}

// @Summary Delete User
// @Description Delete the authenticated user's profile.
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
	return nil
}

// @Summary Confirm Delete User
// @Description Confirm the deletion of the authenticated user's profile.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param confirmationCode query string true "Confirmation Code"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self/confirm [delete]
func (ur *UsersRepo) confirmDelete(c *fiber.Ctx) error {
	return nil
}

type GetPostsByUserIDRequest struct {
	UserID uuid.UUID `json:"userID" validate:"required"`
	Count  int       `json:"count" validate:"omitempty,min=1,max=100"`
	Cursor string    `json:"cursor" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

// @Summary Get Posts by User ID
// @Description Get posts by user ID
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param count query int false "The number of posts to return."
// @Param cursor query string false "Cursor to start the page from."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/posts [get]
func (ur *UsersRepo) getPosts(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "users", "getPosts")

	// Parsing userID from path
	userID := c.Params("userID")
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid User ID")
	}

	// Parsing other parameters
	var request GetPostsByUserIDRequest
	err = c.QueryParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	request.UserID = parsedUUID

	// Create validator instance and validate the request
	v := validator.New()
	errList, err := validateRequest(v, request)
	if err != nil {
		// Send back the validation errors
		return c.Status(fiber.StatusBadRequest).JSON(errList)
	}

	// Set default count if not provided
	if request.Count == 0 {
		request.Count = 25
	}

	// Set default cursor if not provided
	var cursor time.Time
	if request.Cursor == "" {
		cursor = time.Now()
	} else {
		cursor, err = time.Parse(time.RFC3339, request.Cursor)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid cursor")
		}
	}

	// Pass the parameters to the GetByUserID function
	return ur.Srv.Posts.GetByUserID(c, &request.UserID, cursor, request.Count)
}

// @Summary Block User
// @Description Blocks the specified user from interacting with the authenticated user.
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

// @Summary Unblock User
// @Description Removes block on the specified user, allowing them to interact with the authenticated user.
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

// @Summary Retrieve Users Who are Following the Authenticated User
// @Description Retrieves a list of users following the authenticated user.
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param count query string false "The number of follows to return."
// @Param offset query string false "The number of follows to offset the returned follows by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self/followers [get]
func (fr *FollowsRepo) getSelfFollowers(c *fiber.Ctx) error {
	return nil
}

// @Summary Retrieve Users that the Authenticated User is Following
// @Description Retrieves a list of users the authenticated user is following.
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param count query string false "The number of follows to return."
// @Param offset query string false "The number of follows to offset the returned follows by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self/following [get]
func (fr *FollowsRepo) getSelfFollowing(c *fiber.Ctx) error {
	return nil
}

// @Summary Get Users Who are Following the Specified User
// @Description Get a list of all users that are following a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [get]
func (ur *UsersRepo) getFollowersByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Get Users that the Specified User is Following
// @Description Get a list of all users that a user is following by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/following [get]
func (ur *UsersRepo) getFollowingByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Follow a user by user ID
// @Description Follow a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [post]
func (ur *UsersRepo) followUserByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Unfollow a user by user ID
// @Description Unfollow a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [delete]
func (ur *UsersRepo) unfollowUserByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Get likes by user ID
// @Description Get likes by user ID
// @Tags Likes
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param count query string false "The number of likes to return."
// @Param offset query string false "The number of likes to offset the returned likes by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/likes [get]
func (ur *UsersRepo) getLikesByUserID(c *fiber.Ctx) error {
	return nil
}
