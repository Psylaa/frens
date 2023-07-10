package router

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/go-playground/validator/v10"
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
	app    *fiber.App
	config *config.Config
	db     *database.Database
	srv    *service.Service
	repos  Repos
}

type Repos struct {
	Auth      *AuthRepo
	Bookmarks *BookmarksRepo
	Feed      *FeedRepo
	Files     *FilesRepo
	Follows   *FollowsRepo
	Likes     *LikesRepo
	Posts     *PostsRepo
	Users     *UsersRepo
}

// New creates a new router instance
func New(cfg *config.Config, db *database.Database, srv *service.Service) *Router {
	app := fiber.New(fiber.Config{})
	r := &Router{
		app:    app,
		config: cfg,
		db:     db,
		srv:    srv,
		repos: Repos{
			Auth:      NewAuthRepo(db, srv),
			Bookmarks: NewBookmarksRepo(db, srv),
			Feed:      NewFeedRepo(db, srv),
			Files:     NewFilesRepo(db, srv),
			Follows:   NewFollowsRepo(db, srv),
			Likes:     NewLikesRepo(db, srv),
			Posts:     NewPostsRepo(db, srv),
			Users:     NewUsersRepo(db, srv),
		},
	}

	r.app.Get("/swagger/*", swagger.HandlerDefault)

	r.configureMiddleware()
	r.configureRoutes()

	return r
}

// Run starts the server
func (r *Router) Run() {
	port := ":" + r.config.Server.Port
	if err := r.app.Listen(port); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func (r *Router) configureMiddleware() {
	r.app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))

	r.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	r.app.Use(limiter.New(limiter.Config{
		Max:        maxRequestsPerSecond,
		Expiration: requestExpiration,
	}))
}

func (r *Router) configureRoutes() {
	v1 := r.app.Group("/v1")

	r.repos.Auth.ConfigurePublicRoutes(v1.Group("/auth"))
	r.repos.Files.ConfigurePublicRoutes(v1.Group("/files"))
	r.addAuth()
	r.repos.Auth.ConfigureProtectedRoutes(v1.Group("/auth"))
	r.repos.Bookmarks.ConfigureRoutes(v1.Group("/bookmarks"))
	r.repos.Feed.ConfigureRoutes(v1.Group("/feeds"))
	r.repos.Files.ConfigureProtectedRoutes(v1.Group("/files"))
	r.repos.Follows.ConfigureRoutes(v1.Group("/follows"))
	r.repos.Likes.ConfigureRoutes(v1.Group("/likes"))
	r.repos.Posts.ConfigureRoutes(v1.Group("/posts"))
	r.repos.Users.ConfigureRoutes(v1.Group("/users"))

	logger.Log.Info().Msg("Configured routes")
}

func (r *Router) addAuth() {
	r.app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(r.config.Server.JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				logger.Log.Warn().Msg("token is missing or malformed: " + err.Error())
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
			}
			logger.Log.Warn().Msg("unknown token error: " + err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		},
	}))

	r.app.Use(r.extractRequestorID)
}

func (r *Router) extractRequestorID(c *fiber.Ctx) error {
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

	userUUID, err := uuid.Parse(sub)
	if err != nil {
		logger.Log.Warn().Msg("invalid sub claim in token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	uuidPtr := &userUUID

	if userUUID == uuid.Nil {
		logger.Log.Warn().Msg("invalid sub claim in token")
		return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
	}

	c.Locals("requestorID", uuidPtr)

	logger.DebugLogRequestUpdate("router", "extractRequestorID", "extractRequestorID", "parsed userID from token: "+uuidPtr.String())
	return c.Next()
}

// Validation utility
func validateRequest(v *validator.Validate, req interface{}) ([]map[string]interface{}, error) {
	if err := v.Struct(req); err != nil {
		// Convert the error into a validation.Errors object
		errList := err.(validator.ValidationErrors)

		// Create a new slice to hold the JSON API error objects
		jsonErrs := createJsonApiErrors(errList)

		// Return the error response
		return jsonErrs, err
	}
	return nil, nil
}

// JSON API errors utility
func createJsonApiErrors(errList validator.ValidationErrors) []map[string]interface{} {
	jsonErrs := make([]map[string]interface{}, len(errList))

	// Loop through the validation errors
	for i, e := range errList {
		// Create an error object
		jsonErr := make(map[string]interface{})

		jsonErr["status"] = "400"
		jsonErr["source"] = map[string]string{"pointer": "/data/attributes/" + e.Field()}
		jsonErr["title"] = "Invalid Attribute"

		detail := fmt.Sprintf("%s validation failed on the '%s' tag", e.Field(), e.ActualTag())
		if e.Param() != "" {
			detail = fmt.Sprintf("%s validation failed on the '%s' tag, condition: %s", e.Field(), e.ActualTag(), e.Param())
		}
		jsonErr["detail"] = detail

		jsonErrs[i] = jsonErr
	}

	return jsonErrs
}
