package response

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type BookmarkResponse struct {
	Data []*BookmarkData `json:"data"`
}

type BookmarkData struct {
	Type       shared.DataType `json:"type"`
	ID         uuid.UUID       `json:"id"`
	Attributes BookmarkAttr    `json:"attributes"`
	Links      BookmarkLinks   `json:"links"`
}

type BookmarkAttr struct {
	CreatedAt time.Time `json:"createdAt"`
	UserID    uuid.UUID `json:"userId"`
	PostID    uuid.UUID `json:"postId"`
}

type BookmarkLinks struct {
	Self string `json:"self"`
}

func CreateBookmarksResponse(bookmarks []*database.Bookmark) *BookmarkResponse {
	return nil
}
