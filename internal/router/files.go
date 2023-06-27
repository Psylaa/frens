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
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	ext := filepath.Ext(file.Filename)

	fileData, err := database.CreateFile(userId, ext)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	if err := os.MkdirAll(cfg.Storage.Local.Path, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	if err := c.SaveFile(file, filepath.Join(cfg.Storage.Local.Path, fileData.ID.String()+ext)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data: []APIResponseData{
			APIResponseData{
				ID: fileData.ID,
				Attributes: APIResponseDataAttributes{
					Filename:  file.Filename,
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
				Success: false,
				Error:   "File not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
				Success: false,
				Error:   ErrInternal,
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
			Success: false,
			Error:   ErrInternal,
		})
	}
	return c.JSON(APIResponse{
		Success: true,
	})
}
