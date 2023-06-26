package router

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// getChronologicalFeed returns a list of statuses from the users that the
// authenticated user is following, in chronological order.
func getChronologicalFeed(c *fiber.Ctx) error {
	// Get the user ID from the JWT.
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	// Get the cursor from the request query parameters. If it's not present,
	// we default to the current time, which will retrieve the most recent statuses.
	cursorParam := c.Query("cursor")
	cursor := time.Now()
	if cursorParam != "" {
		unixTime, err := strconv.ParseInt(cursorParam, 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid cursor")
		}
		cursor = time.Unix(unixTime, 0)
	}

	// Here is where you'd get the list of users that the authenticated user is
	// following. This depends on your data storage, so replace this with your
	// actual implementation.
	following, err := db.GetFollowing(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Extract the following IDs
	followingIDs := make([]uuid.UUID, len(following))
	for i, follower := range following {
		followingIDs[i] = follower.FollowingID
	}

	followingIDs = append(followingIDs, userID)

	statuses, err := db.GetStatusesByUserIDs(followingIDs, cursor, 10)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(statuses)
}
