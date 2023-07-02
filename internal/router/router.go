package router

import (
	"fmt"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/contrib/fiberzerolog"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var cfg *config.Config
var db *database.Database
var srv *service.Service

// Structure representing the router
type Router struct {
	Config *config.Config
	App    *fiber.App
}

func NewRouter(configuration *config.Config, database *database.Database, service *service.Service) *Router {
	cfg = configuration
	db = database
	srv = service

	app := fiber.New(fiber.Config{})
	r := &Router{cfg, app}

	r.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))
	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	r.addPublicRoutes()

	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.Server.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
			}
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		},
	}))

	r.addProtectedRoutes()
	r.addActivityPubRoutes()

	return r
}

func (r *Router) addPublicRoutes() {

	r.App.Static("/files/default-avatar.png", "./assets/default-avatar.png")
	r.App.Static("/files/default-cover.png", "./assets/default-cover.png")

	r.App.Post("/login", login)
	r.App.Get("/files/:filename", retrieveFile)

	logger.Log.Info().Msg("Added public routes routes")
}

// Order - Type - Get, Post, Patch, Delete - Alphabetical
func (r *Router) addProtectedRoutes() {
	// Login
	r.App.Get("/login/verify", verifyUserToken)

	// Bookmarks
	r.App.Get("/bookmarks/:bookmarkId", getBookmarkByID)
	r.App.Delete("/bookmarks/:bookmarkId", deleteBookmarkByID)

	r.App.Get("/posts/:postId/bookmarks", getBookmarksByPostID)
	r.App.Get("/posts/:postId/bookmarks/count", getBookmarksCountByPostID) // same as getBookmarksByPostID but only returns count
	r.App.Post("/posts/:postId/bookmarks", createBookmarkbyPostID)         // userId from token
	r.App.Delete("/posts/:postId/bookmarks", deleteBookmarkByPostID)       // only callable by owner

	r.App.Get("/users/:userId/bookmarks", getBookmarksByUserID)            // only callable by owner
	r.App.Get("/users/:userId/bookmarks/count", getBookmarksCountByUserID) // same as getBookmarksByUserID but only returns count

	// Likes
	r.App.Get("/likes/:likeId", getLikeByID)
	r.App.Delete("/likes/:likeId", deleteLikeByID)

	r.App.Get("/posts/:postId/likes", getLikesByPostID)
	r.App.Get("/posts/:postId/likes/count", getLikesCountByPostID) // same as getLikesByPostID but only returns count
	r.App.Post("/posts/:postId/likes", createLikeByPostID)         // userId from token
	r.App.Delete("/posts/:postId/likes", deleteLikeByPostID)       // only callable by owner

	r.App.Get("/users/:userId/likes", getLikesByUserID)            // Callable by anyone
	r.App.Get("/users/:userId/likes/count", getLikesCountByUserID) // same as getLikesByUserID but only returns count

	// Users
	r.App.Get("/users", getUsers) // Probably should make this an admin only route
	r.App.Post("/users", createUser)
	r.App.Patch("/users", updateUser)
	r.App.Get("/users/:userId", getUserByID)

	// Posts
	r.App.Get("/posts/:postId", getPostByID)
	r.App.Post("/posts", createPost)
	r.App.Delete("/posts/:id", deletePost)

	// Files
	r.App.Post("/files", createFile)
	r.App.Delete("/files/:filename", deleteFile)

	// Feed
	r.App.Get("/feeds/chronological", getChronoFeed)
	r.App.Get("/feeds/algorithmic", getAlgoFeed)
	r.App.Get("/feeds/explore", getExploreFeed)

	// Follows
	r.App.Get("/users/:userId/followers", getFollowersByUserID)
	r.App.Get("/users/:userId/following", getFollowingByUserID)
	r.App.Post("/users/:userId/followers", createFollowByUserID)
	r.App.Delete("/users/:userId/followers", deleteFollowerByUserID)

	logger.Log.Info().Msg("Protected routes added")
}

func (r *Router) addActivityPubRoutes() {
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
