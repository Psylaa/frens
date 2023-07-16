package service

import (
	"fmt"

	"github.com/bwoff11/frens/pkg/config"
	"github.com/bwoff11/frens/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	Auth     *AuthService
	Block    *BlockService
	Bookmark *BookmarkService
	Feed     *FeedService
	Follow   *FollowService
	Like     *LikeService
	Media    *MediaService
	Post     *PostService
	User     *UserService
}

func New(db *database.Database, config *config.APIConfig) *Service {
	return &Service{
		Auth: &AuthService{
			Database:    db,
			JWTSecret:   []byte(config.TokenSecret),
			JWTDuration: config.TokenDuration,
		},
		Block:    &BlockService{Database: db},
		Bookmark: &BookmarkService{Database: db},
		Feed:     &FeedService{Database: db},
		Follow:   &FollowService{Database: db},
		Like:     &LikeService{Database: db},
		Media:    &MediaService{Database: db},
		Post:     &PostService{Database: db},
		User:     &UserService{Database: db},
	}
}

func getRequestorID(c *fiber.Ctx) (uint32, error) {
	// Retrieve the user from the JWT
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)
	if !ok {
		return 0, fiber.ErrUnauthorized
	}

	// Convert the user ID to a uint32
	var id uint32
	if _, err := fmt.Sscanf(sub, "%d", &id); err != nil {
		return 0, fiber.ErrUnauthorized
	}

	return id, nil
}
