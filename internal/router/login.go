package router

import (
	"strconv"
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
		logger.Log.Error().Err(err).Msg("Cannot parse JSON")
		return err
	}

	// Verify the username and password
	user, err := database.VerifyUser(body.Username, body.Password)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid username or password")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtDuration)).Unix()
	logger.Log.Info().Msg("Added claims to token. Claims: " + claims["user_id"].(uuid.UUID).String() + " " + strconv.FormatInt(claims["exp"].(int64), 10))

	// Sign the token with our secret
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Could not sign token")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the token
	logger.Log.Info().Msg("Successfully logged in")
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
		logger.Log.Error().Err(err).Msg("Cannot parse query")
		return err
	}

	// Verify the token
	token, err := jwt.Parse(body.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Extract the user ID from the token's claims
	claims := token.Claims.(jwt.MapClaims)
	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		logger.Log.Error().Msg("user_id claim is not a string")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id claim"})
	}

	// Convert the user ID string to a UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user_id UUID")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id UUID"})
	}

	// Return the user ID
	logger.Log.Info().Msg("Successfully verified token")
	return c.JSON(fiber.Map{"user_id": userID.String()})
}
