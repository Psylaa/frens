package router

import (
	"strconv"
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	_ "github.com/bwoff11/frens/docs"
)

type BookmarksRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewBookmarksRepo(db *database.Database, srv *service.Service) *BookmarksRepo {
	return &BookmarksRepo{
		DB:  db,
		Srv: srv,
	}
}

func (br *BookmarksRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", br.get)
}

// @Summary Retrieve User's Bookmarks
// @Description Retrieves a list of posts bookmarked by the authenticated user.
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param count query string false "The number of bookmarks to return."
// @Param cursor query string false "Cursor for pagination."
// @Success 200 {object} response.BookmarkResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks [get]
func (br *BookmarksRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "bookmarks", "get")

	// Get the query parameters
	queryCount := c.Query("count", "25") // default value is "25"
	queryCursor := c.Query("cursor")     // no default value

	// Convert them to appropriate types
	count, err := strconv.Atoi(queryCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert count to int")
		return c.Status(fiber.StatusBadRequest).SendString("Failed to convert count to int")
	}

	var cursor time.Time
	if queryCursor == "" {
		cursor = time.Now()
	} else {
		cursor, err = time.Parse(time.RFC3339, queryCursor)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse cursor to time")
			return c.Status(fiber.StatusBadRequest).SendString("Failed to parse cursor to time")
		}
	}

	// Send the request to the service layer
	return br.Srv.Bookmarks.Get(c, count, cursor)
}
