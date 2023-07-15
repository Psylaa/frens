package models

import (
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

type CreatePostRequest struct {
	Text    string  `json:"text"`
	Privacy Privacy `json:"privacy"`
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
