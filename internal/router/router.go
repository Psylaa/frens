package router

import (
	"fmt"
	"net/http"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var cfg *config.Config
var db *database.Database

func NewRouter(configuration *config.Config, database *database.Database) *Router {
	cfg = configuration
	db = database

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Server.AllowOrigins,
	}))

	r := &Router{cfg, app}
	r.setupRoutes()

	return r
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&APIResponse{
		Success: false,
		Error:   APIResponseErr(err.Error()),
	})
}

func (r *Router) setupRoutes() {
	// Unauthenticated routes
	r.App.Post("/login", login)
	r.App.Get("/login/verify", verifyToken)
	r.App.Get("/files/:filename", retrieveFile)

	// Middleware for JWT authentication
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(r.Config.Server.JWTSecret)},
	}))

	// Add authenticated routes
	r.AuthRoutes()
	r.ActivityPubRoutes()
}

func (r *Router) AuthRoutes() {
	// Users
	r.App.Get("/users", getUsers)
	r.App.Get("/users/:id", getUser)
	r.App.Post("/users", createUser)
	r.App.Patch("/users/:id", updateUser)

	// Users
	r.App.Get("/users", getUsers)
	r.App.Get("/users/:id", getUser)
	r.App.Post("/users", createUser)
	r.App.Patch("/users/:id", updateUser)

	// Statuses
	r.App.Get("/posts", getPosts)
	r.App.Post("/posts", createPost)
	r.App.Get("/posts/:id", getPost)
	r.App.Delete("/posts/:id", deletePost)

	// Likes
	r.App.Get("/statuses/:id/likes", getLikes)
	r.App.Post("/statuses/:id/likes", createLike)
	r.App.Delete("/statuses/:id/likes", deleteLike)
	r.App.Get("/statuses/:id/likes/:userId", hasUserLiked)

	// Files
	r.App.Post("/files", createFile)
	r.App.Delete("/files/:filename", deleteFile)

	// Bookmarks
	r.App.Get("/statuses/:id/bookmarks", getBookmarks)
	r.App.Post("/statuses/:id/bookmarks", createBookmark)
	r.App.Delete("/statuses/:id/bookmarks", deleteBookmark)
	r.App.Get("/statuses/:id/bookmarks/:userId", hasUserBookmarked)

	// Feed
	r.App.Get("/feed/chronological", getChronologicalFeed)
	r.App.Get("/feed/algorithmic", getChronologicalFeed)

	// Follows
	r.App.Get("/users/:id/followers", getFollowers)
	r.App.Post("/users/:id/followers", createFollower)
	r.App.Delete("/users/:id/followers", deleteFollower)
}

func (r *Router) ActivityPubRoutes() {
	//r.App.Get("/users/:username", activitypub.GetUserProfile)
	//r.App.Post("/users/:username/inbox", activitypub.HandleInbox)
	//r.App.Get("/users/:username/outbox", activitypub.HandleOutbox)
}

func (r *Router) Run() {
	if err := r.App.Listen(":" + r.Config.Server.Port); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to start server")
	}
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
