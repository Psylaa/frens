package router

// Separate imports into three groups: standard library, third-party, and internal
import (
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
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Router struct {
	App    *fiber.App
	Config *config.Config
	DB     *database.Database
	Srv    *service.Service
	Repos  Repos
}

type Repos struct {
	Bookmarks *BookmarksRepo
	Feed      *FeedRepo
	Files     *FilesRepo
	Follows   *FollowsRepo
	Likes     *LikesRepo
	Login     *LoginRepo
	Posts     *PostsRepo
	Users     *UsersRepo
}

var tokenBlacklist []string

// Initialize a new router
func New(cfg *config.Config, db *database.Database, srv *service.Service) *Router {
	app := fiber.New(fiber.Config{})
	r := &Router{
		App:    app,
		Config: cfg,
		DB:     db,
		Srv:    srv,
		Repos: Repos{
			Bookmarks: NewBookmarksRepo(db, srv),
			Feed:      NewFeedRepo(db, srv),
			Files:     NewFilesRepo(db, srv),
			Follows:   NewFollowsRepo(db, srv),
			Likes:     NewLikesRepo(db, srv),
			Login:     NewLoginRepo(db, srv),
			Posts:     NewPostsRepo(db, srv),
			Users:     NewUsersRepo(db, srv),
		},
	}

	r.App.Get("/swagger/*", swagger.HandlerDefault) // Move at some point. Just needed to get it working.

	r.configureMiddleware()
	r.configureRoutes()

	return r
}

func (r *Router) configureMiddleware() {

	// ZeroLog
	r.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))

	// CORS
	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Limiters
	r.App.Use(limiter.New(limiter.Config{
		Max:        1000,
		Expiration: 30 * time.Second,
	}))
}

func (r *Router) configureRoutes() {
	v1 := r.App.Group("/v1")

	r.Repos.Login.ConfigureRoutes(v1.Group("/login"))
	r.addAuth()
	r.Repos.Bookmarks.ConfigureRoutes(v1.Group("/bookmarks"))
	r.Repos.Feed.ConfigureRoutes(v1.Group("/feeds"))
	r.Repos.Files.ConfigureRoutes(v1.Group("/files"))
	r.Repos.Follows.ConfigureRoutes(v1.Group("/follows"))
	r.Repos.Likes.ConfigureRoutes(v1.Group("/likes"))
	r.Repos.Posts.ConfigureRoutes(v1.Group("/posts"))
	r.Repos.Users.ConfigureRoutes(v1.Group("/users"))

	logger.Log.Info().Msg("Configured routes")
}

// Run starts the server
func (r *Router) Run() {
	port := ":" + r.Config.Server.Port
	if err := r.App.Listen(port); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func (r *Router) addAuth() {

	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(r.Config.Server.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
			}
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		},
	}))

	// Requestor ID
	r.App.Use(r.extractRequestorID)
}

// Middleware function to extract the user ID from the token, validate it, and add it to the context:
func (r *Router) extractRequestorID(c *fiber.Ctx) error {
	// Parse claims from token
	if c.Locals("user") == nil {
		logger.Log.Warn().Msg("no user in context of provided token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)

	if !ok {
		logger.Log.Warn().Msg("no sub claim in token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	// Parse user id as uuid
	userUUID, err := uuid.Parse(sub)
	if err != nil {
		logger.Log.Warn().Msg("invalid sub claim in token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	// Convert to pointer
	uuidPtr := &userUUID

	// Validate uuid is not nil (a bit redundant but whatever)
	if userUUID == uuid.Nil {
		logger.Log.Warn().Msg("invalid sub claim in token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	// Verify user exists
	//if exists := r.DB.Users.Exists(uuidPtr); exists == false {
	//	logger.Log.Warn().Msg("user does not exist")
	//	return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	//}

	c.Locals("requestorId", uuidPtr)

	return c.Next()
}
