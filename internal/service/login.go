package service

import (
	"time"

	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRepo struct{}

func (l *LoginRepo) Login(c *fiber.Ctx, body *string, password *string) error {

	// Verify user credentials
	user, err := db.Users.VerifyUser(body, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidCredentials))
	}

	// Create claims
	expiryDate := time.Now().Add(time.Hour * 24 * 7)
	claims := jwt.RegisteredClaims{
		Subject:   user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(expiryDate),
	}

	// Create token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.Server.JWTSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return c.Status(fiber.StatusOK).JSON(response.CreateLoginResponse(user.ID, token, expiryDate))
}
