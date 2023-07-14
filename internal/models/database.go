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

func (base *BaseModel) BeforeCreate() (err error) {
	base.ID = uuid.New()
	return
}

type User struct {
	BaseModel
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Bio      string
	Verified bool `gorm:"default:false"`
}

func (u *User) ToResponse() *Response {
	return &Response{
		Data: []Data{
			{
				Type: Users,
				ID:   u.ID,
			},
		},
	}
}
