package models

import (
	"time"

	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Text     string      `json:"text" validate:"max=2048"`
	MediaIDs []uuid.UUID `json:"media_ids"`
}

func (req *CreatePostRequest) Validate() error {
	return ValidateStruct(req)
}

func (req *CreatePostRequest) ToPost(requestorID *uuid.UUID) (*Post, error) {
	return &Post{
		UserID: *requestorID,
		Text:   req.Text,
	}, nil
}

type FeedRequest struct {
	Count  *int       `json:"count" validate:"min=1,max=100"`
	Cursor *time.Time `json:"cursor"`
}
