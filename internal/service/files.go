package service

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FilesRepo struct{}

func (fr *FilesRepo) GetByID(c *fiber.Ctx) error {
	return nil
}

func (fr *FilesRepo) Create(c *fiber.Ctx, userID *uuid.UUID, file *multipart.FileHeader) error {
	logger.DebugLogRequestRecieved("service", "files", "Create")

	// Get extension of file
	ext := filepath.Ext(file.Filename)

	// Create file object
	fileObj := &database.File{
		ID:        uuid.New(),
		Extension: ext,
		UserID:    *userID,
	}

	// Create file record in database
	fileData, err := db.Files.Create(fileObj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Make directory if it doesn't exist
	if err := os.MkdirAll(cfg.Storage.Local.Path, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Save file to directory
	if err := c.SaveFile(file, filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+ext)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return file data
	return c.JSON(response.GenerateFilesResponse([]*database.File{fileData}))
}
