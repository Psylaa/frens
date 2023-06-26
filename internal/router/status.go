package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getStatus(c *fiber.Ctx) error {
	// Get the status ID from the URL parameter
	statusID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Get the status update
	status, err := database.GetStatus(statusID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get status"})
	}

	// Return the status update
	return c.JSON(status)
}

func createStatus(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Parse the request body
	var body struct {
		Text  string           `json:"text"`
		Media []database.Media `json:"media"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check if both text and media are empty
	if body.Text == "" && len(body.Media) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "At least one of 'text' or 'media' must be provided"})
	}

	// Create the status update
	status, err := database.CreateStatus(userID, body.Text, body.Media)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not create status"})
	}

	// Return the created status update
	return c.Status(fiber.StatusCreated).JSON(status)
}

func deleteStatus(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Get the status ID from the URL parameter
	statusID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Get the status update
	status, err := database.GetStatus(statusID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get status"})
	}

	// Check if the user is the creator or admin
	if status.User.ID != userID { // <-- Check the user's ID from the embedded User struct
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You do not have permission to delete this status"})
		}
	}

	// Delete the status update
	if err := database.DeleteStatus(statusID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete status"})
	}

	// Return success message
	return c.JSON(fiber.Map{"message": "Status successfully deleted"})
}
