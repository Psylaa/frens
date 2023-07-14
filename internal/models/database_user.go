package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Role     Role   `gorm:"default:user"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Bio      string `gorm:"default:''"`
	Verified bool   `gorm:"default:false"`
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

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (u *User) SetBio(bio string) {
	u.Bio = bio
}
