package router

import (
	"net/http"
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getChronoFeed(c *fiber.Ctx) error {

	// Get the user ID from the JWT
	userID, err := getUserID(c)
	if err != nil || userID == uuid.Nil {
		logger.Log.Error().Err(err).Msg("Error getting user ID")
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Get the cursor from the query string
	cursorString := c.Query("cursor")
	var cursor time.Time
	if cursorString == "" {
		cursor = time.Now()
	} else {
		var err error
		cursor, err = time.Parse(time.RFC3339, cursorString)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing cursor")
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	return srv.Feed.GetChrono(c, &userID, cursor)
}

func getAlgoFeed(c *fiber.Ctx) error {
	return nil
}

// getExploreFeed returns a list of the latest posts from all users
func getExploreFeed(c *fiber.Ctx) error {
	/*
		posts, err := db.Posts.GetLatestPublicPosts(25)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error getting posts")
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		logger.Log.Debug().Int("posts", len(posts)).Msg("Got posts")

		return c.JSON(response.GeneratePostsResponse(posts))
	*/
	return nil
}
