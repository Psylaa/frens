package router

import (
	"io"

	"github.com/bwoff11/frens/internal/config"
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
	fileType := c.FormValue("type")

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

	// Create a new db file instance
	fileData := database.File{
		Type:  fileType,
		Owner: userID,
	}

	// Call the CreateFile function from the database package
	newFile, err := database.CreateFile(&fileData)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot create file")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot create file",
			"message": err.Error(),
		})
	}

	// Save the file to storage
	err = storages[config.FileType(fileType)].SaveFile(newFile.ID, data)
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
	// Get the file ID from the URL parameter
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid ID",
			"message": err.Error(),
		})
	}

	// Get the file type from the URL parameter
	// Todo: verify its a valid type
	fileType := c.Params("type")
	if fileType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid type",
			"message": "Type cannot be empty",
		})
	}

	// Load the file from storage
	fileContent, err := storages[config.FileType(fileType)].LoadFile(id)
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
	// Get the file ID from the URL parameter
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid ID",
			"message": err.Error(),
		})
	}

	// Call the GetFile function from the database package
	file, err := database.GetFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot get file",
			"message": err.Error(),
		})
	}

	// Call the DeleteFile function from the database package
	if err := database.DeleteFile(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot delete file",
			"message": err.Error(),
		})
	}

	// Delete the file from storage
	err = storages[config.FileType(file.Type)].DeleteFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot delete file from storage",
			"message": err.Error(),
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
