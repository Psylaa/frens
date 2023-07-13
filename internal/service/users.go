package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepo struct{}

func (ur *UserRepo) GetByID(c *fiber.Ctx, userID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (ur *UserRepo) Create(c *fiber.Ctx, username string, email string, password string) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (ur *UserRepo) Update(c *fiber.Ctx, bio *string, avatarID *uuid.UUID, coverID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (ur *UserRepo) Delete(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
