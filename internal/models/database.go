package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Verified bool   `gorm:"default:false"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		Data: []struct {
			ID string `json:"id"`
		}{
			{
				ID: u.ID.String(),
			},
		},
	}
}
