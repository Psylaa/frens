package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

type BlockRepo struct {
	Database *database.Database
}

func (br *BlockRepo) Get(c *fiber.Ctx, count int, offset int) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
