package service

import (
	"log"

	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
)

type LikeService struct{ Database *database.Database }

func (ls *LikeService) LikePost(c *fiber.Ctx, postID uint32) error {
	// Get the ID of the user making the request
	userID, err := getRequestorID(c)
	if err != nil {
		// Log and handle error here
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user ID")
	}

	// Check if the post exists
	var post models.Post
	if err := ls.Database.Conn.First(&post, postID).Error; err != nil {
		// Post does not exist
		return c.Status(fiber.StatusNotFound).SendString("Post not found")
	}

	// Check if the user exists
	var user models.User
	if err := ls.Database.Conn.First(&user, userID).Error; err != nil {
		// User does not exist
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Create new Like
	newLike := models.Like{
		UserID: userID,
		PostID: postID,
	}

	// Save the like to the database
	if err := ls.Database.Conn.Create(&newLike).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create a like",
		})
	}

	// Set the content type to application/vnd.api+json
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)

	// Marshal the like into JSON API format
	if err := jsonapi.MarshalPayloadWithoutIncluded(c.Response().BodyWriter(), &newLike); err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal the like",
		})
	}

	// Set the status code to 201 Created
	c.Status(fiber.StatusCreated)

	return nil
}

func (ls *LikeService) UnlikePost(c *fiber.Ctx, postID uint32) error {
	// Get the ID of the user making the request
	userID, err := getRequestorID(c)
	if err != nil {
		// Log and handle error here
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user ID")
	}

	// Find the Like in the database
	var existingLike models.Like
	if err := ls.Database.Conn.Where("user_id = ? AND post_id = ?", userID, postID).First(&existingLike).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find the like",
		})
	}

	// Delete the like from the database
	if err := ls.Database.Conn.Delete(&existingLike).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to unlike a post",
		})
	}

	// Set the content type to application/vnd.api+json
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)

	// Marshal the like into JSON API format
	if err := jsonapi.MarshalPayloadWithoutIncluded(c.Response().BodyWriter(), &existingLike); err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal the like",
		})
	}

	// Set the status code to 200 OK
	c.Status(fiber.StatusOK)

	return nil
}
