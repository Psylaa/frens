package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LikeRepo struct{}

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
