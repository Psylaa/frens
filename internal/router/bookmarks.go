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
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	bookmark, err := db.Bookmarks.GetBookmarkByID(&bookmarkID)
	if err != nil {
		log.Error().Err(err).Msg("Error getting bookmark by ID: " + bookmarkId)
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.JSON(response.CreateBookmarkResponse(bookmark))
}

// @Summary Get post bookmarks
// @Description Retrieve bookmarks for a specific post
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param id path string true "Post ID"
// @Success 200 {object} response.BookmarkResp
// @Router /statuses/{id}/bookmarks [get]
func getBookmarksByPostID(c *fiber.Ctx) error {
	return nil

	/*
		id := c.Params("postId")
		statusID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
		}

		bookmarks, err := db.Bookmarks.GetBookmarksByIDs(statusID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
		}

		resp := response.CreateBookmarksResponse(bookmarks)
		return c.JSON(resp)
	*/
}

func getBookmark(c *fiber.Ctx) error {
	bookmarkId := c.Params("id")
	bookmarkID, err := uuid.Parse(bookmarkId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing bookmark ID: " + bookmarkId)
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	bookmark, err := db.Bookmarks.GetBookmarkByID(&bookmarkID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.JSON(response.CreateBookmarkResponse(bookmark))
}

func createBookmark(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	postId, err := uuid.Parse(c.Params("postId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	newBookmark, err := db.Bookmarks.CreateBookmark(userID, postId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	resp := response.CreateBookmarkResponse(newBookmark)
	return c.JSON(resp)
}

func deleteBookmark(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	id := c.Params("id")
	statusID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	if err := db.Bookmarks.DeleteBookmark(userID, statusID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func hasUserBookmarked(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	id := c.Params("id")
	statusId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	bookmarked, err := db.Bookmarks.HasUserBookmarked(userId, statusId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"bookmarked": bookmarked})
}
