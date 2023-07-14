package service

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookmarkRepo struct {
	Database *database.Database
}

func (br *BookmarkRepo) Get(c *fiber.Ctx, count int, cursor time.Time) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (br *BookmarkRepo) Create(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (br *BookmarkRepo) Delete(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (br *BookmarkRepo) GetByID(c *fiber.Ctx, id *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (br *BookmarkRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID, count, offset *int) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (br *BookmarkRepo) DeleteByPostID(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
