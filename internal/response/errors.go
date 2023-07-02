package response

type RespErr string

const (
	ErrExists       RespErr = "already exists"
	ErrForbidden    RespErr = "forbidden"
	ErrInternal     RespErr = "internal server error"
	ErrUnauthorized RespErr = "unauthorized"

	ErrInvalidBody      RespErr = "invalid body"
	ErrInvalidCursor    RespErr = "invalid cursor"
	ErrInvalidEmail     RespErr = "email is not valid"
	ErrInvalidID        RespErr = "invalid id"
	ErrInvalidMediaUUID RespErr = "unable to parse provided media ids into uuids"
	ErrInvalidPassword  RespErr = "password is not valid"
	ErrInvalidToken     RespErr = "invalid token"
	ErrInvalidUsername  RespErr = "username is not valid"
	ErrInvalidUUID      RespErr = "invalid uuid"
	ErrInvalidFileID    RespErr = "invalid file id"

	ErrMissingMedia     RespErr = "unable to find media with provided ids"
	ErrMissingToken     RespErr = "missing or malformed token"
	ErrMissingExtension RespErr = "missing file extension"

	ErrNotFound RespErr = "not found"

	ErrTakenEmail    RespErr = "email is already in use"
	ErrTakenUsername RespErr = "username is already in use"

	ErrFileIDNotUUID     RespErr = "file id is not a valid uuid"
	ErrFileIDNotProvided RespErr = "file id not provided"
)

type ErrResp struct {
	Error RespErr `json:"error"`
}
