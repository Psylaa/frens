package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FollowsRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewFollowsRepo(db *database.Database, srv *service.Service) *FollowsRepo {
	return &FollowsRepo{
		DB:  db,
		Srv: srv,
	}
}

func (fr *FollowsRepo) ConfigureRoutes(rtr fiber.Router) {
}

// @Summary Get a list of all users that the authenticated user is following
// @Description Get a list of all users that the authenticated user is following
// @Tags Follows
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /follows [get]
func (fr *FollowsRepo) GetFollows(c *fiber.Ctx) error {
	return nil
}

// @Summary Get a count of all users that the authenticated user is following
// @Description Get a count of all users that the authenticated user is following
// @Tags Follows
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /follows/count [get]
func (fr *FollowsRepo) GetFollowsCount(c *fiber.Ctx) error {
	return nil
}

// @Summary Get a list of all users that are following a user by user ID
// @Description Get a list of all users that are following a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /follows/{userID} [get]
func (fr *FollowsRepo) GetFollowsByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Get a count of all users that are following a user by user ID
// @Description Get a count of all users that are following a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /follows/{userID}/count [get]
func (fr *FollowsRepo) GetFollowsCountByUserID(c *fiber.Ctx) error {
	return nil
}
