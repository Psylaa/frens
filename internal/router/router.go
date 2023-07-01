package router

import (
	"fmt"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/contrib/fiberzerolog"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	_ "github.com/bwoff11/frens/docs" // For swagger docs
)

var cfg *config.Config
var db *database.Database

// Structure representing the router
type Router struct {
	Config *config.Config
	App    *fiber.App
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func NewRouter(configuration *config.Config, database *database.Database) *Router {
	cfg = configuration
	db = database

	app := fiber.New(fiber.Config{})

	r := &Router{cfg, app}
	r.setupMiddleware()
	r.setupRoutes()

	return r
}

func (r *Router) setupRoutes() {
	r.addUnauthRoutes()

	// Middleware for JWT authentication
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.Server.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
			}
			return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
		},
	}))

	// Add authenticated routes
	r.AuthRoutes()
	r.ActivityPubRoutes()
}

func (r *Router) setupMiddleware() {
	r.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))
	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}

func (r *Router) addUnauthRoutes() {
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	r.App.Post("/login", login)
	r.App.Get("/files/:filename", retrieveFile)
	r.App.Get("/feeds/explore", getExploreFeed)
	r.App.Get("/users/:id", getUser)
	r.App.Get("/posts", GetPostsByUserID)
	r.App.Get("/posts/:id", getPost)

	logger.Log.Info().Msg("Added unauthenticated routes")
}

func (r *Router) AuthRoutes() {
	// Login
	r.App.Get("/login/verify", verifyToken)

	// Users
	r.App.Get("/users", getUsers)
	r.App.Patch("/users/", updateUser)
	r.App.Get("/users/:id", getUser)
	r.App.Post("/users", createUser)

	// Statuses
	r.App.Post("/posts", createPost)
	r.App.Delete("/posts/:id", deletePost)

	// Files
	r.App.Post("/files", createFile)
	r.App.Delete("/files/:filename", deleteFile)

	// Bookmarks
	r.App.Get("/bookmarks/:bookmarkId", getBookmarkByID)
	r.App.Get("/posts/:postId/bookmarks", getBookmarksByPostID)
	r.App.Post("/posts/:postId/bookmarks", createBookmark)
	r.App.Delete("/posts/:postId/bookmarks", deleteBookmark)
	r.App.Get("/posts/:postId/bookmarks/:userId", hasUserBookmarked)
	//r.App.Get("/users/:userId/bookmarks", getUserBookmarks)

	// Likes
	r.App.Get("/posts/:id/likes", getLikes)
	r.App.Post("/posts/:id/likes", createLike)
	r.App.Delete("/posts/:id/likes", deleteLike)
	r.App.Get("/posts/:id/likes/:userId", hasUserLiked)

	// Feed
	r.App.Get("/feeds/chronological", getChronologicalFeed)
	r.App.Get("/feeds/algorithmic", getChronologicalFeed)

	// Follows
	r.App.Get("/users/:id/followers", getFollows)
	r.App.Get("/users/:id/following", getFollowing)
	r.App.Post("/users/:id/followers", createFollow)
	r.App.Delete("/users/:id/followers", deleteFollow)

	logger.Log.Info().Msg("Authenticated routes added")
}

func (r *Router) ActivityPubRoutes() {
	//r.App.Get("/users/:username", activitypub.GetUserProfile)
	//r.App.Post("/users/:username/inbox", activitypub.HandleInbox)
	//r.App.Get("/users/:username/outbox", activitypub.HandleOutbox)

	logger.Log.Info().Msg("ActivityPub routes added")
}

func (r *Router) Run() {
	if err := r.App.Listen(":" + cfg.Server.Port); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func getUserID(c *fiber.Ctx) (uuid.UUID, error) {
	if c.Locals("user") == nil {
		logger.Log.Warn().Msg("no user in context of provided token")
		return uuid.Nil, fmt.Errorf("no user in context")
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)
	if !ok {
		logger.Log.Warn().Msg("no sub claim in token")
		return uuid.Nil, fmt.Errorf("no sub claim in token")
	}
	return uuid.Parse(sub)
}
