package router

import (
	"io"
	"path/filepath"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func createFile(c *fiber.Ctx) error {
	// Get the user ID from the JWT token
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot get user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot get user ID",
			"message": err.Error(),
		})
	}

	// Parse the multipart form
	fileUpload, err := c.FormFile("file")
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot parse file upload")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse file upload",
			"message": err.Error(),
		})
	}

	// Read the uploaded file content
	file, err := fileUpload.Open()
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot open file")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot open file",
			"message": err.Error(),
		})
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot read file")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot read file",
			"message": err.Error(),
		})
	}

	// Create the file record in the database
	newFile, err := database.CreateFile(data, userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot create file record")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot create file record",
			"message": err.Error(),
		})
	}

	// Extract the extension from the original file name
	ext := filepath.Ext(fileUpload.Filename)

	// Save the file to storage
	newFileName := newFile.ID.String() + ext
	err = storageInstance.SaveFile(data, newFileName)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot save file")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot save file",
			"message": err.Error(),
		})
	}

	// Return the new file
	return c.JSON(newFile)
}

func getFile(c *fiber.Ctx) error {
	// Convert id to uuid
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid ID",
			"message": err.Error(),
		})
	}

	// Load the file from storage
	fileContent, err := storageInstance.LoadFile(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot load file from storage",
			"message": err.Error(),
		})
	}

	// Return the file
	return c.SendStream(fileContent)
}

func deleteFile(c *fiber.Ctx) error {
	// Convert id to uuid
	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid ID",
			"message": err.Error(),
		})
	}

	// Delete the file from storage
	err = storageInstance.DeleteFile(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot delete file from storage",
			"message": err.Error(),
		})
	}

	// Delete the file from the database
	err = database.DeleteFile(uuid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot delete file",
			"message": err.Error(),
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
