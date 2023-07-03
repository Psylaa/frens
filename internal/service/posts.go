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
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Get posts from database
	posts, err := db.Posts.GetByUserID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return the posts
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse(posts))
}

func (pr *PostRepo) GetReplies() {

}

func (pr *PostRepo) Create(
	c *fiber.Ctx,
	userID uuid.UUID,
	text string,
	privacy shared.Privacy,
	mediaIDs []*uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "post", "Create")

	// Set default privacy to public if not provided.
	if privacy == "" {
		privacy = shared.PrivacyPublic
	}

	// Convert the media IDs files
	mediaFiles, err := db.Files.GetManyByID(mediaIDs)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "Create", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Create post in database
	post, err := db.Posts.Create(userID, text, privacy, mediaFiles)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "Create", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Retrieve the post so we can return the author's information.
	post, err = db.Posts.GetByID(post.ID)
	if err != nil {
		logger.ErrorLogRequestError("service", "post", "Create", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return the post
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse([]*database.Post{post}))
}
