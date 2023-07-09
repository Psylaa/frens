package service

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BlockRepo struct{}

func (br *BlockRepo) Get(c *fiber.Ctx, count int, offset int) error {
	logger.DebugLogRequestReceived("service", "blocks", "Get")

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the blocks from the database
	blocks, err := db.Blocks.GetByUserID(requestorID, &count, &offset)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error getting blocks")
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(response.CreateBlocksResponse(blocks))
}
