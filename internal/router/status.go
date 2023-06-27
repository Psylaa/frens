package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getStatus(c *fiber.Ctx) error {
	// Get the status ID from the URL parameter
	statusID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid status ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Get the status update
	status, err := database.GetStatus(statusID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to get status")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get status"})
	}

	// Return the status update
	logger.Log.Info().Msg("Successfully got status")
	return c.JSON(status)
}

func getStatuses(c *fiber.Ctx) error {
	// Get the user ID from the query parameters
	userID, err := uuid.Parse(c.Query("userId"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Get the statuses
	statuses, err := database.GetStatusesByUserID(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to get statuses")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get statuses"})
	}

	// Return the statuses
	logger.Log.Info().Msg("Successfully got statuses")
	return c.JSON(statuses)
}

func createStatus(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Parse the request body
	var body struct {
		Text  string           `json:"text"`
		Media []database.Media `json:"media"`
	}
	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Cannot parse JSON")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check if both text and media are empty
	if body.Text == "" && len(body.Media) == 0 {
		logger.Log.Error().Msg("At least one of 'text' or 'media' must be provided")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "At least one of 'text' or 'media' must be provided"})
	}

	// Create the status update
	status, err := database.CreateStatus(userID, body.Text, body.Media)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Could not create status")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not create status"})
	}

	// Return the created status update
	logger.Log.Info().Msg("Successfully created status")
	return c.Status(fiber.StatusCreated).JSON(status)
}

func deleteStatus(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Get the status ID from the URL parameter
	statusID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid status ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Get the status update
	status, err := database.GetStatus(statusID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to get status")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get status"})
	}

	// Check if the user is the creator or admin
	if status.User.ID != userID { // <-- Check the user's ID from the embedded User struct
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			logger.Log.Error().Msg("You do not have permission to delete this status")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You do not have permission to delete this status"})
		}
	}

	// Delete the status update
	if err := database.DeleteStatus(statusID); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to delete status")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete status"})
	}

	// Return success message
	logger.Log.Info().Msg("Successfully deleted status")
	return c.JSON(fiber.Map{"message": "Status successfully deleted"})
}
