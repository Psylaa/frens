package models

import (
	"github.com/google/uuid"
)

type Post struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"` // ID of the user who created the post
	User   User      `json:"user"`                              // User who created the post
	Text   string    `gorm:"type:text" json:"text"`             // Text content of the post
}
