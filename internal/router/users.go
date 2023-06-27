package router

import (
	"log"

	"github.com/bwoff11/frens/internal/database"
	db "github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getUsers(c *fiber.Ctx) error {
	users, err := db.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := db.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	// Parse the request body
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Create the user
	user, err := database.CreateUser(body.Username, body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	log.Println("Raw Request Body:", string(c.Body())) // Add this log
	type request struct {
		Bio            *string    `json:"bio"`
		ProfilePicture *uuid.UUID `json:"profilePicture"`
		BannerImage    *uuid.UUID `json:"bannerImage"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	userId := c.Params("id")
	user, err := db.GetUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUser, err := db.UpdateUser(user.ID, body.Bio, body.ProfilePicture, body.BannerImage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedUser)
}
