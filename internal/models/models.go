package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/bwoff11/frens/internal/logger"
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

// Scan implements the Scanner interface.
func (p *Privacy) Scan(value interface{}) error {
	if value == nil {
		*p = PrivacyPrivate
		return nil
	}
	switch value.(type) {
	case []byte:
		switch string(value.([]byte)) {
		case "public":
			*p = PrivacyPublic
		case "protected":
			*p = PrivacyProtected
		case "private":
			*p = PrivacyPrivate
		default:
			return fmt.Errorf("invalid privacy value: %s", value)
		}
	default:
		return fmt.Errorf("invalid privacy type: %T", value)
	}
	return nil
}

// Value implements the Valuer interface.
func (p Privacy) Value() (driver.Value, error) {
	return p.ToString(), nil
}

// ToString returns the string representation of the Privacy type.
func (p Privacy) ToString() string {
	switch p {
	case PrivacyPublic:
		return "public"
	case PrivacyProtected:
		return "protected"
	case PrivacyPrivate:
		return "private"
	default:
		logger.Error(logger.LogMessage{
			Package:  "models",
			Function: "Privacy.ToString",
			Message:  "Invalid privacy type provided (somehow). Validator should have prevented this. Unable to convert to string. Defaulting to private.",
		}, nil)
		return "private"
	}
}
