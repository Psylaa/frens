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

	requestorId := c.Locals("requestorId").(*uuid.UUID)

	// Get extension of file
	ext := filepath.Ext(file.Filename)

	// Create file object
	logger.Log.Debug().
		Str("package", "service").
		Str("service", "files").
		Str("method", "Create").
		Str("file_extension", ext).
		Str("user_id", requestorId.String()).
		Msg("Creating file object")

	fileObj := &database.File{
		ID:        uuid.New(),
		Extension: ext,
		UserID:    *requestorId,
	}

	// Create file record in database
	logger.Log.Debug().
		Str("package", "service").
		Str("service", "files").
		Str("method", "Create").
		Str("file_id", fileObj.ID.String()).
		Str("file_extension", fileObj.Extension).
		Str("user_id", fileObj.UserID.String()).
		Msg("Creating file record in database")

	fileData, err := db.Files.Create(fileObj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Make directory if it doesn't exist
	logger.Log.Debug().
		Str("package", "service").
		Str("service", "files").
		Str("method", "Create").
		Str("file_id", fileData.ID.String()).
		Str("file_extension", fileData.Extension).
		Str("user_id", fileData.UserID.String()).
		Str("directory", cfg.Storage.Local.Path).
		Msg("Creating directory if it doesn't exist")
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

func (fr *FilesRepo) RetrieveByID(c *fiber.Ctx, fileId *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "files", "Retrieve")

	// Get file data
	logger.Log.Debug().
		Str("package", "service").
		Str("service", "files").
		Str("method", "RetrieveByID").
		Str("file_id", fileId.String()).
		Msg("Getting file data")
	fileData, err := db.Files.GetByID(fileId)
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
		Str("file_id", fileId.String()).
		Str("file_extension", fileData.Extension).
		Str("user_id", fileData.UserID.String()).
		Str("directory", cfg.Storage.Local.Path).
		Msg("Checking if file exists on disk")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return file
	return c.SendFile(filePath)
}

func (fr *FilesRepo) DeleteByID(c *fiber.Ctx, fileId *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "files", "Delete")

	// Get file data
	logger.Log.Debug().
		Str("package", "service").
		Str("service", "files").
		Str("method", "DeleteByID").
		Str("file_id", fileId.String()).
		Msg("Getting file data")
	fileData, err := db.Files.GetByID(fileId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	if fileData == nil {
		return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
	}

	// Delete file from storage
	logger.Log.Debug().
		Str("package", "service").
		Str("service", "files").
		Str("method", "DeleteByID").
		Str("file_id", fileId.String()).
		Str("file_extension", fileData.Extension).
		Str("user_id", fileData.UserID.String()).
		Str("directory", cfg.Storage.Local.Path).
		Msg("Deleting file from storage")

	if err := os.Remove(filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+fileData.Extension)); err != nil {
		logger.Log.Error().
			Str("package", "service").
			Str("service", "files").
			Str("method", "DeleteByID").
			Str("file_id", fileId.String()).
			Str("file_extension", fileData.Extension).
			Str("user_id", fileData.UserID.String()).
			Str("directory", cfg.Storage.Local.Path).
			Msg("Failed to delete file from storage")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Delete file from database
	if err := db.Files.DeleteByID(fileId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return success
	return c.JSON(response.CreateFilesResponse([]*database.File{fileData}))
}
