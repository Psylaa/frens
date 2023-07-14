package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DataTypes string

const (
	Users DataTypes = "users"
)

type Response struct {
	Data []Data `json:"data"`
}

type Data struct {
	Type          DataTypes     `json:"type"`
	ID            uuid.UUID     `json:"id"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

type Attributes struct {
}

type Relationships struct {
}

func (r *Response) Send(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(r)
}
