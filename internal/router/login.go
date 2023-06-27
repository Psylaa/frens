package router

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func login(c *fiber.Ctx) error {
	// Parse the request body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	// Verify the username and password
	user, err := database.VerifyUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtDuration)).Unix()

	// Sign the token with our secret
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to sign token")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the token
	return c.JSON(fiber.Map{
		"id":    user.ID,
		"token": t})
}

func verifyToken(c *fiber.Ctx) error {
	// Parse the query parameter
	var body struct {
		Token string `query:"token"`
	}
	if err := c.QueryParser(&body); err != nil {
		return err
	}

	// Verify the token
	token, err := jwt.Parse(body.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Extract the user ID from the token's claims
	claims := token.Claims.(jwt.MapClaims)
	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id claim"})
	}

	// Convert the user ID string to a UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id UUID"})
	}

	// Return the user ID
	return c.JSON(fiber.Map{"user_id": userID.String()})
}
