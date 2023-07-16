package router

import (
	"strconv"

	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

type FeedRepo struct {
	Service *service.FeedService
}

func (fr *FeedRepo) addPublicRoutes(rtr fiber.Router) {
}

func (fr *FeedRepo) addPrivateRoutes(rtr fiber.Router) {
	grp := rtr.Group("/feeds")
	//grp.Get("/algorithmic", fr.getAlgorithmic)
	grp.Get("/chronological", fr.getChronological)
	//grp.Get("/explore", fr.getExplore)
}

func (fr *FeedRepo) getChronological(c *fiber.Ctx) error {
	var req ChronoFeedRequest

	// Get the count parameter from the query string
	count := c.Query("count", "0") // the second argument is a default value
	countUint, err := strconv.ParseUint(count, 10, 8)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid count parameter")
	}
	req.Count = uint8(countUint)

	// Get the cursor parameter from the query string
	cursor := c.Query("cursor", "0") // the second argument is a default value
	cursorUint, err := strconv.ParseUint(cursor, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid cursor parameter")
	}
	req.Cursor = uint32(cursorUint)

	if err := validate.Struct(req); err != nil {
		return err
	}

	return fr.Service.GetChronological(c, req.Count, req.Cursor)
}
