package models

import (
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (lr *LoginRequest) Sanitize() {
	p := bluemonday.UGCPolicy()
	lr.Email = p.Sanitize(lr.Email)
}

func (lr *LoginRequest) Validate() error {
	return ValidateStruct(lr)
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (rr *RegisterRequest) Sanitize() {
	p := bluemonday.UGCPolicy()
	rr.Username = p.Sanitize(rr.Username)
	rr.Email = p.Sanitize(rr.Email)
}

func (rr *RegisterRequest) Validate() error {
	return ValidateStruct(rr)
}

func (rr *RegisterRequest) ToUser() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rr.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Role:     RoleUser,
		Username: rr.Username,
		Email:    rr.Email,
		Password: string(hashedPassword),
	}, nil
}

type UpdateUserRequest struct {
	Bio      *string `json:"bio"`
	AvatarID *string `json:"avatar_id"`
	CoverID  *string `json:"cover_id"`
}
