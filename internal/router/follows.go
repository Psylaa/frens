package router

import (
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FollowsRepo struct {
	Service *service.Service
}

func (fr *FollowsRepo) ConfigureRoutes(rtr fiber.Router) {
}
