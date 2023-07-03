package router

import (
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	_ "github.com/bwoff11/frens/docs"
)

// @Summary Retrieve a bookmark by ID
// @Description Get the details of a specific bookmark based on the provided ID
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param bookmarkId path string true "Bookmark ID"
// @Security ApiKeyAuth
// @Router /bookmarks/{bookmarkId} [get]
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
	// Since this is only callable by owner, we can get the user ID from the token
	return srv.Bookmarks.GetByUserID(c, c.Locals("requestorId").(*uuid.UUID))
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
	// Since this is only callable by owner, we can get the user ID from the token
	return srv.Bookmarks.GetCountByUserID(c, c.Locals("requestorId").(*uuid.UUID))
}

func createBookmarkbyPostID(c *fiber.Ctx) error {
	postId := c.Params("postId")
	postID, err := uuid.Parse(postId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing post ID: " + postId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.Create(c, c.Locals("requestorID").(*uuid.UUID), &postID)
}

func deleteBookmarkByID(c *fiber.Ctx) error {
	bookmarkId := c.Params("bookmarkId")
	bookmarkID, err := uuid.Parse(bookmarkId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing bookmark ID: " + bookmarkId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.DeleteByID(c, c.Locals("requestorId").(*uuid.UUID), &bookmarkID)
}

func deleteBookmarkByPostID(c *fiber.Ctx) error {

	// Get post ID from request
	id := c.Params("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Send request to service
	return srv.Bookmarks.DeleteByUserAndPostID(c, c.Locals("requestorId").(*uuid.UUID), &postID)
}
