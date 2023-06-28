package router

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func createFile(c *fiber.Ctx) error {
	userId, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	ext := filepath.Ext(file.Filename)

	fileData, err := db.Files.CreateFile(userId, ext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	if err := os.MkdirAll(cfg.Storage.Local.Path, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	if err := c.SaveFile(file, filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+ext)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Data: []APIResponseData{
			{
				ID: &fileData.ID,
				Attributes: APIResponseDataAttributes{
					Extension: ext,
				},
			},
		},
	})
}

func retrieveFile(c *fiber.Ctx) error {
	filePath := filepath.Join(cfg.Storage.Local.Path, c.Params("filename"))
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).JSON(APIResponse{
				Error: "File not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
				Error: ErrInternal,
			})
		}
	}
	return c.SendFile(filePath)
}

func deleteFile(c *fiber.Ctx) error {
	filePath := filepath.Join(cfg.Storage.Local.Path, c.Params("filename"))
	err := os.Remove(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	return c.JSON(APIResponse{})
}
