package response

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type DataType string

const (
	UserType     DataType = "user"
	BookmarkType DataType = "bookmark"
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
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	ExpiresAt *time.Time     `json:"expiresAt,omitempty"`
	Extenion  *string        `json:"extension,omitempty"`
	Privacy   shared.Privacy `json:"privacy,omitempty"`
	Text      string         `json:"text,omitempty"`
	Token     string         `json:"token,omitempty"`
	Username  string         `json:"username,omitempty"`
}

type Relationships struct {
	Author *Response `json:"author,omitempty"` // Recursive types must be pointers
	Owner  *Response `json:"owner,omitempty"`  // Recursive types must be pointers
	User   *Response `json:"user,omitempty"`   // Recursive types must be pointers
	Media  *Response `json:"media,omitempty"`  // Recursive types must be pointers
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

func CreateErrorResponse(err RespErr) *Response {
	return &Response{
		Errors: []*RespErr{&err},
	}
}

func CreateCountResponse(count int) *Response {
	return &Response{
		Meta: &Meta{
			Count: count,
		},
	}
}

func CreateUsersResponse(users []*database.User) *Response {
	var data []Data
	for _, user := range users {
		selfLink := fmt.Sprintf("%s/users/%s", baseURL, user.ID)
		postsLink := fmt.Sprintf("%s/users/%s/posts", baseURL, user.ID)
		followersLink := fmt.Sprintf("%s/users/%s/followers", baseURL, user.ID)
		followingLink := fmt.Sprintf("%s/users/%s/following", baseURL, user.ID)

		var avatarLink string
		if user.AvatarID == uuid.Nil {
			avatarLink = fmt.Sprintf("%s/files/default-avatar.png", baseURL)
		} else {
			avatarLink = fmt.Sprintf("%s/files/%s%s", baseURL, user.AvatarID, user.Avatar.Extension)
		}

		var coverLink string
		if user.CoverID == uuid.Nil {
			coverLink = fmt.Sprintf("%s/files/default-cover.png", baseURL)
		} else {
			coverLink = fmt.Sprintf("%s/files/%s%s", baseURL, user.CoverID, user.Cover.Extension)
		}

		data = append(data, Data{
			Type: UserType,
			ID:   user.ID,
			Attributes: Attributes{
				CreatedAt: user.CreatedAt,
				UpdatedAt: &user.UpdatedAt,
				Username:  user.Username,
			},
			Links: Links{
				Self:      selfLink,
				Posts:     postsLink,
				Following: followingLink,
				Followers: followersLink,
				Avatar:    avatarLink,
				Cover:     coverLink,
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
			Relationships: &Relationships{},
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

func CreatePostsResponse(posts []*database.Post) *Response {
	var data []Data
	for _, post := range posts {
		selfLink := fmt.Sprintf("%s/posts/%s", baseURL, post.ID)
		//authorLink := fmt.Sprintf("%s/users/%s", baseURL, post.Author.ID)

		data = append(data, Data{
			Type: UserType,
			ID:   post.ID,
			Attributes: Attributes{
				CreatedAt: post.CreatedAt,
				UpdatedAt: &post.UpdatedAt,
				Privacy:   post.Privacy,
				Text:      post.Text,
			},
			Relationships: &Relationships{
				Author: CreateUsersResponse([]*database.User{&post.Author}),
				Media:  CreateFilesResponse(post.Media),
			},
			Links: Links{
				Self: selfLink,
			},
		})
	}

	return &Response{
		Data: data,
	}
}

func CreateFilesResponse(files []*database.File) *Response {
	var data []Data
	for _, file := range files {
		selfLink := fmt.Sprintf("%s/files/%s%s", baseURL, file.ID, file.Extension)

		data = append(data, Data{
			Type: "file",
			ID:   file.ID,
			Attributes: Attributes{
				CreatedAt: file.CreatedAt,
				UpdatedAt: &file.UpdatedAt,
				Extenion:  &file.Extension,
			},
			Links: Links{
				Self: selfLink,
			},
		})
	}

	return &Response{
		Data: data,
	}
}

func CreateLoginResponse(user *database.User, token string, expirationDate time.Time) *Response {
	return &Response{
		Data: []Data{
			{
				Type: "login",
				ID:   user.ID,
				Attributes: Attributes{
					CreatedAt: time.Now(),
					ExpiresAt: &expirationDate,
					Token:     token,
					Username:  user.Username,
				},
				Links: Links{
					Self: fmt.Sprintf("%s/login", baseURL),
				},
			},
		},
	}
}
