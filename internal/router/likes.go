package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bwoff11/frens/internal/service"
)

type LikesRepo struct {
	Service *service.Service
}

func (lr *LikesRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", lr.get)
}

// @Summary Search Likes
// @Description Retrieve likes. If a like ID is provided, it is always used. Otherwise, a postID will return all likes for that post. If a userID is also provided, it will return either the like for that user/post or an empty array. If only a userID is provided, it will return all likes by that user for any post.
// @Tags Likes
// @Accept  json
// @Produce  json
// @Param postID query string false "Post ID"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /likes [get]
func (lr *LikesRepo) get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{})
}
