package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LikeRepo struct{ Database *database.Database }

func (lr *LikeRepo) GetByID(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (lr *LikeRepo) GetByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (lr *LikeRepo) GetByPostIDAndUserID(c *fiber.Ctx, postID *uuid.UUID, userID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (lr *LikeRepo) Create(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (lr *LikeRepo) Delete(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
