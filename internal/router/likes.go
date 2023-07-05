package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/service"
)

type LikesRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewLikesRepo(db *database.Database, srv *service.Service) *LikesRepo {
	return &LikesRepo{
		DB:  db,
		Srv: srv,
	}
}

func (lr *LikesRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", lr.get)
	rtr.Post("/", lr.create)
	rtr.Delete("/", lr.deleteByID)
}

// @Summary Get likes
// @Description Retrieve likes. If a like ID is provided, it is always used. Otherwise, a postID will return all likes for that post. If a userID is also provided, it will return either the like for that user/post or an empty array. If only a userID is provided, it will return all likes by that user for any post.
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param likeId query string false "Like ID"
// @Param postId query string false "Post ID"
// @Param userId query string false "User ID"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes [get]
func (lr *LikesRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "likes", "getLikes")

	// Get the query parameters
	queries := c.Queries()
	queryLikeID := queries["likeId"]
	queryPostID := queries["postId"]
	queryUserID := queries["userId"]

	// If no query parameters were provided, return an error
	if len(queryLikeID) == 0 && len(queryPostID) == 0 && len(queryUserID) == 0 {
		logger.Log.Error().Msg("No query parameters provided")
		return c.Status(fiber.StatusBadRequest).SendString("No query parameters provided")
	}

	// Convert all provided IDs to UUIDs
	var likeID, postID, userID uuid.UUID
	var err error
	if len(queryLikeID) > 0 {
		likeID, err = uuid.Parse(queryLikeID)
	}
	if len(queryPostID) > 0 {
		postID, err = uuid.Parse(queryPostID)
	}
	if len(queryUserID) > 0 {
		userID, err = uuid.Parse(queryUserID)
	}
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing query parameters")
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing query parameters")
	}

	// If a like ID was provided, use that
	if likeID != uuid.Nil {
		return lr.Srv.Likes.GetByID(c, &likeID)
	}

	// Otherwise, if a post ID was provided, use that
	if postID != uuid.Nil {
		// If a user ID was also provided, use that
		if userID != uuid.Nil {
			return lr.Srv.Likes.GetByPostIDAndUserID(c, &postID, &userID)
		}
		// Otherwise, just use the post ID
		return lr.Srv.Likes.GetByPostID(c, &postID)
	}

	// Otherwise, if a user ID was provided, use that
	if userID != uuid.Nil {
		return nil // to be implemented
		//return lr.Srv.Likes.GetByUserID(c, &userID)
	}

	// If we got here, something went wrong
	logger.Log.Error().Msg("Error parsing query parameters")
	return c.Status(fiber.StatusBadRequest).SendString("Error parsing query parameters")
}

// @Summary Create a like
// @Description Create a new like for a user based on the provided token
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param postId body string true "Post ID"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes/ [post]
func (lr *LikesRepo) create(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "likes", "createLike")

	// Get the post ID from the URL
	postID := c.Params("postId")
	if postID == "" {
		logger.Log.Error().Msg("No post ID provided")
		return c.Status(fiber.StatusBadRequest).SendString("No post ID provided")
	}

	// Convert the post ID to a UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing post ID")
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing post ID")
	}

	// Send the request to the service layer
	return lr.Srv.Likes.Create(c, &postUUID)
}

// @Summary Delete a like
// @Description Delete a specific like
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param likeId path string true "Like ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes [delete]
func (lr *LikesRepo) deleteByID(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "likes", "deleteLike")
	return nil
}
