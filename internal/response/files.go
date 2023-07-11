package response

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type FileResponse struct {
	Data []*FileData `json:"data"`
}

type FileData struct {
	Type       shared.DataType `json:"type"`
	ID         uuid.UUID       `json:"id"`
	Attributes FileAttr        `json:"attributes"`
	Links      FileLinks       `json:"links"`
}

type FileAttr struct {
	CreatedAt time.Time `json:"createdAt"`
	OwnerID   uuid.UUID `json:"ownerId"`
	Extension string    `json:"extension"`
}

type FileLinks struct {
	Self string `json:"self"`
}

func CreateFileResponse(files []*database.File) *FileResponse {
	// Initialize filesData as an empty slice
	// This is so its not null in the JSON response
	// because javascript is stupid
	filesData := make([]*FileData, 0)

	for _, file := range files {
		selfLink := baseURL + "/v1/files/" + file.ID.String()

		filesData = append(filesData, &FileData{
			Type: shared.DataTypeFile,
			ID:   file.ID,
			Attributes: FileAttr{
				CreatedAt: file.CreatedAt,
				OwnerID:   file.OwnerID,
				Extension: file.Extension,
			},
			Links: FileLinks{
				Self: selfLink,
			},
		})
	}
	return &FileResponse{
		Data: filesData,
	}
}
