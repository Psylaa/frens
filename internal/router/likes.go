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
	rtr.Get("/:id", lr.getByID)
	rtr.Delete("/:id", lr.deleteByID)
}

func (lr *LikesRepo) get(c *fiber.Ctx) error {
	return nil
}

func (lr *LikesRepo) getByID(c *fiber.Ctx) error {
	return nil
}

func (lr *LikesRepo) deleteByID(c *fiber.Ctx) error {
	return nil
}
