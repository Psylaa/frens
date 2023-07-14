package models

import (
	"github.com/go-playground/validator"
)

type Privacy string

const (
	PrivacyPublic    Privacy = "public"
	PrivacyProtected Privacy = "protected"
	PrivacyPrivate   Privacy = "private"
)

type DataType string

const (
	DataTypeUser     DataType = "user"
	DataTypePost     DataType = "post"
	DataTypeFollow   DataType = "follow"
	DataTypeToken    DataType = "token"
	DataTypeBookmark DataType = "bookmark"
	DataTypeLike     DataType = "like"
	DataTypeFile     DataType = "file"
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
