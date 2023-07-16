package models

import "github.com/google/uuid"

type CreateLikeRequest struct {
	PostID uuid.UUID `json:"post_id" validate:"required"`
}

func (clr *CreateLikeRequest) Validate() error {
	return ValidateStruct(clr)
}

func (clr *CreateLikeRequest) ToLike(requestorID *uuid.UUID) (*Like, error) {
	return &Like{
		UserID: *requestorID,
		PostID: clr.PostID,
	}, nil
}
