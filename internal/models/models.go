package models

import (
	"github.com/go-playground/validator"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
