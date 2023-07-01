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
}

type Links struct {
	Self           string `json:"self"`
	Posts          string `json:"posts,omitempty"`
	Following      string `json:"following,omitempty"`
	Followers      string `json:"followers,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
	CoverImage     string `json:"coverImage,omitempty"`
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
}

var baseURL string

func Init(baseUrl string) {
	baseURL = baseUrl
}

func GenerateErrorResponse(err APIResponseErr) *ErrResp {
	return &ErrResp{
		Error: err,
	}
}

func CreateUserResponse(user *database.User) *Response {
	selfLink := fmt.Sprintf("%s/users/%s", baseURL, user.ID)
	postsLink := fmt.Sprintf("%s/users/%s/posts", baseURL, user.ID)
	followersLink := fmt.Sprintf("%s/users/%s/followers", baseURL, user.ID)
	followingLink := fmt.Sprintf("%s/users/%s/following", baseURL, user.ID)
	ppLink := fmt.Sprintf("%s/files/%s%s", baseURL, user.ProfilePicture.ID, user.ProfilePicture.Extension)
	ciLink := fmt.Sprintf("%s/files/%s%s", baseURL, user.CoverImage.ID, user.CoverImage.Extension)

	return &Response{
		Data: Data{
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
	}
}

func CreateBookmarkResponse(bookmark *database.Bookmark) *Response {
	return &Response{
		Data: Data{
			Type:       BookmarkType,
			ID:         bookmark.ID,
			Attributes: Attributes{},
			Relationships: Relationships{
				Owner: &Response{
					//Data: []Data{CreateUserResponse(&bookmark.Owner)},
				},
			},
			Links: Links{
				Self: baseURL + "/bookmarks/" + bookmark.ID.String(),
			},
		},
	}
}
