package models

import (
	"github.com/bwoff11/frens/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr *RegisterRequest) ToUser() (*database.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rr.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &database.User{
		Username: rr.Username,
		Email:    rr.Email,
		Password: string(hashedPassword),
	}, nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Bio      *string `json:"bio"`
	AvatarID *string `json:"avatar_id"`
	CoverID  *string `json:"cover_id"`
}
