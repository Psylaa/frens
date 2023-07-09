package service

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookmarkRepo struct{}

func (br *BookmarkRepo) Get(c *fiber.Ctx, count int, cursor time.Time) error {
	logger.DebugLogRequestReceived("service", "bookmark", "Get")

	// Get all bookmarks
	bookmarks, err := db.Bookmarks.Get(count, &cursor)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting bookmarks")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Convert to a list of postIDs
	var postIDs []*uuid.UUID
	for _, bookmark := range bookmarks {
		postIDs = append(postIDs, &bookmark.PostID)
	}

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get all posts
	posts, err := db.Posts.GetByIDs(postIDs, requestorID)

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreatePostsResponse(posts))
}

func (br *BookmarkRepo) Create(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "bookmark", "Create")

	// Get the userID from the token.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Construct the bookmark object
	bookmark := &database.Bookmark{
		BaseModel: database.BaseModel{
			ID: uuid.New(),
		},
		UserID: *requestorID,
		PostID: *postID,
	}

	// Insert the bookmark into the database
	err := db.Bookmarks.Create(bookmark)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error creating bookmark")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarksResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) Delete(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "bookmark", "Delete")

	// Get the userID from the token.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the bookmark
	bookmark, err := db.Bookmarks.GetByPostIDAndUserID(postID, requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting bookmark")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Delete the bookmark
	err = db.Bookmarks.Delete(bookmark)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error deleting bookmark")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarksResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) GetByID(c *fiber.Ctx, id *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "bookmark", "GetByID")

	// Make sure the user owns the bookmark
	isOwner := db.Bookmarks.Exists(id, c.Locals("requestorID").(*uuid.UUID))
	if !isOwner {
		return c.Status(fiber.StatusForbidden).JSON(response.CreateErrorResponse(response.ErrForbidden))
	}

	// Retrieve the bookmark from the database
	bookmark, err := db.Bookmarks.GetByID(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting bookmark")
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarksResponse([]*database.Bookmark{bookmark}))
}

func (br *BookmarkRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID, count, offset *int) error {
	logger.DebugLogRequestReceived("service", "bookmark", "GetByUserID")

	// Get the userID from the token.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get all bookmarks owned by the user
	bookmarks, err := db.Bookmarks.GetByUserID(requestorID, count, offset)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting bookmarks")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Filter bookmarks that belong to the user
	var userBookmarks []*database.Bookmark
	for _, bookmark := range bookmarks {
		if bookmark.UserID == *userID {
			userBookmarks = append(userBookmarks, bookmark)
		}
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarksResponse(userBookmarks))
}

func (br *BookmarkRepo) DeleteByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "bookmark", "DeleteByPostID")

	// Get the userID from the token.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the bookmark
	bookmark, err := db.Bookmarks.GetByPostIDAndUserID(postID, requestorID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting bookmark")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Delete the bookmark
	err = db.Bookmarks.Delete(bookmark)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error deleting bookmark")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Send the response
	return c.Status(fiber.StatusOK).JSON(response.CreateBookmarksResponse([]*database.Bookmark{bookmark}))
}
