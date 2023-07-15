package service

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

const defaultCount = 25

type FeedRepo struct{ Database *database.Database }

func (f *FeedRepo) GetChrono(c *fiber.Ctx, req models.FeedRequest) error {
	// If Count is not provided, we set it to the default value.
	if req.Count == nil {
		defaultCount := defaultCount
		req.Count = &defaultCount
	}

	// If Cursor is not provided, we set it to the current time.
	if req.Cursor == nil {
		now := time.Now()
		req.Cursor = &now
	}

	posts, err := f.Database.Posts.Read(req.Count, req.Cursor)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c)
	}

	// Convert posts to response format and collect all users.
	var data []models.PostData
	var included []models.UserData
	includedMap := make(map[string]struct{}) // Used to track included users.

	for _, post := range posts {
		postData, userData := post.ToResponseData()

		// Check if user is already included.
		if _, ok := includedMap[userData.ID.String()]; !ok {
			// If not, add user to included slice and map.
			included = append(included, userData)
			includedMap[userData.ID.String()] = struct{}{}
		}

		// Add post data to data slice.
		data = append(data, postData)
	}

	// Create and send response.
	resp := models.CreateFeedResponse(data, included)
	return c.Status(fiber.StatusOK).JSON(resp)
}
