package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getFollows(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	followers, err := db.Follows.GetFollows(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting followers")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	var data []APIResponseData
	for _, follower := range followers {
		data = append(data, createFollowAPIResponseData(&follower))
	}

	logger.Log.Debug().Msg("Fetched followers successfully")

	return c.JSON(APIResponse{
		Data: data,
	})
}

func createFollow(c *fiber.Ctx) error {
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
		logger.Log.Warn().Msg("Follow already exists")
		return c.Status(fiber.StatusConflict).JSON(APIResponse{
			Error: "Follow already exists",
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

func deleteFollow(c *fiber.Ctx) error {
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
		logger.Log.Warn().Msg("Follow does not exist")
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Error: "Follow does not exist",
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

func getFollowing(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	following, err := db.Follows.GetFollowing(userID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting following")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	var data []APIResponseData
	for _, follow := range following {
		data = append(data, createFollowAPIResponseData(&follow))
	}

	logger.Log.Debug().Msg("Fetched following successfully")

	return c.JSON(APIResponse{
		Data: data,
	})
}

func createFollowAPIResponseData(follow *database.Follow) APIResponseData {
	return APIResponseData{
		Type: shared.DataTypeFollow,
		ID:   &follow.ID,
		Attributes: APIResponseDataAttributes{
			CreatedAt: &follow.CreatedAt,
			UpdatedAt: &follow.UpdatedAt,
			SourceID:  &follow.SourceID,
			TargetID:  &follow.TargetID,
		},
		Links: APIResponseDataLinks{
			Self:   "/follows/" + follow.ID.String(),
			Source: "/users/" + follow.SourceID.String(),
			Target: "/users/" + follow.TargetID.String(),
		},
	}
}
