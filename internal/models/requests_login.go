package models

import "github.com/microcosm-cc/bluemonday"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (lr *LoginRequest) Sanitize() {
	p := bluemonday.UGCPolicy()
	lr.Email = p.Sanitize(lr.Email)
	// Password is not sanitized as it will be hashed and we don't want to unintentionally alter it
}

func (lr *LoginRequest) Validate() error {
	return ValidateStruct(lr)
}
