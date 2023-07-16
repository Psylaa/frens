package router

import (
	"github.com/bwoff11/frens/pkg/config"
	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

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
			Auth:      &AuthRepo{Service: service.User},
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

func (r *Router) Start() error {
	return r.App.Listen(":" + r.Port)
}

func addRoutes(router *Router) {
	addPrivateRoutes(router)
	//middleware.AddAuthenticator(router.Token.Secret)
	addPublicRoutes(router)
}

func addPublicRoutes(router *Router) {
	// Authentication routes
	authGroup := router.App.Group("/auth")
	authGroup.Post("/login", router.Repos.Auth.Login)
	authGroup.Post("/register", router.Repos.Auth.Register)
}

func addPrivateRoutes(router *Router) {
	/*
		// User routes
		userGroup := app.Group("/user")
		userGroup.Get("/:id", router.Users.Get)
		userGroup.Post("/", router.Users.Create)
		userGroup.Put("/:id", router.Users.Update)
		userGroup.Delete("/:id", router.Users.Delete)

		// Like routes
		likeGroup := app.Group("/like")
		likeGroup.Post("/:postId", router.Likes.LikePost)
		likeGroup.Delete("/:postId", router.Likes.UnlikePost)

		// Bookmark routes
		bookmarkGroup := app.Group("/bookmark")
		bookmarkGroup.Post("/:postId", router.Bookmarks.BookmarkPost)
		bookmarkGroup.Delete("/:postId", router.Bookmarks.UnbookmarkPost)

		// Block routes
		blockGroup := app.Group("/block")
		blockGroup.Post("/:userId", router.Blocks.BlockUser)
		blockGroup.Delete("/:userId", router.Blocks.UnblockUser)

		// Feed routes
		feedGroup := app.Group("/feed")
		feedGroup.Get("/chronological", router.Feed.Chronological)
		feedGroup.Get("/algorithmic", router.Feed.Algorithmic)
		feedGroup.Get("/explore", router.Feed.Explore)

		// Media routes
		mediaGroup := app.Group("/media")
		mediaGroup.Get("/:mediaId", router.Media.Get)
		mediaGroup.Post("/", router.Media.Upload)
		mediaGroup.Delete("/:mediaId", router.Media.Delete)
	*/
}
