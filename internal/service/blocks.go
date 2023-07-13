package service

import (
	"github.com/gofiber/fiber/v2"
)

type BlockRepo struct{}

func (br *BlockRepo) Get(c *fiber.Ctx, count int, offset int) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
