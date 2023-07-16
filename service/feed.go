package service

import (
	"time"

	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
)

type FeedService struct{ Database *database.Database }

func (f *FeedService) GetChronological(c *fiber.Ctx, count uint8, cursor uint32) error {

	// Set default values for count and cursor if they are not provided
	if count == 0 {
		count = 10
	}
	if cursor == 0 {
		cursor = uint32(time.Now().Unix())
	}

	// get the user ID from the context
	userID, err := getRequestorID(c)
	if err != nil {
		return err
	}

	// get the IDs of all users the current user is following
	var follows []models.Follow
	f.Database.Conn.Where("user_id = ?", userID).Find(&follows)

	followedIDs := make([]uint32, len(follows))
	for i, follow := range follows {
		followedIDs[i] = follow.FollowedID
	}

	// add the current user's ID to the list
	followedIDs = append(followedIDs, userID)

	// get all posts from these users, before the cursor date
	var posts []*models.Post
	f.Database.Conn.
		Preload("User").
		Where("user_id IN (?) AND created_at < ?", followedIDs, time.Unix(int64(cursor), 0)).
		Order("created_at desc").
		Limit(int(count)).
		Find(&posts)

	// render the posts as JSON API
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)
	if err := jsonapi.MarshalPayload(c.Response().BodyWriter(), posts); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal the posts",
		})
	}
	return nil
}
