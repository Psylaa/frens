package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type BlockRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewBlocksRepo(db *database.Database, srv *service.Service) *BlockRepo {
	return &BlockRepo{
		DB:  db,
		Srv: srv,
	}
}

func (br *BlockRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", br.get)
}

// @Summary Retrieve Blocked Users
// @Description Retrieves a list of users blocked by the authenticated user.
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param count query string false "The number of blocks to return."
// @Param offset query string false "The number of blocks to offset the returned blocks by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /blocks [get]
func (br *BlockRepo) get(c *fiber.Ctx) error {
	return nil
}
