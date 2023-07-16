package service

import (
	"log"

	"github.com/bwoff11/frens/models"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
)

type PostService struct{ Database *database.Database }

func (ps *PostService) Create(c *fiber.Ctx, text string, privacy string) error {
	// Get the ID of the user making the request
	userID, err := getRequestorID(c)
	if err != nil {
		// Log and handle error here
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user ID")
	}

	// Get the user from the database
	var user models.User
	if result := ps.Database.Conn.First(&user, userID); result.Error != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get the user from the database",
		})
	}

	// Assign the User to the newPost before saving it to the database
	newPost := models.Post{
		UserID:  userID,
		Text:    text,
		Privacy: privacy,
		User:    &user,
	}

	// Save the post to the database
	if err := ps.Database.Conn.Create(&newPost).Error; err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create a post",
		})
	}

	// Set the content type to application/vnd.api+json
	c.Response().Header.Set(fiber.HeaderContentType, jsonapi.MediaType)

	// Marshal the post into JSON API format
	if err := jsonapi.MarshalPayload(c.Response().BodyWriter(), &newPost); err != nil {
		// Log and handle error here
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal the post",
		})
	}

	// Set the status code to 201 Created
	c.Status(fiber.StatusCreated)

	return nil
}
