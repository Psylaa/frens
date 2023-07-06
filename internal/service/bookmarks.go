package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookmarkRepo struct{}

func (br *BookmarkRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID, count *int, offset *int) error {
	logger.DebugLogRequestReceived("service", "bookmark", "GetByUserID")

	// Get bookmarks from database
	bookmarks, err := db.Bookmarks.GetByUserID(userID, count, offset)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "GetByUserID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "GetByUserID", "bookmark found")

	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse(bookmarks))
}

func (br *BookmarkRepo) Create(c *fiber.Ctx, userID *uuid.UUID, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "bookmark", "Create")

	// Construct a new bookmark
	newBookmark := &database.Bookmark{
		UserID: *userID,
		PostID: *postID,
	}

	// Insert bookmark into database
	err := db.Bookmarks.Create(newBookmark)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "Create", "bookmark not created", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Get bookmark from database with owner
	bookmark, err := db.Bookmarks.GetByPostAndUserID(userID, postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "Create", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "Create", "bookmark created")

	// Return bookmark
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) Delete(c *fiber.Ctx, userID *uuid.UUID, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "bookmark", "DeleteByUserAndPostID")

	// Get bookmark from database
	bookmark, err := db.Bookmarks.GetByPostAndUserID(userID, postID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "DeleteByUserAndPostID", "bookmark not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Delete bookmark from database
	err = db.Bookmarks.DeleteByID(&bookmark.ID)
	if err != nil {
		logger.ErrorLogRequestError("service", "bookmark", "DeleteByUserAndPostID", "bookmark not deleted", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "bookmark", "DeleteByUserAndPostID", "bookmark deleted")

	// Return bookmark
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarkResponse([]*database.Bookmark{bookmark}))
}
