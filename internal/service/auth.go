package service

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthRepo struct{}

func (l *AuthRepo) Login(c *fiber.Ctx, body string, password string) error {
	logger.DebugLogRequestReceived("service", "login", "Login")

	// Verify user credentials
	user, err := db.Users.CheckCredentials(body, password)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error verifying user credentials")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidCredentials))
	}

	// Create claims
	expiryDate := time.Now().Add(time.Hour * 24 * 7)
	claims := jwt.RegisteredClaims{
		Issuer:    "frens",
		Subject:   user.ID.String(),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiryDate),
		Audience:  jwt.ClaimStrings{"frens"},
	}

	// Create token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.Server.JWTSecret))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateAuthResponse(token, user, expiryDate))
}
