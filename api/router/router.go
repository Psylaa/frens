package router

import (
	"github.com/bwoff11/frens/pkg/config"
	"github.com/bwoff11/frens/service"
	"github.com/go-playground/validator"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Router struct {
	App   *fiber.App
	Port  string
	Repos Repos
	Token struct {
		Secret   []byte
		Duration int
	}
}

type Repos struct {
	Auth      *AuthRepo
	Bookmarks *BookmarksRepo
	Feed      *FeedRepo
	Follows   *FollowsRepo
	Likes     *LikesRepo
	Media     *MediaRepo
	Posts     *PostsRepo
	Users     *UsersRepo
}

func New(service *service.Service, config *config.APIConfig) *Router {
	app := fiber.New()

	router := &Router{
		App:  app,
		Port: config.Port,
		Repos: Repos{
			Auth:      &AuthRepo{Service: service.Auth},
			Bookmarks: &BookmarksRepo{Service: service.Bookmark},
			Feed:      &FeedRepo{Service: service.Feed},
			Follows:   &FollowsRepo{Service: service.Follow},
			Likes:     &LikesRepo{Service: service.Like},
			Media:     &MediaRepo{Service: service.Media},
			Posts:     &PostsRepo{Service: service.Post},
			Users:     &UsersRepo{Service: service.User},
		},
		Token: struct {
			Secret   []byte
			Duration int
		}{
			Secret:   []byte(config.TokenSecret),
			Duration: config.TokenDuration,
		},
	}

	addRoutes(router)
	return router
}

func (r *Router) Start() {
	if err := r.App.Listen(":" + r.Port); err != nil {
		panic(err)
	}
}

func addRoutes(router *Router) {
	v1 := router.App.Group("/v1")
	router.Repos.Auth.addPublicRoutes(v1)

	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: router.Token.Secret},
	}))

	router.Repos.Auth.addPrivateRoutes(v1)
	router.Repos.Posts.addPrivateRoutes(v1)
}

/*
func addPublicRoutes(v1 fiber.Router, router *Router) {
	authGroup := v1.Group("/auth")
	authGroup.Post("/login", router.Repos.Auth.Login)
	authGroup.Post("/register", router.Repos.Auth.Register)
}

func addPrivateRoutes(v1 fiber.Router, router *Router) {
		// User routes
		userGroup := v1.Group("/user")
		userGroup.Get("/:id", router.Users.Get)
		userGroup.Post("/", router.Users.Create)
		userGroup.Put("/:id", router.Users.Update)
		userGroup.Delete("/:id", router.Users.Delete)

		// Like routes
		likeGroup := v1.Group("/like")
		likeGroup.Post("/:postId", router.Likes.LikePost)
		likeGroup.Delete("/:postId", router.Likes.UnlikePost)

		// Bookmark routes
		bookmarkGroup := v1.Group("/bookmark")
		bookmarkGroup.Post("/:postId", router.Bookmarks.BookmarkPost)
		bookmarkGroup.Delete("/:postId", router.Bookmarks.UnbookmarkPost)

		// Block routes
		blockGroup := v1.Group("/block")
		blockGroup.Post("/:userId", router.Blocks.BlockUser)
		blockGroup.Delete("/:userId", router.Blocks.UnblockUser)

		// Feed routes
		feedGroup := v1.Group("/feed")
		feedGroup.Get("/chronological", router.Feed.Chronological)
		feedGroup.Get("/algorithmic", router.Feed.Algorithmic)
		feedGroup.Get("/explore", router.Feed.Explore)

		// Media routes
		mediaGroup := v1.Group("/media")
		mediaGroup.Get("/:mediaId", router.Media.Get)
		mediaGroup.Post("/", router.Media.Upload)
		mediaGroup.Delete("/:mediaId", router.Media.Delete)
*/
