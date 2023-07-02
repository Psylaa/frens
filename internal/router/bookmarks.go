package router

import (
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func getBookmarkByID(c *fiber.Ctx) error {
	bookmarkId := c.Params("bookmarkId")
	bookmarkID, err := uuid.Parse(bookmarkId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing bookmark ID: " + bookmarkId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.GetByID(c, &bookmarkID)
}

func getBookmarksByPostID(c *fiber.Ctx) error {
	postId := c.Params("postId")
	postID, err := uuid.Parse(postId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing post ID: " + postId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.GetByPostID(c, &postID)
}

func getBookmarksByUserID(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	return srv.Bookmarks.GetByUserID(c, &userID)
}

func getBookmarksCountByPostID(c *fiber.Ctx) error {
	postId := c.Params("postId")
	postID, err := uuid.Parse(postId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing post ID: " + postId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.GetCountByPostID(c, &postID)
}

func getBookmarksCountByUserID(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	return srv.Bookmarks.GetCountByUserID(c, &userID)
}

func createBookmarkbyPostID(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	postId := c.Params("postId")
	postID, err := uuid.Parse(postId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing post ID: " + postId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.Create(c, &userID, &postID)
}

func deleteBookmarkByID(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	bookmarkId := c.Params("bookmarkId")
	bookmarkID, err := uuid.Parse(bookmarkId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing bookmark ID: " + bookmarkId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.DeleteByID(c, &userID, &bookmarkID)
}

func deleteBookmarkByPostID(c *fiber.Ctx) error {

	// Get user ID from token
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Get post ID from request
	id := c.Params("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Send request to service
	return srv.Bookmarks.DeleteByUserAndPostID(c, &userID, &postID)
}
