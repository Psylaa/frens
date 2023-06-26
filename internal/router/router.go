package router

import (
	"fmt"

	"github.com/bwoff11/frens/internal/activitypub"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	jwtSecret   string
	jwtDuration int
)

func Init(port string, secret string, duration int) {
	// Set global variables
	jwtSecret = secret
	jwtDuration = duration

	// Initialize Fiber
	app := fiber.New()

	// Apply middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // change to your front-end origin
	}))

	// Define routes
	app.Post("/login", login)
	app.Get("/login/verify", verifyToken)
	app.Post("/users", createUser)

	// Authenticated routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	}))

	// Other
	app.Get("/users", getUsers)

	// Users
	app.Post("/statuses", createStatus)
	app.Get("/statuses/:id", getStatus)
	app.Delete("/statuses/:id", deleteStatus)

	// Likes
	app.Get("/statuses/:id/likes", getLikes)
	app.Post("/statuses/:id/likes", createLike)
	app.Delete("/statuses/:id/likes", deleteLike)

	// Feed
	app.Get("/feed/chronological", getChronologicalFeed)
	app.Get("/feed/algorithmic", getChronologicalFeed) //placeholder

	// Follows
	app.Get("/users/:id/followers", getFollowers)
	app.Post("/users/:id/followers", createFollower)
	app.Delete("/users/:id/followers", deleteFollower)

	// ActivityPub routes
	app.Get("/users/:username", activitypub.GetUserProfile)
	app.Post("/users/:username/inbox", activitypub.HandleInbox)
	app.Get("/users/:username/outbox", activitypub.HandleOutbox)

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
