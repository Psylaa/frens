package router

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func getBookmarkByID(c *fiber.Ctx) error {
	logger.DebugLogRequestRecieved("router", "bookmark", "getBookmarkByID")

	bookmarkId := c.Params("bookmarkId")
	bookmarkID, err := uuid.Parse(bookmarkId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing bookmark ID: " + bookmarkId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return srv.Bookmarks.GetByBookmarkID(c, &bookmarkID)
}

func getBookmarksByPostID(c *fiber.Ctx) error {
	return nil

	/*
		id := c.Params("postId")
		postID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		bookmarks, err := db.Bookmarks.GetBookmarksByIDs(postID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		resp := response.CreateBookmarksResponse(bookmarks)
		return c.JSON(resp)
	*/
}

// @Summary Get bookmark
// @Description Retrieve a specific bookmark
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param id path string true "Bookmark ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /bookmarks/{id} [get]
func getBookmark(c *fiber.Ctx) error {
	return nil
	/*
		bookmarkId := c.Params("id")
		bookmarkID, err := uuid.Parse(bookmarkId)
		if err != nil {
			log.Error().Err(err).Msg("Error parsing bookmark ID: " + bookmarkId)
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		bookmark, err := db.Bookmarks.GetBookmarkByID(&bookmarkID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.JSON(response.CreateBookmarkResponse(bookmark))
	*/
}

func createBookmark(c *fiber.Ctx) error {
	return nil
	/*
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
	*/
}

func deleteBookmark(c *fiber.Ctx) error {
	/*
		userID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
		}

		id := c.Params("id")
		postID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
		}

		if err := db.Bookmarks.DeleteBookmark(userID, postID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(fiber.StatusOK)
	*/

	return nil
}
