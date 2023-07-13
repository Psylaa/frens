package service

import (
	"github.com/gofiber/fiber/v2"
)

type AuthRepo struct{}

func (l *AuthRepo) Login(c *fiber.Ctx, body string, password string) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
