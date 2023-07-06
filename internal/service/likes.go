package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LikeRepo struct{}

func (lr *LikeRepo) GetByID(c *fiber.Ctx, postId *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "like", "GetByID")

		// Get the like
		like, err := db.Likes.GetByID(postId)
		if err != nil {
			logger.ErrorLogRequestError("service", "like", "GetByID", "Error getting like", err)
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		logger.DebugLogRequestUpdate("service", "like", "GetByID", "Like retrieved")

		// Return the like
		return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse(nil, nil, []*database.Like{like}))
	*/
	return nil
}

func (lr *LikeRepo) GetByPostID(c *fiber.Ctx, postId *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "like", "GetByPostID")

		// Get the likes
		likes, err := db.Likes.GetByPostID(postId)
		if err != nil {
			logger.ErrorLogRequestError("service", "like", "GetByPostID", "Error getting likes", err)
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		logger.DebugLogRequestUpdate("service", "like", "GetByPostID", "Likes retrieved")

		// Return the likes
		return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse(nil, nil, likes))
	*/
	return nil
}

func (lr *LikeRepo) GetByPostIDAndUserID(c *fiber.Ctx, postId *uuid.UUID, userId *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "like", "GetByPostIDAndUserID")

		// Get the likes
		like, err := db.Likes.GetByPostIDAndUserID(postId, userId)
		if err != nil {
			logger.ErrorLogRequestError("service", "like", "GetByPostIDAndUserID", "Error getting likes", err)
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		logger.DebugLogRequestUpdate("service", "like", "GetByPostIDAndUserID", "Likes retrieved")

		// Return the likes
		return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse(nil, nil, []*database.Like{like}))
	*/
	return nil
}

func (lr *LikeRepo) Create(c *fiber.Ctx, postId *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "like", "Create")

		// Get the user id
		userId := c.Locals("requestorId").(*uuid.UUID)

		// Check if the user has already liked the post
		liked, err := db.Likes.ExistsByPostAndUserID(postId, userId)
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
	*/
	return nil
}

func (lr *LikeRepo) Delete(c *fiber.Ctx, postId *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "like", "Delete")

		// Get the user id
		userId := c.Locals("requestorId").(*uuid.UUID)

		// Check if the user has already liked the post
		liked, err := db.Likes.ExistsByPostAndUserID(postId, userId)
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
	*/
	return nil
}
