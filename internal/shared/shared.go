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

func HashPassword(password string, salt []byte) (*string, error) {
	// Generate the bcrypt hash with salt
	hashedPassword, err := bcrypt.GenerateFromPassword(append([]byte(password), salt...), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Convert the hashed password to a string
	hashedPasswordString := string(hashedPassword)

	return &hashedPasswordString, nil
}
