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
		//AllowOrigins: cfg.Server.AllowOrigins,
		AllowOrigins: "*",
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
		Error: APIResponseErr(err.Error()),
	})
}

func (r *Router) setupRoutes() {
	r.addUnauthRoutes()

	// Middleware for JWT authentication
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.Server.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
					Error: ErrMissingToken,
				})
			}

			return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
				Error: ErrUnauthorized,
			})
		},
	}))

	// Add authenticated routes
	r.AuthRoutes()
	r.ActivityPubRoutes()
}

func (r *Router) addUnauthRoutes() {
	// Unauthenticated routes
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
	r.App.Get("/users/:id", getUser)
	r.App.Post("/users", createUser)
	r.App.Patch("/users/:id", updateUser)

	// Statuses
	r.App.Post("/posts", createPost)
	r.App.Delete("/posts/:id", deletePost)

	// Files
	r.App.Post("/files", createFile)
	r.App.Delete("/files/:filename", deleteFile)

	// Bookmarks
	r.App.Get("/statuses/:id/bookmarks", getBookmarks)
	r.App.Post("/statuses/:id/bookmarks", createBookmark)
	r.App.Delete("/statuses/:id/bookmarks", deleteBookmark)
	r.App.Get("/statuses/:id/bookmarks/:userId", hasUserBookmarked)

	// Likes
	r.App.Get("/statuses/:id/likes", getLikes)
	r.App.Post("/statuses/:id/likes", createLike)
	r.App.Delete("/statuses/:id/likes", deleteLike)
	r.App.Get("/statuses/:id/likes/:userId", hasUserLiked)

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
