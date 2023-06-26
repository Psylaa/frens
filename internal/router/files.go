package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func createFile(c *fiber.Ctx) error {
	// Parse the request body into a File object
	var file database.File
	if err := c.BodyParser(&file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Call the CreateFile function from the database package
	newFile, err := database.CreateFile(&file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create file",
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
			"error": "Invalid ID",
		})
	}

	// Call the GetFile function from the database package
	file, err := database.GetFile(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get file",
		})
	}

	// Return the file
	return c.JSON(file)
}

func updateFile(c *fiber.Ctx) error {
	// Get the file ID from the URL parameter
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	// Parse the request body into a File object
	var file database.File
	if err := c.BodyParser(&file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Ensure the file ID matches the ID in the URL
	if file.ID != id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID mismatch",
		})
	}

	// Call the UpdateFile function from the database package
	if err := database.UpdateFile(&file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update file",
		})
	}

	// Return the updated file
	return c.JSON(file)
}

func deleteFile(c *fiber.Ctx) error {
	// Get the file ID from the URL parameter
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	// Call the DeleteFile function from the database package
	if err := database.DeleteFile(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete file",
		})
	}

	// Return a success message
	return c.JSON(fiber.Map{
		"message": "File deleted",
	})
}
