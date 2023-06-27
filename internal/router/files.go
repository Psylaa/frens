package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func createFile(c *fiber.Ctx) error {
	// Parse the request body into a File object
	var file database.File
	if err := c.BodyParser(&file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Cannot parse JSON",
			"message": err.Error(),
		})
	}

	// Call the CreateFile function from the database package
	newFile, err := database.CreateFile(&file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot create file",
			"message": err.Error(),
		})
	}

	// Save the file to storage
	err = storage.SaveFile(storage.FileType(newFile.Type), c.Params("path"), c.Body())
	if err != nil {
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

	// Call the GetFile function from the database package
	file, err := database.GetFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot get file",
			"message": err.Error(),
		})
	}

	// Load the file from storage
	fileContent, err := storage.LoadFile(storage.FileType(file.Type), c.Params("path"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Cannot load file",
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
	err = storage.DeleteFile(storage.FileType(file.Type), c.Params("path"))
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
