package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LikeRepo struct{ Database *database.Database }

func (lr *LikeRepo) Create(c *fiber.Ctx, req *models.CreateLikeRequest) error {
	logger.Debug(logger.LogMessage{
		Package:  "service",
		Function: "LikeRepo.Create",
		Message:  "Creating like",
	})

	// Validate request
	err := req.Validate()
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	// Get requesting user
	requestorID, ok := c.Locals("requestorID").(*uuid.UUID)
	if !ok {
		logger.Error(logger.LogMessage{
			Package:  "service",
			Function: "LikeRepo.Create",
			Message:  "Error parsing requestorID from context.",
		}, nil)
		return models.ErrInternalServerError.SendResponse(c, "Error parsing requestorID from context.")
	}

	// Convert request to like
	newLike, err := req.ToLike(requestorID)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c)
	}

	// Create like in database and get the like with preloaded user
	newLike, err = lr.Database.Likes.Create(newLike)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Convert to response and send
	likeData := newLike.ToResponseData()
	resp := models.CreateLikeResponse(likeData)
	return c.Status(fiber.StatusCreated).JSON(resp)
}
