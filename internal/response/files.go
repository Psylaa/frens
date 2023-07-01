package response

import (
	"fmt"

	"github.com/bwoff11/frens/internal/database"
)

type FileResp struct {
	Links    FileResp_Links      `json:"links,omitempty"`
	Data     []FileResp_Data     `json:"data,omitempty"`
	Included []FileResp_Included `json:"included,omitempty"`
}

type FileResp_Links struct {
	Self string `json:"self"`
}

type FileResp_Data struct {
	Type       string                  `json:"type"`
	ID         string                  `json:"id"`
	Attributes FileResp_DataAttributes `json:"attributes"`
}

type FileResp_DataAttributes struct {
	Extension string `json:"extension"`
}

type FileResp_Included struct {
	// For future use
}

func GenerateFileResponse(file *database.File) *FileResp {
	selfLink := fmt.Sprintf("%s/files/%s%s", baseURL, file.ID, file.Extension)

	return &FileResp{
		Links: FileResp_Links{
			Self: selfLink,
		},
		Data: []FileResp_Data{
			{
				Type: "file",
				ID:   file.ID.String(),
				Attributes: FileResp_DataAttributes{
					Extension: file.Extension,
				},
			},
		},
	}
}

func GeneratesFilesResponse(files []*database.File) *FileResp {
	var data []FileResp_Data

	for _, file := range files {
		data = append(data, FileResp_Data{
			Type: "file",
			ID:   file.ID.String(),
			Attributes: FileResp_DataAttributes{
				Extension: file.Extension,
			},
		})
	}

	return &FileResp{
		Data: data,
	}
}
