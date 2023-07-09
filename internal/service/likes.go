package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LikeRepo struct{}

func (lr *LikeRepo) GetByID(c *fiber.Ctx, postID *uuid.UUID) error {
	return nil
}

func (lr *LikeRepo) GetByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	return nil
}

func (lr *LikeRepo) GetByPostIDAndUserID(c *fiber.Ctx, postID *uuid.UUID, userID *uuid.UUID) error {
	return nil
}

func (lr *LikeRepo) Create(c *fiber.Ctx, postID *uuid.UUID) error {

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Construct the like object
	newLike := &database.Like{
		BaseModel: database.BaseModel{
			ID: uuid.New(),
		},
		UserID: requestorID,
		PostID: postID,
	}

	// Insert the like into the database
	err := db.Likes.Create(newLike)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Create", "Error creating like", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse([]*database.Like{newLike}))
}

func (lr *LikeRepo) Delete(c *fiber.Ctx, postID *uuid.UUID) error {

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the like
	like, err := db.Likes.GetByPostIDAndUserID(postID, requestorID)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Delete", "Error getting like", err)
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Delete the like from the database
	err = db.Likes.Delete(like)
	if err != nil {
		logger.ErrorLogRequestError("service", "like", "Delete", "Error deleting like", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(response.CreateLikesResponse([]*database.Like{like}))
}
