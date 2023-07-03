package router

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/contrib/fiberzerolog"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/gofiber/swagger"
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

	r.addPublicMiddleware()
	r.addPublicRoutes()
	r.addAuthMiddleware()
	r.addProtectedRoutes()
	r.addActivityPubRoutes()

	return r
}

func (r *Router) addPublicMiddleware() {
	r.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))
	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	r.App.Use(limiter.New(limiter.Config{
		Max:        1000,
		Expiration: 30 * time.Second,
	}))
	r.App.Use(func(c *fiber.Ctx) error {
		requestorId, err := getUserIDFromToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		}
		c.Locals("requestorId", requestorId)
		return c.Next()
	})
}

func (r *Router) addAuthMiddleware() {
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.Server.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
			}
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		},
	}))
}

func (r *Router) addPublicRoutes() {

	r.App.Static("/files/default-avatar.png", "./assets/default-avatar.png")
	r.App.Static("/files/default-cover.png", "./assets/default-cover.png")

	r.App.Post("/login", login)
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	logger.Log.Info().Msg("Added public routes routes")
}

// Order - Type - Get, Post, Patch, Delete - Alphabetical
func (r *Router) addProtectedRoutes() {
	// Login
	r.App.Get("/login/verify", verifyUserToken)

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

	r.App.Get("/users/:userId/likes", getLikesByUserID)            // only callable by owner
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
	r.App.Get("/users/:userId/posts", getPostsByUserID)

	// Files
	r.App.Get("/files/:fileId", retrieveFileByID)
	r.App.Post("/files", createFile)
	r.App.Delete("/files/:fileId", deleteFileByID)

	// Feed
	r.App.Get("/feeds/chronological", getChronoFeed)
	r.App.Get("/feeds/algorithmic", getChronoFeed) // temp
	r.App.Get("/feeds/explore", getChronoFeed)     // temp

	// Follows
	r.App.Get("/users/:userId/followers", getFollowersByUserID)
	r.App.Get("/users/:userId/following", getFollowingByUserID)
	r.App.Post("/users/:userId/followers", createFollowByUserID)
	r.App.Delete("/users/:userId/followers", deleteFollowerByUserID)

	// Future possible routes
	// r.App.Get("/posts/:postId/likes/user/:userId", checkIfPostIsLikedByUser)
	// r.App.Get("/posts/:postId/bookmarks/user/:userId", checkIfPostIsBookmarkedByUser)
	// r.App.Get("/users/:userId/followers/user/:userId", checkIfUserIsFollowedByUser)

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

// In many, many functions, the requestors userID is required.
// This function will get the userID from the token and return it as a pointer.
// It can then be added to the context and used in any function that requires it.
func getUserIDFromToken(c *fiber.Ctx) (*uuid.UUID, error) {

	// Parse claims from token
	if c.Locals("user") == nil {
		logger.Log.Warn().Msg("no user in context of provided token")
		return nil, fmt.Errorf("no user in context")
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)
	if !ok {
		logger.Log.Warn().Msg("no sub claim in token")
		return nil, fmt.Errorf("no sub claim in token")
	}

	// Parse user id as uuid
	userUUID, err := uuid.Parse(sub)
	if err != nil {
		logger.Log.Warn().Msg("invalid sub claim in token")
		return nil, fmt.Errorf("invalid sub claim in token")
	}
	// Convert to pointer
	uuidPtr := &userUUID

	// Validate uuid is not nil (a bit redundant but whatever)
	if userUUID == uuid.Nil {
		logger.Log.Warn().Msg("invalid sub claim in token")
		return nil, fmt.Errorf("invalid sub claim in token")
	}

	// Verify user exists
	if exists := db.Users.Exists(uuidPtr); exists == false {
		logger.Log.Warn().Msg("user does not exist")
		return nil, fmt.Errorf("user does not exist")
	}

	return uuidPtr, nil
}
