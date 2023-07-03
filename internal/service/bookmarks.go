package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookmarkRepo struct{}

func (br *BookmarkRepo) GetByID(c *fiber.Ctx, bookmarkID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByBookmarkID")

	// Get bookmark from database
	bookmark, err := db.Bookmarks.GetByID(bookmarkID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetByBookmarkID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "GetByBookmarkID", "bookmark found")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByUserID")

	// Get bookmarks from database
	bookmarks, err := db.Bookmarks.GetByUserID(userID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetByUserID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "GetByUserID", "bookmark found")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse(bookmarks))
}

func (br *BookmarkRepo) GetByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByPostID")

	// Get bookmarks from database
	bookmarks, err := db.Bookmarks.GetByPostID(postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetByPostID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "GetByPostID", "bookmark found")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse(bookmarks))
}

func (br *BookmarkRepo) GetCountByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetCountByPostID")

	// Get bookmark count from database
	count, err := db.Bookmarks.GetCountByPostID(postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetCountByPostID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "GetCountByPostID", "bookmark found")

	return c.Status(fiber.StatusOK).JSON(response.CreateCountResponse(count))
}

func (br *BookmarkRepo) GetCountByUserID(c *fiber.Ctx, userID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetCountByUserID")

	// Get bookmark count from database
	count, err := db.Bookmarks.GetCountByUserID(userID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetCountByUserID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "GetCountByUserID", "bookmark found")

	resp := response.CreateCountResponse(count)
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (br *BookmarkRepo) Create(c *fiber.Ctx, userID *uuid.UUID, postID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "Create")

	// Create bookmark in database
	bookmark, err := db.Bookmarks.Create(userID, postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "Create", "bookmark not created", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "Create", "bookmark created")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) DeleteByID(c *fiber.Ctx, userID *uuid.UUID, bookmarkID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "DeleteByID")

	// Delete bookmark from database
	bookmark, err := db.Bookmarks.DeleteByID(userID, bookmarkID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "DeleteByID", "bookmark not deleted", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "DeleteByID", "bookmark deleted")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) DeleteByUserAndPostID(c *fiber.Ctx, userID *uuid.UUID, postID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "bookmark", "DeleteByUserAndPostID")

	// Delete bookmark from database
	bookmark, err := db.Bookmarks.DeleteByUserAndPostID(userID, postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "DeleteByUserAndPostID", "bookmark not deleted", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "DeleteByUserAndPostID", "bookmark deleted")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}
