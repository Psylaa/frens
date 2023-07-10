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

func (fr *FilesRepo) Create(c *fiber.Ctx, file *multipart.FileHeader) error {
	logger.DebugLogRequestReceived("service", "files", "Create")

	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get extension of file
	ext := filepath.Ext(file.Filename)

	// Create file object

	fileObj := &database.File{
		BaseModel: database.BaseModel{
			ID: uuid.New(),
		},
		Extension: ext,
		OwnerID:   *requestorID,
	}

	// Create file record in database
	err := db.Files.Create(fileObj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Retrieve file data from database with owner
	fileData, err := db.Files.GetByID(&fileObj.ID)
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
	return c.JSON(response.CreateFilesResponse([]*database.File{fileData}))
}

func (fr *FilesRepo) RetrieveByID(c *fiber.Ctx, fileID *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "files", "Retrieve")

		// Get file data
		logger.Log.Debug().
			Str("package", "service").
			Str("service", "files").
			Str("method", "RetrieveByID").
			Str("file_id", fileID.String()).
			Msg("Getting file data")
		fileData, err := db.Files.GetByID(fileID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		// Get file path
		filePath := filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+fileData.Extension)

		// Check if file exists
		logger.Log.Debug().
			Str("package", "service").
			Str("service", "files").
			Str("method", "RetrieveByID").
			Str("file_id", fileID.String()).
			Str("file_extension", fileData.Extension).
			Str("user_id", fileData.UserID.String()).
			Str("directory", cfg.Storage.Local.Path).
			Msg("Checking if file exists on disk")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		// Return file
		return c.SendFile(filePath)
	*/
	return nil
}

func (fr *FilesRepo) DeleteByID(c *fiber.Ctx, fileID *uuid.UUID) error {
	/*
		logger.DebugLogRequestReceived("service", "files", "Delete")

		// Get file data
		fileData, err := db.Files.GetByID(fileID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		if fileData == nil {
			return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
		}

		// Delete file from storage

		if err := os.Remove(filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+fileData.Extension)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		// Delete file from database
		if err := db.Files.DeleteByID(fileID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		// Return success
		return c.JSON(response.CreateFilesResponse([]*database.File{fileData}))
	*/
	return nil
}
