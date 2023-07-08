package router

import (
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// @Summary Search Bookmarks
// @Description Search for bookmarks with query parameters. If no query parameters are provided, all bookmarks will be returned. Since bookmarks are private, only the authenticated user's bookmarks will be returned.
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param bookmarkID query string false "The ID of a specific bookmark to retrieve"
// @Param count query string false "The number of bookmarks to return."
// @Param offset query string false "The number of bookmarks to offset the returned bookmarks by. For example, offset=10&count=10 would return bookmarks 10-20"
// @Success 200 {object} response.BookmarkResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks [get]
func (br *BookmarksRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "bookmarks", "get")

	// Get the requestorID from the token
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the query parameters
	bookmarkID := c.Query("bookmarkID", "")
	queryCount := c.Query("count", "")
	queryOffset := c.Query("offset", "")

	// If bookmark is not nil, convert it to a UUID
	var bookmarkUUID uuid.UUID
	var err error
	if bookmarkID != "" {
		bookmarkUUID, err = uuid.Parse(bookmarkID)
		if err != nil {
			logger.DebugLogRequestUpdate("router", "bookmarks", "get", "Error parsing bookmarkID: "+bookmarkID)
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}
	}

	// If a bookmarkID is provided, fetch the specific bookmark
	if bookmarkID != "" {
		return br.Srv.Bookmarks.GetByID(c, &bookmarkUUID)
	}

	// If a count was provided, parse it
	var count *int
	if queryCount != "" {
		countInt, err := strconv.Atoi(queryCount)
		if err != nil {
			logger.DebugLogRequestUpdate("router", "bookmarks", "get", "Error parsing count: "+queryCount)
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidCount))
		}
		count = &countInt
	}

	// If an offset was provided, parse it
	var offset *int
	if queryOffset != "" {
		offsetInt, err := strconv.Atoi(queryOffset)
		if err != nil {
			log.Error().Err(err).Msg("Error parsing offset: " + queryOffset)
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidOffset))
		}
		offset = &offsetInt
	}

	// Send request to service layer
	return br.Srv.Bookmarks.GetByUserID(c, requestorID, count, offset)
}
