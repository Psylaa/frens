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
	var likeData []*LikeData

	for _, like := range likes {

		likeData = append(likeData, &LikeData{
			Type: shared.DataTypeLike,
			ID:   like.ID.String(),
			Attributes: LikeAttr{
				CreatedAt: like.CreatedAt.String(),
				UserID:    like.UserID,
				PostID:    like.PostID,
			},
		})
	}

	return &LikeResponse{
		Data: likeData,
	}
}
