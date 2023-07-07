package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostRepo struct{}

func (pr *PostRepo) Get(c *fiber.Ctx, userID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "post", "Get")

	return nil
}

func (ur *PostRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID) error {
	return nil
}

func (pr *PostRepo) GetReplies() {

}

func (pr *PostRepo) Create(c *fiber.Ctx, text string, privacy shared.Privacy, mediaIDs []string) error {
	logger.DebugLogRequestReceived("service", "post", "Create")

	// Get the userID from the token.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Construct the post object
	post := &database.Post{
		BaseModel: database.BaseModel{
			ID: uuid.New(),
		},
		AuthorID: *requestorID,
		Text:     text,
		Privacy:  privacy,
		MediaIDs: mediaIDs,
	}

	// Insert the post into the database
	err := db.Posts.Create(post)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error creating post")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Send the response
	return c.Status(fiber.StatusCreated).JSON(response.CreatePostsResponse([]*database.Post{post}))
}
