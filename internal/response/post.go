package response

import (
	"fmt"
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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Text      string    `json:"text"`
	Privacy   string    `json:"privacy"`
	MediaIDs  []string  `json:"media_ids"`
}

type PostLinks struct {
	Self string `json:"self"`
}

type PostRel struct {
	Author UserResponse `json:"author"`
}

func CreatePostsResponse(posts []*database.Post) *PostResponse {
	var postData []*PostData

	for _, post := range posts {
		selfLink := fmt.Sprintf("%s/posts/%s", baseURL, post.ID.String())

		// Extract media IDs
		var mediaIDs []string
		for _, media := range post.Media {
			mediaIDs = append(mediaIDs, media.ID.String())
		}

		postData = append(postData, &PostData{
			Type: shared.DataTypePost,
			ID:   post.ID,
			Attributes: PostAttr{
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
				Privacy:   string(post.Privacy),
				Text:      post.Text,
				MediaIDs:  mediaIDs,
			},
			Links: PostLinks{
				Self: selfLink,
			},
			Relationships: PostRel{
				Author: *CreateUsersResponse([]*database.User{&post.Author}),
			},
		})
	}

	return &PostResponse{
		Data: postData,
	}
}
