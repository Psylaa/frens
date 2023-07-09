package service

import (
	"net/http"
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FeedRepo struct{}

func (f *FeedRepo) GetChrono(c *fiber.Ctx, userID *uuid.UUID, cursor time.Time) error {
	logger.DebugLogRequestReceived("service", "feed", "GetChrono")

	// Get requestor id from context
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the users that the requestor user is following
	following, err := db.Follows.GetFollowing(requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting users followed")
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	logger.Log.Debug().
		Int("following", len(following)).
		Msg("Got users followed")

	// Convert following to a slice of uuid.UUID
	var followingIDs []*uuid.UUID
	for _, user := range following {
		followingIDs = append(followingIDs, &user.ID)
	}
	logger.Log.Debug().
		Int("followingIDs", len(followingIDs)).
		Msg("Got following IDs")

	// Include the authenticated user's ID in the list of IDs to get posts from
	followingIDs = append(followingIDs, requestorID)

	// Get the posts from the users that the authenticated user is following
	posts, err := db.Posts.GetByUserIDs(followingIDs, cursor, 25, requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting posts")
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	logger.Log.Debug().
		Int("posts", len(posts)).
		Msg("Got posts")

	// Return the posts
	return c.JSON(response.CreatePostsResponse(posts))
}
