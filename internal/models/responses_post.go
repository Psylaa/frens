package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostResponse struct {
	Links PostLinks  `json:"links"`
	Data  []PostData `json:"data"`
}

type PostLinks struct {
	Self string `json:"self"`
}

type PostData struct {
	Type       DataType       `json:"type"`
	ID         uuid.UUID      `json:"id"`
	Attributes PostAttributes `json:"attributes"`
}

type PostAttributes struct {
	UserID  uuid.UUID `json:"userId"`
	Text    string    `json:"text"`
	Privacy Privacy   `json:"privacy"`
}

func (pr *PostResponse) Send(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(pr)
}
