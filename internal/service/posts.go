package service

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostRepo struct{}

func (pr *PostRepo) Create(c *fiber.Ctx, text string, privacy shared.Privacy, mediaIDs []*uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "post", "Create")

	// Get the userID from the token.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Construct the post object
	post := &database.Post{
		BaseModel: database.BaseModel{
			ID: uuid.New(),
		},
		UserID:   *requestorID,
		Text:     text,
		Privacy:  privacy,
		MediaIDs: mediaIDs,
	}

	// Create the post and update the files
	err := db.Posts.CreateWithMedia(post, mediaIDs, db.Files)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error creating post")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Send the response
	return c.Status(fiber.StatusCreated).JSON(response.CreatePostsResponse([]*database.Post{post}))
}

func (pr *PostRepo) GetByID(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "post", "Get")

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Retrieve the posts from the database
	post, err := db.Posts.GetByID(postID, requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting post")
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse([]*database.Post{post}))
}

func (ur *PostRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID, cursor time.Time, count int) error {

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Retrieve the posts from the database
	posts, err := db.Posts.GetByUserIDs([]*uuid.UUID{userID}, cursor, count, requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting posts")
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse(posts))
}

func (pr *PostRepo) GetReplies() {

}

func (pr *PostRepo) Update() {}

func (pr *PostRepo) Delete(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "post", "Delete")

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the post
	post, err := db.Posts.GetByID(postID, requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting post")
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Verify that the requestor is the author of the post or an admin
	isAdmin := c.Locals("role").(shared.Role) == shared.RoleAdmin
	if post.AuthorID != *requestorID || !isAdmin {
		logger.Log.Error().Err(err).Msg("requestor is not the author of the post or an admin")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrUnauthorized))
	}

	// Send the request to the database layer
	err = db.Posts.Delete(post)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error deleting post")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Send the response
	return c.Status(fiber.StatusNoContent).JSON(response.CreatePostsResponse([]*database.Post{post}))
}
