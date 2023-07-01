package shared

import "github.com/google/uuid"

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

type APIResponseErr string

const (
	ErrInternal      APIResponseErr = "internal server error"
	ErrNotFound      APIResponseErr = "not found"
	ErrInvalidID     APIResponseErr = "invalid id"
	ErrInvalidJSON   APIResponseErr = "invalid json"
	ErrInvalidToken  APIResponseErr = "invalid token"
	ErrUnauthorized  APIResponseErr = "unauthorized"
	ErrMissingToken  APIResponseErr = "missing or malformed token"
	ErrAlreadyExists APIResponseErr = "already exists"
)

func UUIDsFromStrings(ids []string) ([]uuid.UUID, error) {
	var uuids []uuid.UUID
	for _, id := range ids {
		uuid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, uuid)
	}
	return uuids, nil
}
