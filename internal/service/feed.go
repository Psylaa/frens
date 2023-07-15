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

	// Call the Read function from PostRepository using the provided or default values.
	posts, err := f.Database.Posts.Read(req.Count, req.Cursor)

	if err != nil {
		// handle the error, maybe return a 500 status with some information
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve posts",
		})
	}

	// If there's no error, send the posts as a response.
	return c.Status(fiber.StatusOK).JSON(posts)
}
