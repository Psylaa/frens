package service

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FeedRepo struct {
	Database *database.Database
}

func (f *FeedRepo) GetChrono(c *fiber.Ctx, userID *uuid.UUID, cursor time.Time) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
