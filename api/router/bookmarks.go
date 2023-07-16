package router

import (
	"strconv"

	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

type BookmarksRepo struct {
	Service *service.BookmarkService
}

func (br *BookmarksRepo) addPrivateRoutes(rtr fiber.Router) {
	grp := rtr.Group("/bookmarks")
	grp.Post("/:postID", br.bookmarkPost)
	grp.Delete("/:postID", br.unbookmarkPost)
}

func (br *BookmarksRepo) bookmarkPost(c *fiber.Ctx) error {
	postID, err := strconv.ParseUint(c.Params("postID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid post ID")
	}

	return br.Service.BookmarkPost(c, uint32(postID))
}

func (br *BookmarksRepo) unbookmarkPost(c *fiber.Ctx) error {
	postID, err := strconv.ParseUint(c.Params("postID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid post ID")
	}

	return br.Service.UnbookmarkPost(c, uint32(postID))
}
