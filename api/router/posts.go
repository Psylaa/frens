package router

import (
	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

type PostsRepo struct {
	Service *service.PostService
}

func (pr *PostsRepo) addPublicRoutes(rtr fiber.Router) {
}

func (pr *PostsRepo) addPrivateRoutes(rtr fiber.Router) {
	grp := rtr.Group("/posts")
	grp.Post("/", pr.create)
}

func (pr *PostsRepo) create(c *fiber.Ctx) error {
	var req CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return pr.Service.Create(c, req.Text, req.Privacy)
}
