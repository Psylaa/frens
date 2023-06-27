package router

import (
	"github.com/bwoff11/frens/internal/database"
	db "github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func getUsers(c *fiber.Ctx) error {
	users, err := db.GetUsers()
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting users")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully got users")
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := db.GetUser(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully got user")
	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	// Parse the request body
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Cannot parse JSON")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Create the user
	user, err := database.CreateUser(body.Username, body.Email, body.Password)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Could not create user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	logger.Log.Info().Msg("Successfully created user")
	return c.Status(fiber.StatusCreated).JSON(user)
}
