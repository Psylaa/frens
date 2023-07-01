package response

import (
	"fmt"

	"github.com/bwoff11/frens/internal/database"
	"github.com/google/uuid"
)

type DataType string

const (
	UserType     DataType = "user"
	BookmarkType DataType = "bookmark"
)

type Response struct {
	Data     interface{} `json:"data,omitempty"` // Can be an array or a single data object
	Errors   []Error     `json:"errors,omitempty"`
	Included []*Response `json:"included,omitempty"` // Recursive types must be pointers
	Meta     Meta        `json:"meta,omitempty"`
}

type Links struct {
	Self           string `json:"self"`
	Posts          string `json:"posts,omitempty"`
	Following      string `json:"following,omitempty"`
	Followers      string `json:"followers,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
	CoverImage     string `json:"coverImage,omitempty"`
	Owner          string `json:"owner,omitempty"`
}

type Data struct {
	Type          DataType      `json:"type"`
	ID            uuid.UUID     `json:"id"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
	Links         Links         `json:"links"`
}

type Attributes struct {
}

type Relationships struct {
	Author *Response `json:"author,omitempty"` // Recursive types must be pointers
	Owner  *Response `json:"owner,omitempty"`  // Recursive types must be pointers
	User   *Response `json:"user,omitempty"`   // Recursive types must be pointers
}

type Error struct {
}

type Meta struct {
	Count int `json:"count,omitempty"`
}

var baseURL string

func Init(baseUrl string) {
	baseURL = baseUrl
}

func CreateErrorResponse(err APIResponseErr) *ErrResp {
	return &ErrResp{
		Error: err,
	}
}

func CreateCountResponse(count int) *Response {
	return &Response{
		Meta: Meta{
			Count: count,
		},
	}
}

func CreateUserResponse(users []*database.User) *Response {
	var data []Data
	for _, user := range users {
		selfLink := fmt.Sprintf("%s/users/%s", baseURL, user.ID)
		postsLink := fmt.Sprintf("%s/users/%s/posts", baseURL, user.ID)
		followersLink := fmt.Sprintf("%s/users/%s/followers", baseURL, user.ID)
		followingLink := fmt.Sprintf("%s/users/%s/following", baseURL, user.ID)
		ppLink := fmt.Sprintf("%s/files/%s%s", baseURL, user.ProfilePicture.ID, user.ProfilePicture.Extension)
		ciLink := fmt.Sprintf("%s/files/%s%s", baseURL, user.CoverImage.ID, user.CoverImage.Extension)

		data = append(data, Data{
			Type:       UserType,
			ID:         user.ID,
			Attributes: Attributes{},
			Links: Links{
				Self:           selfLink,
				Posts:          postsLink,
				Following:      followingLink,
				Followers:      followersLink,
				ProfilePicture: ppLink,
				CoverImage:     ciLink,
			},
		},
		)
	}

	return &Response{
		Data: data,
	}
}

func CreateBookmarkResponse(bookmark []*database.Bookmark) *Response {
	var data []Data
	for _, b := range bookmark {
		selfLink := fmt.Sprintf("%s/bookmarks/%s", baseURL, b.ID)
		ownerLink := fmt.Sprintf("%s/users/%s", baseURL, b.Owner.ID)

		data = append(data, Data{
			Type:          BookmarkType,
			ID:            b.ID,
			Attributes:    Attributes{},
			Relationships: Relationships{},
			Links: Links{
				Self:  selfLink,
				Owner: ownerLink,
			},
		})
	}

	return &Response{
		Data: data,
	}
}
