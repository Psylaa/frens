package shared

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
