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
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	Text         string `json:"text"`
	IsLiked      bool   `json:"isLiked"`
	IsBookmarked bool   `json:"isBookmarked"`
}

func (pr *PostResponse) Send(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(pr)
}
