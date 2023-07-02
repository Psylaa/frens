package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookmarkRepo struct{}

func (br *BookmarkRepo) GetByBookmarkID(c *fiber.Ctx, bookmarkID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByBookmarkID")

	// Get bookmark from database
	bookmark, err := db.Bookmarks.GetByID(bookmarkID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetByBookmarkID", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) GetByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByPostID")

	// Get bookmarks from database
	bookmarks, err := db.Bookmarks.GetByPostID(postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetByPostID", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse(bookmarks))
}

func (br *BookmarkRepo) GetCountByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetCountByPostID")

	// Get bookmark count from database
	count, err := db.Bookmarks.GetCountByPostID(postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetCountByPostID", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateCountResponse(count))
}

func (br *BookmarkRepo) GetCountByUserID(c *fiber.Ctx, userID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetCountByUserID")

	// Get bookmark count from database
	count, err := db.Bookmarks.GetCountByUserID(userID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetCountByUserID", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	resp := response.CreateCountResponse(count)
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (br *BookmarkRepo) DeleteByID(c *fiber.Ctx, userID *uuid.UUID, bookmarkID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "DeleteByID")

	// Delete bookmark from database
	bookmark, err := db.Bookmarks.DeleteByID(userID, bookmarkID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "DeleteByID", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))

	}

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}
