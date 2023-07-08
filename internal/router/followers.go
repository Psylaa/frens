package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FollowersRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewFollowersRepo(db *database.Database, srv *service.Service) *FollowsRepo {
	return &FollowsRepo{
		DB:  db,
		Srv: srv,
	}
}

func (fr *FollowersRepo) ConfigureRoutes(rtr fiber.Router) {
}

// @Summary Get a list of all users that are following the authenticated user
// @Description Get a list of all users that are following the authenticated user
// @Tags Followers
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /followers [get]
func (fr *FollowersRepo) GetFollowers(c *fiber.Ctx) error {
	return nil
}

// @Summary Get a count of all users that are following the authenticated user
// @Description Get a count of all users that are following the authenticated user
// @Tags Followers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /followers/count [get]
func (fr *FollowsRepo) GetFollowersCount(c *fiber.Ctx) error {
	return nil
}
