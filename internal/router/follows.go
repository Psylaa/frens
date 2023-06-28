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
			Error: ErrInvalidID,
		})
	}

	followers, err := db.Follows.GetFollowers(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting followers")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	var data []APIResponseData
	for _, follower := range followers {
		data = append(data, createFollowerAPIResponseData(&follower))
	}

	logger.Log.Debug().Msg("Fetched followers successfully")

	return c.JSON(APIResponse{
		Data: data,
	})
}

func createFollower(c *fiber.Ctx) error {
	sourceID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrUnauthorized,
		})
	}

	id := c.Params("id")
	targetID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	// Check if the follower record already exists
	exists, err := db.Follows.DoesFollowExist(sourceID, targetID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error checking follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	if exists {
		logger.Log.Warn().Msg("Follower already exists")
		return c.Status(fiber.StatusConflict).JSON(APIResponse{
			Error: "Follower already exists",
		})
	}

	if _, err := db.Follows.CreateFollow(sourceID, targetID); err != nil {
		logger.Log.Error().Err(err).Msg("Error creating follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	logger.Log.Debug().Msg("Created follower successfully")

	return c.SendStatus(fiber.StatusOK)
}

func deleteFollower(c *fiber.Ctx) error {
	SourceID, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID in token")
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInvalidToken,
		})
	}

	id := c.Params("id")
	TargetID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	// Check if the follower record exists
	exists, err := db.Follows.DoesFollowExist(SourceID, TargetID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error checking follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	if !exists {
		logger.Log.Warn().Msg("Follower does not exist")
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Error: "Follower does not exist",
		})
	}

	if err := db.Follows.DeleteFollow(SourceID, TargetID); err != nil {
		logger.Log.Error().Err(err).Msg("Error deleting follower")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	logger.Log.Debug().Msg("Deleted follower successfully")

	return c.SendStatus(fiber.StatusOK)
}

func createFollowerAPIResponseData(follower *database.Follow) APIResponseData {
	return APIResponseData{
		Type: shared.DataTypeFollower,
		ID:   &follower.ID,
		Attributes: APIResponseDataAttributes{
			CreatedAt: &follower.CreatedAt,
			UpdatedAt: &follower.UpdatedAt,
			SourceID:  &follower.SourceID,
			TargetID:  &follower.TargetID,
		},
		Links: APIResponseDataLinks{
			Self:   "/followers/" + follower.ID.String(),
			Source: "/users/" + follower.ID.String(),
			Target: "/users/" + follower.TargetID.String(),
		},
	}
}
