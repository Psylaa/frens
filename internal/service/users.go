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

	// Sanitize and validate request
	req.Sanitize()
	err := req.Validate()
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	// Convert request to user
	newUser, err := req.ToUser()
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c)
	}

	// Add default bio
	newUser.SetBio(defaultBio)

	// Create user in database
	if err := ur.Database.Create(&newUser).Error; err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Convert to response
	resp := newUser.ToResponse()

	// Add token
	err = resp.AddToken(JWTSigningKey, JWTDuration)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Send response
	return resp.Send(c)
}

func (ur *UserRepo) Login(c *fiber.Ctx, req *models.LoginRequest) error {
	logger.Debug(logger.LogMessage{
		Package:  "service",
		Function: "UserRepo.Login",
		Message:  "Logging in user: " + req.Email,
	})

	// Sanitize and validate request
	req.Sanitize()
	err := req.Validate()
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	// Find user in database
	user, err := ur.Database.Users.ReadByEmail(req.Email)
	if err != nil {
		return models.ErrUnauthorized.SendResponse(c, "No user found with that email")
	}

	// Check password
	if !user.IsPasswordValid(req.Password) {
		return models.ErrUnauthorized.SendResponse(c, "invalid password")
	}

	// Convert to response
	resp := user.ToResponse()

	// Add token
	err = resp.AddToken(JWTSigningKey, JWTDuration)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Send response
	return resp.Send(c)
}
