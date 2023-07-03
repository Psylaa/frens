package router

import (
	"github.com/bwoff11/frens/internal/database"
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
	rtr.Get("/", br.getSelf)
	rtr.Post("/", br.create)
	rtr.Delete("/:id", br.delete)
}

// @Summary Get bookmarks
// @Description Get bookmarks
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param userId query string false "User ID"
// @Param count query string false "Count"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks [get]
func (br *BookmarksRepo) getSelf(c *fiber.Ctx) error {
	return br.Srv.Bookmarks.GetSelf(c, c.Locals("requestorId").(*uuid.UUID))
}

// @Summary Create a bookmark for a post
// @Description Create a bookmark for a specific post based on the provided ID
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param postId body string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks [post]
func (br *BookmarksRepo) create(c *fiber.Ctx) error {
	postId := c.Params("postId")
	postID, err := uuid.Parse(postId)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing post ID: " + postId)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	return br.Srv.Bookmarks.Create(c, c.Locals("requestorID").(*uuid.UUID), &postID)
}

// @Summary Delete a bookmark by ID
// @Description Delete a specific bookmark based on the provided ID
// @Tags Bookmarks
// @Accept  json
// @Produce  json
// @Param bookmarkId path string true "Bookmark ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks/{bookmarkId} [delete]
func (br *BookmarksRepo) delete(c *fiber.Ctx) error {

	// Get post ID from request
	id := c.Params("id")
	postID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status ID"})
	}

	// Send request to service
	return br.Srv.Bookmarks.DeleteByUserAndPostID(c, c.Locals("requestorId").(*uuid.UUID), &postID)
}
