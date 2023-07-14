package router

import (
	"log"
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/models"
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

const (
	maxRequestsPerSecond = 1000
	requestExpiration    = 30 * time.Second
)

type Router struct {
	App       *fiber.App
	Port      string
	JWTSecret []byte
	Repos     Repos
}

type Repos struct {
	Auth      *AuthRepo
	Bookmarks *BookmarksRepo
	Feed      *FeedRepo
	Follows   *FollowsRepo
	Likes     *LikesRepo
	Posts     *PostsRepo
	Users     *UsersRepo
}

// New creates a new router instance
func New(configuration *config.Config) *Router {

	// Create service
	service := service.New(configuration)

	r := &Router{
		Repos: Repos{
			Auth:      &AuthRepo{Service: service},
			Bookmarks: &BookmarksRepo{Service: service},
			Feed:      &FeedRepo{Service: service},
			Follows:   &FollowsRepo{Service: service},
			Likes:     &LikesRepo{Service: service},
			Posts:     &PostsRepo{Service: service},
			Users:     &UsersRepo{Service: service},
		},
	}

	// Store config values
	r.Port = configuration.Server.Port
	r.JWTSecret = []byte(configuration.Server.JWTSecret)

	// Create fiber App
	r.App = fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	r.App.Get("/swagger/*", swagger.HandlerDefault)

	r.configureMiddleware()
	r.configureRoutes()

	return r
}

// Run starts the server
func (r *Router) Run() {

	port := ":" + r.Port
	if err := r.App.Listen(port); err != nil {
		log.Fatal(err)
	}
}

func (r *Router) configureMiddleware() {
	// Create separate logger for fiber
	r.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Logger,
	}))

	r.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	r.App.Use(limiter.New(limiter.Config{
		Max:        maxRequestsPerSecond,
		Expiration: requestExpiration,
	}))
}

func (r *Router) configureRoutes() {
	v1 := r.App.Group("/v1")

	r.Repos.Auth.ConfigurePublicRoutes(v1.Group("/auth"))
	r.addAuth()
	r.Repos.Auth.ConfigureProtectedRoutes(v1.Group("/auth"))
	r.Repos.Bookmarks.ConfigureRoutes(v1.Group("/bookmarks"))
	r.Repos.Feed.ConfigureRoutes(v1.Group("/feeds"))
	r.Repos.Follows.ConfigureRoutes(v1.Group("/follows"))
	r.Repos.Likes.ConfigureRoutes(v1.Group("/likes"))
	r.Repos.Posts.ConfigureRoutes(v1.Group("/posts"))
	r.Repos.Users.ConfigureRoutes(v1.Group("/users"))

	logger.Info(logger.LogMessage{
		Package:  "router",
		Function: "configureRoutes",
		Message:  "routes configured",
	})
}

func (r *Router) addAuth() {
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: r.JWTSecret},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return models.ErrInvalidToken.SendResponse(c)
		},
	}))

	r.App.Use(r.extractRequestorID)
}

func (r *Router) extractRequestorID(c *fiber.Ctx) error {
	if c.Locals("user") == nil {
		return models.ErrInvalidToken.SendResponse(c, "no user in context")
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)

	if !ok {
		return models.ErrInvalidToken.SendResponse(c, "no sub claim in token")
	}

	userUUID, err := uuid.Parse(sub)
	if err != nil {
		return models.ErrInvalidToken.SendResponse(c, "invalid sub claim in token")
	}

	uuidPtr := &userUUID

	if userUUID == uuid.Nil {
		return models.ErrInvalidToken.SendResponse(c, "invalid sub claim in token")
	}

	c.Locals("requestorID", uuidPtr)
	return c.Next()
}
