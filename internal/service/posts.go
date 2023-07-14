package service

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostRepo struct {
	Database *database.Database
}

func (pr *PostRepo) Create(c *fiber.Ctx, text string, privacy shared.Privacy, mediaIDs []*uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (pr *PostRepo) GetByID(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (ur *PostRepo) GetByUserID(c *fiber.Ctx, userID *uuid.UUID, cursor time.Time, count int) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (pr *PostRepo) GetReplies(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (pr *PostRepo) Update(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (pr *PostRepo) Delete(c *fiber.Ctx, postID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
