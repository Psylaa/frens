package router

import (
	db "github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getFollowers(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	followers, err := db.GetFollowers(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Cannot get followers")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully retrieved followers")
	return c.JSON(followers)
}

func createFollower(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	id := c.Params("id")
	followingID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if _, err := db.CreateFollower(userID, followingID); err != nil {
		logger.Log.Error().Err(err).Msg("Cannot create follower")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully created follower")
	return c.SendStatus(fiber.StatusOK)
}

func deleteFollower(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	id := c.Params("id")
	followingID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := db.DeleteFollower(userID, followingID); err != nil {
		logger.Log.Error().Err(err).Msg("Cannot delete follower")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Log.Info().Msg("Successfully deleted follower")
	return c.SendStatus(fiber.StatusOK)
}
