package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostResponse struct {
	Links    PostLinks  `json:"links"`
	Data     []PostData `json:"data"`
	Included []UserData `json:"included"`
}

type PostLinks struct {
	Self string `json:"self"`
}

type PostData struct {
	Type          DataType       `json:"type"`
	ID            uuid.UUID      `json:"id"`
	Attributes    PostAttributes `json:"attributes"`
	Relationships Relationship   `json:"relationships"`
}

type PostAttributes struct {
	Text string `json:"text"`
}

type Relationship struct {
	User RelationshipData `json:"user"`
}

type RelationshipData struct {
	Data RelationshipDetails `json:"data"`
}

type RelationshipDetails struct {
	Type string    `json:"type"`
	ID   uuid.UUID `json:"id"`
}

func (pr *PostResponse) Send(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(pr)
}
