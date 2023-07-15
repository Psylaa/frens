package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	Role     Role   `gorm:"default:user"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Bio      string `gorm:"default:''"`
	Verified bool   `gorm:"default:false"`
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (u *User) SetBio(bio string) {
	u.Bio = bio
}

func (u *User) ToResponse() *UserRespone {
	return &UserRespone{
		Links: UserLinks{
			Self: "todo",
		},
		Data: []UserData{
			{
				Type: DataTypeUser,
				ID:   u.ID,
				Attributes: UserAttributes{
					Role:      u.Role,
					Username:  u.Username,
					Bio:       u.Bio,
					Verrified: u.Verified,
				},
			},
		},
	}
}

type Post struct {
	BaseModel
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	User    User      `gorm:"foreignKey:UserID;references:ID"`
	Text    string    `gorm:"type:text"`
	Privacy Privacy   `gorm:""`
	Media   []Media
}

func (p *Post) ToResponse() *PostResponse {
	return &PostResponse{
		Links: PostLinks{
			Self: "todo",
		},
		Data: []PostData{
			{
				Type: DataTypePost,
				ID:   p.ID,
				Attributes: PostAttributes{
					UserID:  p.UserID,
					Text:    p.Text,
					Privacy: p.Privacy,
				},
			},
		},
	}
}

type Follow struct {
	BaseModel
	SourceUserID uuid.UUID `gorm:"type:uuid;not null"`
	TargetUserID uuid.UUID `gorm:"type:uuid;not null"`
}

type Like struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	PostID uuid.UUID `gorm:"type:uuid;not null"`
}

type Bookmark struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	PostID uuid.UUID `gorm:"type:uuid;not null"`
}

type Block struct {
	BaseModel
	SourceUserID uuid.UUID `gorm:"type:uuid;not null"`
	TargetUserID uuid.UUID `gorm:"type:uuid;not null"`
}

type Media struct {
	BaseModel
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	User      User       `gorm:"foreignKey:UserID;references:ID"`
	PostID    *uuid.UUID `gorm:"type:uuid"`
	Post      Post       `gorm:"foreignKey:PostID;references:ID"`
	Extension string     `gorm:"not null"`
}

func (m *Media) ToResponse() *MediaResponse {
	return &MediaResponse{
		Links: MediaLinks{
			Self: "todo",
		},
		Data: []MediaData{
			{
				Type: DataTypeMedia,
				ID:   m.ID,
				Attributes: MediaAttributes{
					UserID:    m.UserID,
					PostID:    m.PostID,
					Extension: m.Extension,
				},
			},
		},
	}
}
