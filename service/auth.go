package service

import (
	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{ Database *database.Database }

func (a *AuthService) Login(c *fiber.Ctx, email, password string) error {
	var user models.User

	// Check if user exists
	if err := a.Database.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Check if password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (a *AuthService) Register(c *fiber.Ctx, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Respond with error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	newUser := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := a.Database.Conn.Create(&newUser).Error; err != nil {
		// Respond with error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Prepare the response
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)
	c.Response().SetStatusCode(fiber.StatusCreated)

	// Respond with the created user
	return jsonapi.MarshalPayload(c.Response().BodyWriter(), &newUser)
}
