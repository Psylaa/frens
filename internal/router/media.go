package router

import (
	"github.com/bwoff11/frens/internal/models"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type MediaRepo struct {
	Service *service.Service
}

func (mr *MediaRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", mr.get)
	rtr.Post("/", mr.create)
	rtr.Delete("/:id", mr.delete)
}

// @Summary Get Media
// @Description Retrieve media by ID
// @Tags Media
// @Accept  json
// @Produce  json
// @Param id path string true "Media ID"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /media/{id} [get]
func (r *MediaRepo) get(c *fiber.Ctx) error {
	return nil
}

// @Summary Create Media
// @Description Create media
// @Tags Media
// @Accept  json
// @Produce  json
// @Param file formData file true "File"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /media [post]
func (r *MediaRepo) create(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return models.ErrInvalidBody.SendResponse(c, err.Error())
	}

	return r.Service.Media.Create(c, file)
}

// @Summary Delete Media
// @Description Delete media by ID
// @Tags Media
// @Accept  json
// @Produce  json
// @Param id path string true "Media ID"
// @Success 200
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /media/{id} [delete]
func (r *MediaRepo) delete(c *fiber.Ctx) error {
	return nil
}
