package response

type RespErr string

const (
	ErrExists       RespErr = "already exists"
	ErrForbidden    RespErr = "forbidden"
	ErrInternal     RespErr = "internal server error"
	ErrUnauthorized RespErr = "unauthorized"
	ErrNotFound     RespErr = "not found"

	ErrInvalidBody        RespErr = "invalid body"
	ErrInvalidCount       RespErr = "invalid count"
	ErrInvalidCredentials RespErr = "invalid credentials"
	ErrInvalidCursor      RespErr = "invalid cursor"
	ErrInvalidEmail       RespErr = "invalid email"
	ErrInvalidFileID      RespErr = "invalid file id"
	ErrInvalidID          RespErr = "invalid id"
	ErrInvalidMediaUUID   RespErr = "unable to parse provided media ids into uuids"
	ErrInvalidOffset      RespErr = "invalid offset"
	ErrInvalidPassword    RespErr = "invalid password"
	ErrInvalidToken       RespErr = "invalid token"
	ErrInvalidUsername    RespErr = "invalid username"
	ErrInvalidUUID        RespErr = "invalid uuid"

	ErrMissingMedia     RespErr = "unable to find media with provided ids"
	ErrMissingToken     RespErr = "missing or malformed token"
	ErrMissingExtension RespErr = "missing file extension"

	ErrTakenEmail    RespErr = "email is already in use"
	ErrTakenUsername RespErr = "username is already in use"

	ErrFileIDNotUUID     RespErr = "file id is not a valid uuid"
	ErrFileIDNotProvided RespErr = "file id not provided"
)

type ErrResp struct {
	Error RespErr `json:"error"`
}
