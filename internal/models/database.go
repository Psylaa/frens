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

type InteractorModel struct {
	BaseModel
	SourceID uuid.UUID `gorm:"type:uuid;not null"`
	TargetID uuid.UUID `gorm:"type:uuid;not null"`
}

func (interactor *InteractorModel) BeforeCreate() (err error) {
	interactor.ID = uuid.New()
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
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"` // ID of the user who created the post
	User   User      `json:"user"`                              // User who created the post
	Text   string    `gorm:"type:text" json:"text"`             // Text content of the post
}

type Follow struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target User `gorm:"foreignKey:UserID"`
}

type Like struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target Post `gorm:"foreignKey:PostID"`
}

type Bookmark struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target Post `gorm:"foreignKey:PostID"`
}

type Block struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target User `gorm:"foreignKey:UserID"`
}
