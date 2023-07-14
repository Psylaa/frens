package router

import (
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"

	_ "github.com/bwoff11/frens/docs"
)

type BookmarksRepo struct {
	Service *service.Service
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
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /bookmarks [get]
func (br *BookmarksRepo) get(c *fiber.Ctx) error {
	return nil
}
