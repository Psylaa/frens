package router

import (
	"github.com/bwoff11/frens/internal/database"
	db "github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

func getUsers(c *fiber.Ctx) error {
	users, err := db.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
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
