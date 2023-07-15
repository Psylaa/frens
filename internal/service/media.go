package service

import (
	"mime/multipart"
	"path/filepath"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MediaRepo struct {
	Database *database.Database
}

func (mr *MediaRepo) Create(c *fiber.Ctx, file *multipart.FileHeader) error {

	// Get requesting user
	requestorID, ok := c.Locals("requestorID").(*uuid.UUID)
	if !ok {
		logger.Error(logger.LogMessage{
			Package:  "service",
			Function: "MediaRepo.Create",
			Message:  "Error parsing requestorID from context.",
		}, nil)
		return models.ErrInternalServerError.SendResponse(c, "Error parsing requestorID from context.")
	}

	// Extract file extension
	extension := filepath.Ext(file.Filename)

	// Create Media object
	media := &models.Media{
		UserID:    *requestorID,
		Extension: extension,
	}

	// Save Media object to database
	newMedia, err := mr.Database.Media.Create(media)
	if err != nil {
		logger.Error(logger.LogMessage{
			Package:  "service",
			Function: "MediaRepo.Create",
			Message:  "Error creating media object.",
		}, err)
		return models.ErrInternalServerError.SendResponse(c, "Error creating media object.")
	}

	// Save file to disk
	// Todo. Create a storage package to handle this.

	// Convert to response
	resp := newMedia.ToResponse()

	// Send response
	return resp.Send(c)
}
