package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/service"
)

type LikesRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewLikesRepo(db *database.Database, srv *service.Service) *LikesRepo {
	return &LikesRepo{
		DB:  db,
		Srv: srv,
	}
}

func (lr *LikesRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", lr.get)
	rtr.Post("/", lr.create)
	rtr.Delete("/:id", lr.deleteByID)
}

// @Summary Get likes
// @Description Retrieve a like object consisting of a user and a target post
// @Tags Likes
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes [get]
func (lr *LikesRepo) get(c *fiber.Ctx) error {
	return nil
}

// @Summary Create a like
// @Description Create a new like for a user based on the provided token
// @Tags Likes
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes/{:postId} [post]
func (lr *LikesRepo) create(c *fiber.Ctx) error {
	return nil
}

// @Summary Delete a like
// @Description Delete a specific like
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param likeId path string true "Like ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes [delete]
func (lr *LikesRepo) deleteByID(c *fiber.Ctx) error {
	return nil
}
