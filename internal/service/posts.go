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

func (pr *PostRepo) GetByID() {

}

func (ur *PostRepo) GetByUserID(c *fiber.Ctx, userId *uuid.UUID) error {

	// Verify that the user exists
	if !db.Users.Exists(userId) {
		logger.ErrorLogRequestError("service", "post", "GetByUserID", "user not found", nil)
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}
	logger.DebugLogRequestUpdate("service", "post", "GetByUserID", "user found")

	// Get posts from database
	posts, err := db.Posts.GetByUserID(userId)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "GetByUserID", "posts not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "post", "GetByUserID", "posts found")

	// Return the posts
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse(posts))
}

func (pr *PostRepo) GetReplies() {

}

func (pr *PostRepo) Create(
	c *fiber.Ctx,
	text string,
	privacy shared.Privacy,
	mediaIDs []*uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "post", "Create")

	requestorId := c.Locals("requestorId").(*uuid.UUID)

	// Set default privacy to public if not provided.
	if privacy == "" {
		logger.DebugLogRequestUpdate("service", "post", "Create", "privacy not provided, setting to public")
		privacy = shared.PrivacyPublic
	}
	logger.DebugLogRequestUpdate("service", "post", "Create", "privacy set")

	// Convert the media IDs files
	mediaFiles, err := db.Files.GetManyByID(mediaIDs)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "Create", "error getting media files", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "post", "Create", "media files retrieved")

	// Create post in database
	post, err := db.Posts.Create(*c.Locals("requestorId").(*uuid.UUID), text, privacy, mediaFiles)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "Create", "error creating post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "post", "Create", "post created")

	// Retrieve the post so we can return the author's information.
	post, err = db.Posts.GetByID(requestorId, &post.ID)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "Create", "error retrieving post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "post", "Create", "post retrieved")

	// Return the post
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse([]*database.Post{post}))
}
