package router

import (
	"fmt"

	"github.com/bwoff11/frens/internal/activitypub"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // change to your front-end origin
	}))

	// Define routes
	app.Post("/login", login)
	app.Get("/login/verify", verifyToken)

	// Authenticated routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	}))

	// Users
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)
	app.Post("/users", createUser)
	app.Patch("/users/:id", updateUser)

	// Statuses
	app.Get("/statuses", getStatuses)
	app.Post("/statuses", createStatus)
	app.Get("/statuses/:id", getStatus)
	app.Delete("/statuses/:id", deleteStatus)

	// Likes
	app.Get("/statuses/:id/likes", getLikes)
	app.Post("/statuses/:id/likes", createLike)
	app.Delete("/statuses/:id/likes", deleteLike)
	app.Get("/statuses/:id/likes/:userId", hasUserLiked)

	// Files
	app.Post("/files", createFile)
	app.Get("/files/:id", getFile)
	app.Delete("/files/:id", deleteFile)

	// Bookmarks
	app.Get("/statuses/:id/bookmarks", getBookmarks)
	app.Post("/statuses/:id/bookmarks", createBookmark)
	app.Delete("/statuses/:id/bookmarks", deleteBookmark)
	app.Get("/statuses/:id/bookmarks/:userId", hasUserBookmarked)

	// Feed
	app.Get("/feed/chronological", getChronologicalFeed)
	app.Get("/feed/algorithmic", getChronologicalFeed)

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
