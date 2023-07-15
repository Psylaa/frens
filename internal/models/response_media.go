package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MediaResponse struct {
	Links MediaLinks  `json:"links"`
	Data  []MediaData `json:"data"`
}

type MediaLinks struct {
	Self string `json:"self"`
}

type MediaData struct {
	Type       DataType        `json:"type"`
	ID         uuid.UUID       `json:"id"`
	Attributes MediaAttributes `json:"attributes"`
}

type MediaAttributes struct {
	UserID    uuid.UUID  `json:"userId"`
	PostID    *uuid.UUID `json:"postId,omitempty"`
	Extension string     `json:"extension"`
}

func (mr *MediaResponse) Send(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(mr)
}
