package router

import (
	"github.com/bwoff11/frens/internal/models"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FeedRepo struct {
	Service *service.Service
}

func (fr *FeedRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/chronological", fr.getChrono)
	rtr.Get("/algorithmic", fr.getChrono) // temp until algo is finished
	rtr.Get("/explore", fr.getChrono)     // temp
}

// @Summary Retrieve Algorithmic Feed
// @Description Retrieves the authenticated user's feed, sorted by an algorithm to highlight relevant content.
// @Tags Feed
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /feeds/algorithmic [get]
func (fr *FeedRepo) getAlgo(c *fiber.Ctx) error {
	return nil
}

// @Summary Retrieve Chronological Feed
// @Description Retrieves the authenticated user's feed, sorted by the time of the post's creation.
// @Tags Feed
// @Accept  json
// @Produce  json
// @Param cursor query string false "Cursor for pagination"
// @Param count query int false "Number of posts to return"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /feeds/chronological [get]
func (fr *FeedRepo) getChrono(c *fiber.Ctx) error {

	var req models.FeedRequest
	if err := c.QueryParser(&req); err != nil {
		return err
	}

	return fr.Service.Feed.GetChrono(c, req)
}

// @Summary Retrieve Explore Feed
// @Description Retrieves a feed of trending or recommended content for the authenticated user to discover.
// @Tags Feed
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /feeds/explore [get]
func (fr *FeedRepo) getExplore(c *fiber.Ctx) error {
	return nil
}
