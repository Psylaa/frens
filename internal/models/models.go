package models

import (
	"github.com/go-playground/validator"
)

type DataType string

const (
	DataTypeBlock    DataType = "block"
	DataTypeBookmark DataType = "bookmark"
	DataTypeFollow   DataType = "follow"
	DataTypeLike     DataType = "like"
	DataTypeMedia    DataType = "media"
	DataTypePost     DataType = "post"
	DataTypeToken    DataType = "token"
	DataTypeUser     DataType = "user"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
