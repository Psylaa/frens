package service

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/jsonapi"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Database    *database.Database
	JWTSecret   []byte
	JWTDuration int
}

type Token struct {
	ID string `jsonapi:"primary,token"`
}

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

	// Create the Claims
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(a.JWTDuration))),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encryptedToken, err := token.SignedString(a.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	// Prepare the response
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)
	c.Response().SetStatusCode(fiber.StatusOK)

	// Respond with the user
	return jsonapi.MarshalPayload(c.Response().BodyWriter(), &Token{ID: encryptedToken})
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

func (a *AuthService) Refresh(c *fiber.Ctx) error {
	userID, err := getRequestorID(c)
	if err != nil {
		return err
	}

	// Create the Claims
	newClaims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(a.JWTDuration)).Unix(),
	}

	// Create token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	encryptedToken, err := newToken.SignedString(a.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	// Prepare the response
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)
	c.Response().SetStatusCode(fiber.StatusOK)

	// Respond with the user
	return jsonapi.MarshalPayload(c.Response().BodyWriter(), &Token{ID: encryptedToken})
}
