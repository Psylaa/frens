package router

import (
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
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

// @Summary Search Blocks
// @Description Search for blocks with query parameters. If no query parameters are provided, all blocks will be returned. Since blocks are private, only the authenticated user's blocks will be returned.
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
	logger.DebugLogRequestReceived("router", "blocks", "get")

	// Get the query parameters
	queryCount := c.Query("count", "")
	queryOffset := c.Query("offset", "")

	// Parse the query parameters
	count, err := strconv.Atoi(queryCount)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing count parameter")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidCount))
	}

	offset, err := strconv.Atoi(queryOffset)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing offset parameter")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidCount))
	}

	// Send to the service layer
	return br.Srv.Blocks.Get(c, count, offset)
}
