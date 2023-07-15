package models

import (
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

type CreatePostRequest struct {
	Text     string      `json:"text" validate:"max=2048"`
	Privacy  Privacy     `json:"privacy" validate:"oneof=private public protected"`
	MediaIDs []uuid.UUID `json:"media_ids"`
}

func (req *CreatePostRequest) Sanitize() {
	p := bluemonday.UGCPolicy()
	req.Text = p.Sanitize(req.Text)
}

func (req *CreatePostRequest) Validate() error {
	return ValidateStruct(req)
}

func (req *CreatePostRequest) ToPost(requestorID *uuid.UUID) (*Post, error) {
	return &Post{
		UserID:  *requestorID,
		Privacy: req.Privacy,
		Text:    req.Text,
	}, nil
}
