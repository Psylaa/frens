package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LikeRepo struct{}

func (lr *LikeRepo) Create(c *fiber.Ctx, postId *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "like", "Create")

	// Get the user id
	userId := c.Locals("requestorId").(*uuid.UUID)

	// Check if the user has already liked the post
	liked, err := db.Likes.Exists(userId, postId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Create", "Error checking if like exists", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Create", "Like exists")

	// If the user has already liked the post, return an error
	if liked {
		logger.ErrorLogRequestError("service", "like", "Create", "Like already exists", err)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrExists))
	}
	logger.DebugLogRequestUpdate("service", "like", "Create", "Like does not exist")

	// Create the like
	like, err := db.Likes.Create(userId, postId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Create", "Error creating like", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Create", "Like created")

	// Retrieve the user - probably can be replaced at some point by preloading the user directly in the create function
	user, err := db.Users.GetByID(userId, userId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Create", "Error retrieving user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Create", "User retrieved")

	// Retrieve the post - probably can be replaced at some point by preloading the post directly in the create function
	post, err := db.Posts.GetByID(userId, postId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Create", "Error retrieving post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Create", "Post retrieved")

	// Return the like
	return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse(user, post, []*database.Like{like}))
}

func (lr *LikeRepo) Delete(c *fiber.Ctx, postId *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "like", "Delete")

	// Get the user id
	userId := c.Locals("requestorId").(*uuid.UUID)

	// Check if the user has already liked the post
	liked, err := db.Likes.Exists(userId, postId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Delete", "Error checking if like exists", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Delete", "Like exists")

	// If the user has not liked the post, return an error
	if !liked {
		logger.ErrorLogRequestError("service", "like", "Delete", "Like does not exist", err)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}
	logger.DebugLogRequestUpdate("service", "like", "Delete", "Like does exist")

	// Delete the like
	err = db.Likes.Delete(userId, postId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Delete", "Error deleting like", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Delete", "Like deleted")

	// Retrieve the user - probably can be replaced at some point by preloading the user directly in the create function
	user, err := db.Users.GetByID(userId, userId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Delete", "Error retrieving user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Delete", "User retrieved")

	// Retrieve the post - probably can be replaced at some point by preloading the post directly in the create function
	post, err := db.Posts.GetByID(userId, postId)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Delete", "Error retrieving post", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "like", "Delete", "Post retrieved")

	// Return the like
	return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse(user, post, []*database.Like{}))
}