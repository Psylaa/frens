package response

import (
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type DataType string

const (
	UserType     DataType = "user"
	BookmarkType DataType = "bookmark"
	LikeType     DataType = "like"
)

type Response struct {
	Data     []Data      `json:"data,omitempty"` // Can be an array or a single data object
	Errors   []*RespErr  `json:"errors,omitempty"`
	Included []*Response `json:"included,omitempty"` // Recursive types must be pointers
	Meta     *Meta       `json:"meta,omitempty"`
}

type Links struct {
	Self      string `json:"self"`
	Posts     string `json:"posts,omitempty"`
	Following string `json:"following,omitempty"`
	Followers string `json:"followers,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Cover     string `json:"cover,omitempty"`
	Owner     string `json:"owner,omitempty"`
}

type Data struct {
	Type          DataType       `json:"type"`
	ID            uuid.UUID      `json:"id"`
	Attributes    Attributes     `json:"attributes"`
	Relationships *Relationships `json:"relationships,omitempty"`
	Links         Links          `json:"links"`
}

type Attributes struct {
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	ExpiresAt    *time.Time     `json:"expiresAt,omitempty"`
	Extenion     *string        `json:"extension,omitempty"`
	Privacy      shared.Privacy `json:"privacy,omitempty"`
	Text         string         `json:"text,omitempty"`
	Token        string         `json:"token,omitempty"`
	Username     string         `json:"username,omitempty"`
	IsLiked      *bool          `json:"isLiked,omitempty"`      // Pointer so its not ommited if false
	IsBookmarked *bool          `json:"isBookmarked,omitempty"` // Pointer so its not ommited if false
	IsFollowing  *bool          `json:"isFollowing,omitempty"`  // Pointer so its not ommited if false
	AvatarID     *uuid.UUID     `json:"avatarID,omitempty"`
	CoverID      *uuid.UUID     `json:"coverID,omitempty"`
	Bio          *string        `json:"bio,omitempty"`
}

type Relationships struct {
	Author *Response `json:"author,omitempty"` // Recursive types must be pointers
	Owner  *Response `json:"owner,omitempty"`  // Recursive types must be pointers
	User   *Response `json:"user,omitempty"`   // Recursive types must be pointers
	Media  *Response `json:"media,omitempty"`  // Recursive types must be pointers
	Post   *Response `json:"post,omitempty"`   // Recursive types must be pointers
}

type Error struct {
}

type Meta struct {
	Count int `json:"count,omitempty"`
}

var baseURL string
var defaultBio string

func Init(config *config.Config) {
	baseURL = config.Server.BaseURL
	defaultBio = config.Users.DefaultBio
}

func CreateCountResponse(count int) *Response {
	return &Response{
		Meta: &Meta{
			Count: count,
		},
	}
}
