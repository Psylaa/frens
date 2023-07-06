package shared

import "golang.org/x/crypto/bcrypt"

type Privacy string

const (
	PrivacyPublic  Privacy = "public"
	PrivacyPrivate Privacy = "private"
)

type DataType string

const (
	DataTypeUser   DataType = "user"
	DataTypePost   DataType = "post"
	DataTypeFollow DataType = "follow"
	DataTypeToken  DataType = "token"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func HashPassword(password string) (*string, error) {
	// Generate the bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Convert the hashed password to a string
	hashedPasswordString := string(hashedPassword)

	return &hashedPasswordString, nil
}
