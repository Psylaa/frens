package router

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Init(port string, jwtSecret string) {
	// Initialize Fiber
	app := fiber.New()

	// Apply middleware
	app.Use(logger.New())

	// Define routes
	app.Post("/login", func(c *fiber.Ctx) error {
		return login(c, jwtSecret)
	})
	app.Post("/users", createUser)

	// Authenticated routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	}))

	app.Post("/statuses", func(c *fiber.Ctx) error {
		return createStatus(c, jwtSecret)
	})
	app.Get("/statuses/:id", getStatus)
	app.Delete("/statuses/:id", func(c *fiber.Ctx) error {
		return deleteStatus(c, jwtSecret)
	})

	// Start the server
	app.Listen(":" + port)
}

func getUserID(c *fiber.Ctx) (uuid.UUID, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("no sub claim in token")
	}
	return uuid.Parse(sub)
}
