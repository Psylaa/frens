package router

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(port string, jwtSecret string) {
	// Initialize Fiber
	app := fiber.New()

	// Apply middleware
	app.Use(logger.New())

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return login(c, jwtSecret)
	})
	app.Post("/users", createUser)

	// Start the server
	app.Listen(":" + port)
}

func login(c *fiber.Ctx, jwtSecret string) error {
	// Parse the request body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	// Verify the username and password
	user, err := database.VerifyUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires after 72 hours

	// Sign the token with our secret
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the token
	return c.JSON(fiber.Map{"token": t})
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
