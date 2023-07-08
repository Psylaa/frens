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

// @Summary Get blocks
// @Description Get blocks for the authenticated user.
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

// @Summary Get a block by block ID
// @Description Retrieve a block by block ID
// @Tags Blocks
// @Accept json
// @Produce json
// @Param blockID path string true "Block ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /blocks/{blockID} [get]
func (br *BlockRepo) getByID(c *fiber.Ctx) error {
	return nil
}

// @Summary Delete all blocks
// @Description Delete all blocks for the authenticated user.
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /blocks [delete]
func (br *BlockRepo) deleteAll(c *fiber.Ctx) error {
	return nil
}

// @Summary Delete a block by block ID
// @Description Delete a block by block ID
// @Tags Blocks
// @Accept json
// @Produce json
// @Param blockID path string true "Block ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /blocks/{blockID} [delete]
func (br *BlockRepo) deleteByID(c *fiber.Ctx) error {
	return nil
}
