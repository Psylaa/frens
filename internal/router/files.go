package router

import (
	"os"
	"path/filepath"

	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
)

func createFile(c *fiber.Ctx) error {
	userId, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return srv.Files.Create(c, &userId, file)
}

func retrieveFile(c *fiber.Ctx) error {
	filePath := filepath.Join(cfg.Storage.Local.Path, c.Params("filename"))
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
	}
	return c.SendFile(filePath)
}

func deleteFile(c *fiber.Ctx) error {
	filePath := filepath.Join(cfg.Storage.Local.Path, c.Params("filename"))
	err := os.Remove(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	return c.SendStatus(fiber.StatusOK)
}
