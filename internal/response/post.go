package response

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type PostResp struct {
	Links    PostResp_Links      `json:"links,omitempty"`
	Data     []PostResp_Data     `json:"data,omitempty"`
	Included []PostResp_Included `json:"included,omitempty"`
}

type PostResp_Links struct {
	Self   string `json:"self"`
	Author string `json:"author"`
}

type PostResp_Data struct {
	Type          string                     `json:"type"`
	ID            uuid.UUID                  `json:"id,omitempty"`
	Attributes    PostResp_DataAttributes    `json:"attributes"`
	Relationships PostResp_DataRelationships `json:"relationships,omitempty"`
}

type PostResp_DataAttributes struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Privacy   shared.Privacy `json:"privacy"`
	Text      string         `json:"text"`
}

type PostResp_DataRelationships struct {
	Author PostResp_DataRelationshipsAuthor `json:"author"`
}

type PostResp_DataRelationshipsAuthor struct {
	Data PostResp_DataRelationshipsAuthorData `json:"data"`
}

type PostResp_DataRelationshipsAuthorData struct {
	Type string    `json:"type"`
	ID   uuid.UUID `json:"id"`
}

type PostResp_Included struct {
}

func GeneratePostResponse(post *database.Post) *PostResp {
	selfLink := fmt.Sprintf("/posts/%s", post.ID)
	authorLink := fmt.Sprintf("/users/%s", post.AuthorID)

	return &PostResp{
		Links: PostResp_Links{
			Self:   selfLink,
			Author: authorLink,
		},
		Data: []PostResp_Data{
			{
				Type: "post",
				ID:   post.ID,
				Attributes: PostResp_DataAttributes{
					CreatedAt: post.CreatedAt,
					UpdatedAt: post.UpdatedAt,
					Privacy:   post.Privacy,
					Text:      post.Text,
				},
				Relationships: PostResp_DataRelationships{
					Author: PostResp_DataRelationshipsAuthor{
						Data: PostResp_DataRelationshipsAuthorData{
							Type: "user",
							ID:   post.AuthorID,
						},
					},
				},
			},
		},
	}
}

func GeneratePostsResponse(posts []*database.Post) *PostResp {
	postsResp := PostResp{
		Data: make([]PostResp_Data, 0, len(posts)),
	}

	for _, post := range posts {
		postResp := GeneratePostResponse(post)
		postsResp.Data = append(postsResp.Data, postResp.Data[0])
	}

	return &postsResp
}
