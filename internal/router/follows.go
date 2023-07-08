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
