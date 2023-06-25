package database

import "github.com/google/uuid"

// Media represents a media item attached to a status update.
type Media struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	StatusID uuid.UUID `json:"statusId"`
	Type     string    `json:"type"`
	URL      string    `json:"url"`
}
