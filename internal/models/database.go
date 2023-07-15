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
	Role     Role   `gorm:"default:user" json:"-"`
	Username string `gorm:"unique" json:"-"`
	Email    string `gorm:"unique" json:"-"`
	Password string `gorm:"not null" json:"-"`
	Bio      string `gorm:"default:''" json:"-"`
	Verified bool   `gorm:"default:false" json:"-"`
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (u *User) ToResponseData() UserData {
	return UserData{
		Type: DataTypeUser,
		ID:   u.ID,
		Attributes: UserAttributes{
			Role:     u.Role,
			Username: u.Username,
			Bio:      u.Bio,
			//Verified: u.Verified,
		},
	}
}

type Post struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	User   User      `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Text   string    `gorm:"type:text" json:"-"`
	Media  []Media   `json:"-"`
}

func (p *Post) ToResponseData() (PostData, UserData) {
	postData := PostData{
		Type: DataTypePost,
		ID:   p.ID,
		Attributes: PostAttributes{
			Text: p.Text,
		},
		Relationships: Relationship{
			User: RelationshipData{
				Data: RelationshipDetails{
					Type: "user",
					ID:   p.UserID,
				},
			},
		},
	}

	userData := p.User.ToResponseData()

	return postData, userData
}

type Follow struct {
	BaseModel
	SourceUserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	TargetUserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
}

type Like struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	PostID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
}

type Bookmark struct {
	BaseModel
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	PostID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
}

type Block struct {
	BaseModel
	SourceUserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	TargetUserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
}

type Media struct {
	BaseModel
	UserID    uuid.UUID  `gorm:"type:uuid;not null" json:"-"`
	User      User       `gorm:"foreignKey:UserID;references:ID" json:"-"`
	PostID    *uuid.UUID `gorm:"type:uuid" json:"-"`
	Post      Post       `gorm:"foreignKey:PostID;references:ID" json:"-"`
	Extension string     `gorm:"not null" json:"-"`
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
