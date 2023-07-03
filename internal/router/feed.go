package router

import (
	"net/http"
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FeedRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewFeedRepo(db *database.Database, srv *service.Service) *FeedRepo {
	return &FeedRepo{
		DB:  db,
		Srv: srv,
	}
}

func (fr *FeedRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/chrono", fr.getChrono)
	rtr.Get("/algo", fr.getAlgo)
	rtr.Get("/explore", fr.getExplore)
}

// getChrono handles the request to get the chronological feed for a user.
// @Summary Get chronological feed
// @Description Get posts in a chronological order
// @Tags Feed
// @Accept  json
// @Produce  json
// @Param cursor query string false "Cursor for pagination"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /feed/chrono [get]
func (fr *FeedRepo) getChrono(c *fiber.Ctx) error {

	// Get the cursor from the query string
	cursorString := c.Query("cursor")
	var cursor time.Time
	if cursorString == "" {
		cursor = time.Now()
	} else {
		var err error
		cursor, err = time.Parse(time.RFC3339, cursorString)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing cursor")
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
	}

	return fr.Srv.Feed.GetChrono(c, c.Locals("requestorId").(*uuid.UUID), cursor)
}

// getAlgo handles the request to get the algorithm based feed for a user.
// @Summary Get algorithm-based feed
// @Description Get posts based on a particular algorithm (Not Implemented)
// @Tags Feed
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /feed/algo [get]
func (fr *FeedRepo) getAlgo(c *fiber.Ctx) error {
	return nil
}

// getExplore handles the request to get the explore feed for a user.
// @Summary Get explore feed
// @Description Get posts for the explore section (Not Implemented)
// @Tags Feed
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /feed/explore [get]
func (fr *FeedRepo) getExplore(c *fiber.Ctx) error {
	return nil
}
