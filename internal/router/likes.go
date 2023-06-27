package router

import (
	db "github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getLikes(c *fiber.Ctx) error {
	id := c.Params("id")
	statusID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid status ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	likes, err := db.GetLikes(statusID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot get likes")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully retrieved likes")
	return c.JSON(likes)
}

func createLike(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	id := c.Params("id")
	statusID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid status ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	if _, err := db.CreateLike(userID, statusID); err != nil {
		logger.Log.Error().Err(err).Msg("Cannot create like")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteLike(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	id := c.Params("id")
	statusID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid status ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	if err := db.DeleteLike(userID, statusID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully deleted like")
	return c.SendStatus(fiber.StatusOK)
}

func hasUserLiked(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	id := c.Params("id")
	statusId, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid status ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	liked, err := db.HasUserLiked(userId, statusId)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot get like")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully retrieved like")
	return c.JSON(fiber.Map{"liked": liked})
}
