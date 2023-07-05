package router

import (
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
	rtr.Post("/:postId", br.create)
	rtr.Delete("/", br.delete)
}

// @Summary Get bookmarks
// @Description Get bookmarks
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param userId query string false "The ID of the user to get bookmarks for. Given bookmarks are private, this must be the same as the requestor (defaults to the requestor). Admins can get any user's bookmarks"
// @Param count query string false "The number of bookmarks to return."
// @Param offset query string false "The number of bookmarks to offset the returned bookmarks by e.g. offset=10&count=10 would return bookmarks 10-20"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks [get]
func (br *BookmarksRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "bookmarks", "get")
	return br.Srv.Bookmarks.GetSelf(c, c.Locals("requestorId").(*uuid.UUID))
}

// @Summary Create a bookmark for a post
// @Description Create a bookmark for a specific post based on the provided ID
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param postId path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks/{:postId} [post]
func (br *BookmarksRepo) create(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "bookmarks", "create")

	// Parse the post ID
	postId := c.Params("postId")
	postID, err := uuid.Parse(postId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing post ID: " + postId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send request to service
	return br.Srv.Bookmarks.Create(c, c.Locals("requestorID").(*uuid.UUID), &postID)
}

// @Summary Delete a bookmark by ID
// @Description Delete a specific bookmark. Either a bookmark ID or a post ID must be provided. If both are provided, the bookmark ID will be used. Only the owner of the bookmark can delete it. Admins can delete any bookmark.
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param bookmarkId query string true "The ID of the bookmark to delete"
// @Param postId query string true "The ID of the post to delete the bookmark for"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks/{bookmarkId} [delete]
func (br *BookmarksRepo) delete(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "bookmarks", "delete")

	// Get post ID from request
	id := c.Params("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Send request to service layer
	return br.Srv.Bookmarks.DeleteByUserAndPostID(c, c.Locals("requestorId").(*uuid.UUID), &postID)
}
