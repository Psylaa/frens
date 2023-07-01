package router

import (
	"time"

	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// login handles the HTTP request for user login.
func login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	user, err := db.Users.VerifyUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
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

	return c.Status(fiber.StatusOK).JSON(response.GenerateLoginResponse(token, expiryDate, user.ID))
}

// verifyToken handles the HTTP request for token verification.
func verifyUserToken(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
