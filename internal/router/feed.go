package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// getChronologicalFeed returns a list of posts from the users that the
// authenticated user is following, in chronological order.
func getChronologicalFeed(c *fiber.Ctx) error {
	// Get the user ID from the JWT.
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrUnauthorized))
	}
	logger.Log.Debug().Str("userID", userID.String()).Msg("Got user ID from JWT")

	// Get the cursor from the request query parameters. If it's not present,
	// we default to the current time, which will retrieve the most recent posts.
	cursorParam := c.Query("cursor")
	cursor := time.Now()
	if cursorParam != "" {
		unixTime, err := strconv.ParseInt(cursorParam, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidCursor))
		}
		cursor = time.Unix(unixTime, 0)
	}
	logger.Log.Debug().Time("cursor", cursor).Msg("Got cursor from query parameters")

	// Get the list of users that the authenticated user is following.
	following, err := db.Follows.GetFollowing(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.Log.Debug().Int("following", len(following)).Msg("Got following users")

	// Extract the following IDs
	followingIDs := make([]uuid.UUID, len(following))
	for i, follower := range following {
		followingIDs[i] = follower.TargetID
	}
	logger.Log.Debug().Interface("followingIDs", followingIDs).Msg("Got following user IDs")

	followingIDs = append(followingIDs, userID)

	// Get posts by following users and the current user
	posts, err := db.Posts.GetPostsByUserIDs(followingIDs, cursor, 10)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting posts")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.Log.Debug().Int("posts", len(posts)).Msg("Got posts")

	return c.JSON(response.GeneratePostsResponse(posts))
}

// getExploreFeed returns a list of the latest posts from all users
func getExploreFeed(c *fiber.Ctx) error {
	posts, err := db.Posts.GetLatestPublicPosts(25)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting posts")
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	logger.Log.Debug().Int("posts", len(posts)).Msg("Got posts")

	return c.JSON(response.GeneratePostsResponse(posts))
}
