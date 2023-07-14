package models

import (
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (rr *RegisterRequest) Sanitize() {
	p := bluemonday.UGCPolicy()
	rr.Username = p.Sanitize(rr.Username)
	rr.Email = p.Sanitize(rr.Email)
	// Password is not sanitized as it will be hashed and we don't want to unintentionally alter it
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
