package response

import "github.com/bwoff11/frens/internal/database"

type BookmarkResp struct {
	Links    BookmarkResp_Links      `json:"links,omitempty"`
	Data     []BookmarkResp_Data     `json:"data,omitempty"`
	Included []BookmarkResp_Included `json:"included,omitempty"`
}

type BookmarkResp_Links struct {
	Self string `json:"self"`
}

type BookmarkResp_Data struct {
	Type       string                      `json:"type"`
	ID         string                      `json:"id"`
	Attributes BookmarkResp_DataAttributes `json:"attributes"`
}

type BookmarkResp_DataAttributes struct {
	// For future use
}

type BookmarkResp_Included struct {
	// For future use
}

func GenerateBookmarkResponse(bookmark *database.Bookmark) *BookmarkResp {
	return &BookmarkResp{
		Links: BookmarkResp_Links{
			Self: "/bookmarks",
		},
		Data: []BookmarkResp_Data{
			{
				Type:       "bookmark",
				ID:         "1",
				Attributes: BookmarkResp_DataAttributes{
					// For future use
				},
			},
		},
	}
}

func GenerateBookmarksResponse(bookmarks []*database.Bookmark) *BookmarkResp {
	return &BookmarkResp{
		Links: BookmarkResp_Links{
			Self: "/bookmarks",
		},
		Data: []BookmarkResp_Data{
			{
				Type:       "bookmark",
				ID:         "1",
				Attributes: BookmarkResp_DataAttributes{
					// For future use
				},
			},
			{
				Type:       "bookmark",
				ID:         "2",
				Attributes: BookmarkResp_DataAttributes{
					// For future use
				},
			},
		},
	}
}
