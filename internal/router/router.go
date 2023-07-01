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
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	_ "github.com/bwoff11/frens/docs" // For swagger docs
)

var cfg *config.Config
var db *database.Database
var srv *service.Service

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
				return c.Status(fiber.StatusBadRequest).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
			}
			return c.Status(fiber.StatusUnauthorized).JSON(response.GenerateErrorResponse(response.ErrInvalidToken))
		},
	}))

	r.addProtectedRoutes()
	r.addActivityPubRoutes()

	return r
}

func (r *Router) addPublicRoutes() {
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	r.App.Post("/login", login)
	r.App.Get("/files/:filename", retrieveFile)
	r.App.Get("/feeds/explore", getExploreFeed)
	r.App.Get("/posts", GetPostsByUserID)
	r.App.Get("/posts/:id", getPost)

	logger.Log.Info().Msg("Added public routes routes")
}

func (r *Router) addProtectedRoutes() {
	// Login
	r.App.Get("/login/verify", verifyUserToken)

	// Users
	r.App.Get("/users", retrieveAllUsers)
	r.App.Post("/users", registerUser)
	r.App.Patch("/users/", updateUserDetails)
	r.App.Get("/users/:id", retrieveUserDetails)

	/*
		// Posts
		r.App.Post("/posts", createNewPost)
		r.App.Delete("/posts/:id", removePost)

		// Files
		r.App.Post("/files", uploadFile)
		r.App.Delete("/files/:filename", deleteFile)

		// Bookmarks
		r.App.Get("/bookmarks/:bookmarkId", retrieveBookmarkById)
		r.App.Get("/posts/:postId/bookmarks", retrievePostBookmarks)
		r.App.Post("/posts/:postId/bookmarks", addBookmarkToPost)
		r.App.Delete("/posts/:postId/bookmarks", removeBookmarkFromPost)

		// Likes
		r.App.Get("/posts/:id/likes", retrievePostLikes)
		r.App.Post("/posts/:id/likes", addLikeToPost)
		r.App.Delete("/posts/:id/likes", removeLikeFromPost)
		r.App.Get("/posts/:id/likes/:userId", checkUserLike)

		// Feed
		r.App.Get("/feeds/chronological", retrieveChronologicalFeed)
		r.App.Get("/feeds/algorithmic", retrieveAlgorithmicFeed)

		// Follows
		r.App.Get("/users/:id/followers", retrieveFollowers)
		r.App.Get("/users/:id/following", retrieveFollowing)
		r.App.Post("/users/:id/followers", addFollower)
		r.App.Delete("/users/:id/followers", removeFollower)
	*/
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
