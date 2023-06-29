package router

import (
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getFollows(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	follows, err := db.Follows.GetFollows(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.JSON(response.GenerateFollowsResponse(follows))
}

func createFollow(c *fiber.Ctx) error {
	sourceID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
	}

	id := c.Params("id")
	targetID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	// Check if the follower record already exists
	exists, err := db.Follows.DoesFollowExist(sourceID, targetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(response.GenerateErrorResponse(response.ErrExists))
	}

	follow, err := db.Follows.CreateFollow(sourceID, targetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.Status(fiber.StatusCreated).JSON(response.GenerateFollowResponse(follow))
}

func deleteFollow(c *fiber.Ctx) error {
	SourceID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
	}

	id := c.Params("id")
	TargetID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	// Check if the follower record exists
	exists, err := db.Follows.DoesFollowExist(SourceID, TargetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(response.GenerateErrorResponse(response.ErrNotFound))
	}

	if err := db.Follows.DeleteFollow(SourceID, TargetID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.SendStatus(fiber.StatusOK)
}

func getFollowing(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidID))
	}

	following, err := db.Follows.GetFollowing(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GenerateErrorResponse(response.ErrInternal))
	}

	return c.JSON(response.GenerateFollowsResponse(following))
}
