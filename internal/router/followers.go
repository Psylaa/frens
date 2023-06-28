package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getFollowers(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	followers, err := db.Followers.GetFollowers(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting followers")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	var data []APIResponseData
	for _, follower := range followers {
		data = append(data, createFollowerAPIResponseData(&follower))
	}

	logger.Log.Debug().Msg("Fetched followers successfully")

	return c.JSON(APIResponse{
		Success: true,
		Data:    data,
	})
}

func createFollower(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Success: false,
			Error:   ErrUnauthorized,
		})
	}

	id := c.Params("id")
	followingID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	// Check if the follower record already exists
	exists, err := db.Followers.FollowerExists(userID, followingID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error checking follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	if exists {
		logger.Log.Warn().Msg("Follower already exists")
		return c.Status(fiber.StatusConflict).JSON(APIResponse{
			Success: false,
			Error:   "Follower already exists",
		})
	}

	if _, err := db.Followers.CreateFollower(userID, followingID); err != nil {
		logger.Log.Error().Err(err).Msg("Error creating follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	logger.Log.Debug().Msg("Created follower successfully")

	return c.SendStatus(fiber.StatusOK)
}

func deleteFollower(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidToken,
		})
	}

	id := c.Params("id")
	followingID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	// Check if the follower record exists
	exists, err := db.Followers.FollowerExists(userID, followingID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error checking follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	if !exists {
		logger.Log.Warn().Msg("Follower does not exist")
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Success: false,
			Error:   "Follower does not exist",
		})
	}

	if err := db.Followers.DeleteFollower(userID, followingID); err != nil {
		logger.Log.Error().Err(err).Msg("Error deleting follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	logger.Log.Debug().Msg("Deleted follower successfully")

	return c.SendStatus(fiber.StatusOK)
}

func createFollowerAPIResponseData(follower *database.Follower) APIResponseData {
	return APIResponseData{
		Type:       shared.DataTypeFollower,
		ID:         follower.ID,
		Attributes: APIResponseDataAttributes{
			// You should add appropriate attributes here, I've left it blank since the structure of your Follower wasn't provided
		},
		Relationships: APIResponseDataRelationships{
			OwnerID: follower.FollowerID,
		},
	}
}
