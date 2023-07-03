package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type LoginRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewLoginRepo(db *database.Database, srv *service.Service) *LoginRepo {
	return &LoginRepo{
		DB:  db,
		Srv: srv,
	}
}

func (lr *LoginRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Post("/login", lr.login)
	rtr.Get("/verify", lr.verifyToken)
}

func (lr *LoginRepo) login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	if body.Username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	return lr.Srv.Login.Login(c, &body.Username, &body.Password)
}

func (lr *LoginRepo) verifyToken(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
