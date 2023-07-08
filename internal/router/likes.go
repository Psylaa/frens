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
}

// @Summary Search Likes
// @Description Retrieve likes. If a like ID is provided, it is always used. Otherwise, a postID will return all likes for that post. If a userID is also provided, it will return either the like for that user/post or an empty array. If only a userID is provided, it will return all likes by that user for any post.
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param likeID query string false "Like ID"
// @Param postID query string false "Post ID"
// @Param userID query string false "User ID"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes [get]
func (lr *LikesRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "likes", "getLikes")

	// Get the query parameters
	queries := c.Queries()
	queryLikeID := queries["likeID"]
	queryPostID := queries["postID"]
	queryUserID := queries["userID"]

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
