package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UserRepo struct {
	Database *database.Database
}

func (ur *UserRepo) Create(c *fiber.Ctx, req *models.RegisterRequest) error {
	logger.Debug(logger.LogMessage{
		Package:  "service",
		Function: "UserRepo.Create",
		Message:  "Creating user: " + req.Username,
	})

	req.Sanitize()
	err := req.Validate()
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	newUser, err := req.ToUser()
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c)
	}

	// Create user in database
	if err := ur.Database.Create(&newUser).Error; err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	return newUser.ToResponse().Send(c)
}

func (ur *UserRepo) Login(c *fiber.Ctx, req *models.LoginRequest) error {
	logger.Debug(logger.LogMessage{
		Package:  "service",
		Function: "UserRepo.Login",
		Message:  "Logging in user: " + req.Email,
	})

	req.Sanitize()
	err := req.Validate()
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	// Find user in database
	user, err := ur.Database.Users.ReadByEmail(req.Email)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return models.ErrUnauthorized.SendResponse(c)
	}

	return nil
}
