package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepo struct {
	Database *database.Database
}

func (ur *UserRepo) GetByID(c *fiber.Ctx, userID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (ur *UserRepo) Create(c *fiber.Ctx, req models.RegisterRequest) error {
	logger.DebugLogRequestReceived("service", "user", "create")

	// Convert request to user model
	newUser, err := req.ToUser()
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c)
	}

	// Create user in database
	if err := ur.Database.Create(&newUser).Error; err != nil {
		return models.ErrInternalServerError.SendResponse(c, "Failed to insert user into database")
	}

	// Convert user to response
	return nil
}

func (ur *UserRepo) Update(c *fiber.Ctx, bio *string, avatarID *uuid.UUID, coverID *uuid.UUID) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (ur *UserRepo) Delete(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
