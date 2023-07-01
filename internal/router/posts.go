package router

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// getPost handles the HTTP request to fetch a specific post.
func getPost(c *fiber.Ctx) error {
	postID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}
	logger.Log.Debug().
		Str("package", "router").
		Str("func", "getPost").
		Msgf("successfully parsed provided post id into uuid: %v", postID)

	post, err := db.Posts.GetPost(postID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusNotFound).JSON(response.GenerateErrorResponse(response.ErrNotFound))
	default:
		logger.Log.Error().Err(err).
			Str("package", "router").
			Str("func", "getPost").
			Msg("unknown error fetching post from database")
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.JSON(response.GeneratePostResponse(post))
}

// getPosts handles the HTTP request to fetch all posts by a user.
func GetPostsByUserID(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Query("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}
	logger.Log.Debug().Msgf("successfully parsed provided user id into uuid: %v", userID)

	posts, err := db.Posts.GetPostsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}
	return c.JSON(response.GeneratePostsResponse(posts))
}

// createPost handles the HTTP request to create a new post.
func createPost(c *fiber.Ctx) error {
	var body struct {
		Text     string         `json:"text"`
		Privacy  shared.Privacy `json:"privacy"`
		MediaIDs []string       `json:"mediaIds"`
	}
	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidBody))
	}

	// Set default privacy to public if not provided.
	if body.Privacy == "" {
		body.Privacy = shared.PrivacyPublic
	}

	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
	}
	logger.Log.Debug().Msgf("userID: %v", userID)

	// Convert the media IDs to UUIDs.
	mediaIDs, err := shared.UUIDsFromStrings(body.MediaIDs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	// Convert the media IDs files
	mediaFiles, err := db.Files.GetFiles(mediaIDs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	post, err := db.Posts.CreatePost(userID, body.Text, body.Privacy, mediaFiles)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	// Retrieve the post so we can return the author's information.
	post, err = db.Posts.GetPost(post.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.JSON(response.GeneratePostResponse(post))
}

// deletePost handles the HTTP request to delete a post.
func deletePost(c *fiber.Ctx) error {
	// Parse the post ID from the URL parameter.
	postID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	// Get the user ID from the JWT.
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
	}

	// First, check if the post exists and belongs to the user.
	post, err := db.Posts.GetPost(postID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(response.GenerateErrorResponse(response.ErrNotFound))
		}
	}

	// Check if the user owns the post.
	if post.AuthorID != userID {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrUnauthorized))
	}

	// Delete the post.
	err = db.Posts.DeletePost(postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.SendStatus(fiber.StatusNoContent)
}
