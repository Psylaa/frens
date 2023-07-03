package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
)

func getLikeByID(c *fiber.Ctx) error {
	return nil /*
		id := c.Params("id")
		postID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
		}

		likes, err := db.Likes.GetLikes(postID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(likes)
	*/
}

func deleteLikeByID(c *fiber.Ctx) error {
	return nil
}

func getLikesByPostID(c *fiber.Ctx) error {
	return nil
}

func getLikesCountByPostID(c *fiber.Ctx) error {
	return nil
}

func createLikeByPostID(c *fiber.Ctx) error {

	// Get post ID from the params
	postID := c.Params("postId")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
	}

	// Convert the post ID to a UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing post ID to UUID")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
	}

	// Send the request to the service
	return srv.Likes.Create(c, &postUUID)
}

func deleteLikeByPostID(c *fiber.Ctx) error {
	return nil
}

func getLikesByUserID(c *fiber.Ctx) error {
	return nil
}

func getLikesCountByUserID(c *fiber.Ctx) error {
	return nil
}

func deleteLike(c *fiber.Ctx) error {
	return nil
	/*
		userID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
		}

		id := c.Params("id")
		postID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
		}

		if err := db.Likes.DeleteLike(userID, postID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(fiber.StatusOK)
	*/
}

func hasUserLiked(c *fiber.Ctx) error {
	return nil
	/*
		userId, err := uuid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		id := c.Params("id")
		postId, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
		}

		liked, err := db.Likes.HasUserLiked(userId, postId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"liked": liked})
	*/
}
