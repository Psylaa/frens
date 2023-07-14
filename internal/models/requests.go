package models

import (
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr *RegisterRequest) Sanitize() {
	p := bluemonday.UGCPolicy()
	rr.Username = p.Sanitize(rr.Username)
	rr.Email = p.Sanitize(rr.Email)
	// Password is not sanitized as it will be hashed and we don't want to unintentionally alter it
}

func (rr *RegisterRequest) ToUser() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rr.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
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
