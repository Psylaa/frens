package service

import (
	"log"

	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
)

type BookmarkService struct{ Database *database.Database }

func (bs *BookmarkService) BookmarkPost(c *fiber.Ctx, postID uint32) error {
	// Get the ID of the user making the request
	userID, err := getRequestorID(c)
	if err != nil {
		// Log and handle error here
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user ID")
	}

	// Check if the post exists
	var post models.Post
	if err := bs.Database.Conn.First(&post, postID).Error; err != nil {
		// Post does not exist
		return c.Status(fiber.StatusNotFound).SendString("Post not found")
	}

	// Check if the user exists
	var user models.User
	if err := bs.Database.Conn.First(&user, userID).Error; err != nil {
		// User does not exist
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Create new Bookmark
	newBookmark := models.Bookmark{
		UserID: userID,
		PostID: postID,
	}

	// Save the bookmark to the database
	if err := bs.Database.Conn.Create(&newBookmark).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create a bookmark",
		})
	}

	// Preload User and Post on the newly created Bookmark
	if err := bs.Database.Conn.
		Preload("User").
		Preload("Post").
		Preload("Post.User").
		First(&newBookmark, newBookmark.ID).
		Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve the bookmark",
		})
	}

	// Set the content type to application/vnd.api+json
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)

	// Marshal the bookmark into JSON API format
	if err := jsonapi.MarshalPayload(c.Response().BodyWriter(), &newBookmark); err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal the bookmark",
		})
	}

	// Set the status code to 201 Created
	c.Status(fiber.StatusCreated)

	return nil
}

func (bs *BookmarkService) UnbookmarkPost(c *fiber.Ctx, postID uint32) error {
	// Get the ID of the user making the request
	userID, err := getRequestorID(c)
	if err != nil {
		// Log and handle error here
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user ID")
	}

	// Find the Bookmark in the database
	var existingBookmark models.Bookmark
	if err := bs.Database.Conn.
		Preload("User").
		Preload("Post").
		Preload("Post.User").
		Where("user_id = ? AND post_id = ?", userID, postID).First(&existingBookmark).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find the bookmark",
		})
	}

	// Set the content type to application/vnd.api+json
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)

	// Marshal the bookmark into JSON API format
	if err := jsonapi.MarshalPayload(c.Response().BodyWriter(), &existingBookmark); err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal the bookmark",
		})
	}

	// Delete the bookmark from the database
	if err := bs.Database.Conn.Delete(&existingBookmark).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to unbookmark a post",
		})
	}

	// Set the status code to 200 OK
	c.Status(fiber.StatusOK)

	return nil
}
