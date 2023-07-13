package response

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type PostResponse struct {
	Data []*PostData `json:"data"`
}

type PostData struct {
	Type          shared.DataType `json:"type"`
	ID            uuid.UUID       `json:"id"`
	Attributes    PostAttr        `json:"attributes"`
	Links         PostLinks       `json:"links"`
	Relationships PostRel         `json:"relationships"`
}

type PostAttr struct {
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Text         string    `json:"text"`
	Privacy      string    `json:"privacy"`
	MediaIDs     []string  `json:"media_ids"`
	IsLiked      bool      `json:"isLiked"`
	IsBookmarked bool      `json:"isBookmarked"`
}

type PostLinks struct {
	Self string `json:"self"`
}

type PostRel struct {
	User UserResponse `json:"user"`
	// Media here
}

func CreatePostsResponse(posts []*database.Post) *PostResponse {
	return nil
}
