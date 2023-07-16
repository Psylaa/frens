package service

import (
	"log"

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

	// Validate request
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
	newUser.Bio = defaultBio

	// Create user in database
	if err := ur.Database.Users.Create(newUser); err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Convert to response data
	respData := newUser.ToResponseData()

	// Create response
	resp := models.CreateUserResponse(respData)

	// Add token
	err = resp.AddToken(JWTSigningKey, JWTDuration)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (ur *UserRepo) Login(c *fiber.Ctx, req *models.LoginRequest) error {
	logger.Debug(logger.LogMessage{
		Package:  "service",
		Function: "UserRepo.Login",
		Message:  "Logging in user: " + req.Email,
	})

	// Validate request
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
	log.Println(req.Password)
	if err := user.CheckPassword(req.Password); err != nil {
		logger.Info(logger.LogMessage{
			Package:  "service",
			Function: "UserRepo.Login",
			Message:  "Checkpassword failed for user: " + user.Username + " with error: " + err.Error(),
		})
		return models.ErrUnauthorized.SendResponse(c, "invalid password")
	}

	// Convert to response data
	respData := user.ToResponseData()

	// Create response
	resp := models.CreateUserResponse(respData)

	// Add token
	err = resp.AddToken(JWTSigningKey, JWTDuration)
	if err != nil {
		return models.ErrInternalServerError.SendResponse(c, err.Error())
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(resp)
}
