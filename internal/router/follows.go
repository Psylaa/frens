package router

import (
	"github.com/gofiber/fiber/v2"
)

func getFollows(c *fiber.Ctx) error {
	/*
		id := c.Params("id")
		userID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		follows, err := db.Follows.GetFollows(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.JSON(response.GenerateFollowsResponse(follows))
	*/
	return nil
}

func createFollow(c *fiber.Ctx) error {
	/*
		sourceID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		}

		id := c.Params("id")
		targetID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		// Check if the follower record already exists
		exists, err := db.Follows.DoesFollowExist(sourceID, targetID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		if exists {
			return c.Status(fiber.StatusConflict).JSON(response.CreateErrorResponse(response.ErrExists))
		}

		follow, err := db.Follows.CreateFollow(sourceID, targetID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusCreated).JSON(response.GenerateFollowResponse(follow))
	*/
	return nil
}

func deleteFollow(c *fiber.Ctx) error {
	/*
		SourceID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		}

		id := c.Params("id")
		TargetID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		// Check if the follower record exists
		exists, err := db.Follows.DoesFollowExist(SourceID, TargetID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
		}

		if err := db.Follows.DeleteFollow(SourceID, TargetID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.SendStatus(fiber.StatusOK)
	*/
	return nil
}

func getFollowing(c *fiber.Ctx) error {
	/*
		id := c.Params("id")
		userID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		following, err := db.Follows.GetFollowing(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.JSON(response.GenerateFollowsResponse(following))
	*/
	return nil
}
