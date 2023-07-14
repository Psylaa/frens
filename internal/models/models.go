package models

import (
	"github.com/go-playground/validator"
)

type DataTypes string

type Role string

const (
	DataTypeUser DataTypes = "users"

	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
