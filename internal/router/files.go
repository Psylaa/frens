package router

import (
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

func createFile(c *fiber.Ctx) error {
	userId, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ext := filepath.Ext(file.Filename)

	fileData, err := database.CreateFile(userId, ext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(cfg.Storage.Local.Path, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.SaveFile(file, filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+ext)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fileData)
}

func retrieveFile(c *fiber.Ctx) error {
	filePath := filepath.Join(cfg.Storage.Local.Path, c.Params("filename"))
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
		} else {
			// other error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}
	return c.SendFile(filePath)
}

func deleteFile(c *fiber.Ctx) error {
	filePath := filepath.Join(cfg.Storage.Local.Path, c.Params("filename"))
	err := os.Remove(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "File deleted"})
}

/*
func updateFile(c *fiber.Ctx) error {
	oldFileName := c.Params("filename")
	userId, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ext := filepath.Ext(file.Filename)

	fileData, err := database.UpdateFile(userId, oldFileName, ext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	oldFilePath := filepath.Join(cfg.Storage.Local.Path, oldFileName)
	err = os.Remove(oldFilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	newFilePath := filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+ext)
	out, err := os.Create(newFilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer out.Close()

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer fileContent.Close()

	_, err = io.Copy(out, fileContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fileData)
}
*/
