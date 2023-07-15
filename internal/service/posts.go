package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostRepo struct {
	Database *database.Database
}

func (pr *PostRepo) Create(c *fiber.Ctx, req *models.CreatePostRequest) error {
	logger.Debug(logger.LogMessage{
		Package:  "service",
		Function: "PostRepo.Create",
		Message:  "Creating post",
	})

	// Sanitize and validate request
	req.Sanitize()
	err := req.Validate()
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	// Get requesting user
	requestorID, ok := c.Locals("requestorID").(*uuid.UUID)
	if !ok {
		logger.Error(logger.LogMessage{
			Package:  "service",
			Function: "PostRepo.Create",
			Message:  "Error parsing requestorID from context.",
		}, nil)
		return models.ErrInternalServerError.SendResponse(c, "Error parsing requestorID from context.")
	}

	// Convert request to post
	newPost, err := req.ToPost(requestorID)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c)
	}

	// Create post in database
	if err := pr.Database.Posts.Create(newPost); err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Convert to response
	resp := newPost.ToResponse()

	// Send response
	return resp.Send(c)
}
