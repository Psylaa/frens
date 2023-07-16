package models

import "github.com/google/uuid"

type LikeResponse struct {
	Links LikeLinks  `json:"links"`
	Data  []LikeData `json:"data"`
}

type LikeLinks struct {
	Self string `json:"self"`
}

type LikeData struct {
	Type          DataType       `json:"type"`
	ID            uuid.UUID      `json:"id"`
	Attributes    LikeAttributes `json:"attributes"`
	Relationships Relationship   `json:"relationships"`
}

type LikeAttributes struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
