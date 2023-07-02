package router

import (
	"github.com/gofiber/fiber/v2"
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
	return nil
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

func createLike(c *fiber.Ctx) error {
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

		// Check if user has already liked this status
		userHasLiked, err := db.Likes.HasUserLiked(userID, postID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if userHasLiked {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User has already liked this status"})
		}

		// Create the like
		if _, err := db.Likes.CreateLike(userID, postID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(fiber.StatusOK)
	*/
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
