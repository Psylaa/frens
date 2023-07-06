package router_test

import (
	"testing"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MockRouter struct {
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

func TestRouter_extractRequestorID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}

	r := &Router{DB: &MockDB{}}

	tests := []struct {
		name    string
		r       *Router
		args    args
		wantErr bool
	}{
		{
			name: "No user in local claims",
			r:    r,
			args: args{
				c: &fiber.Ctx{},
			},
			wantErr: true,
		},
		{
			name: "Invalid UUID in user claim",
			r:    r,
			args: args{
				c: func() *fiber.Ctx {
					token := jwt.New(jwt.SigningMethodHS256)
					claims := token.Claims.(jwt.MapClaims)
					claims["sub"] = "invalidUUID"
					c := &fiber.Ctx{}
					c.Locals("user", token)
					return c
				}(),
			},
			wantErr: true,
		},
		{
			name: "Valid UUID but no user exists in database",
			r:    r,
			args: args{
				c: func() *fiber.Ctx {
					token := jwt.New(jwt.SigningMethodHS256)
					claims := token.Claims.(jwt.MapClaims)
					claims["sub"] = "00000000-0000-0000-0000-000000000000"
					c := &fiber.Ctx{}
					c.Locals("user", token)
					return c
				}(),
			},
			wantErr: true,
		},
		{
			name: "Valid UUID and user exists in database",
			r:    r,
			args: args{
				c: func() *fiber.Ctx {
					token := jwt.New(jwt.SigningMethodHS256)
					claims := token.Claims.(jwt.MapClaims)
					claims["sub"] = uuid.New().String()
					c := &fiber.Ctx{}
					c.Locals("user", token)
					return c
				}(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.extractRequestorID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Router.extractRequestorID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
