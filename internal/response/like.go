package response

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type LikeResponse struct {
	Data []*LikeData `json:"data"`
}

type LikeData struct {
	Type       shared.DataType `json:"type"`
	ID         string          `json:"id"`
	Attributes LikeAttr        `json:"attributes"`
}

type LikeAttr struct {
	CreatedAt string     `json:"createdAt"`
	UserID    *uuid.UUID `json:"userId"`
	PostID    *uuid.UUID `json:"postId"`
}

func CreateLikesResponse(likes []*database.Like) *LikeResponse {
	return nil
}
