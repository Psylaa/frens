package response

type RespErr string

const (
	ErrForbidden        RespErr = "forbidden"
	ErrInternal         RespErr = "internal server error"
	ErrInvalidCursor    RespErr = "invalid cursor"
	ErrInvalidID        RespErr = "invalid id"
	ErrInvalidToken     RespErr = "invalid token"
	ErrInvalidUUID      RespErr = "invalid uuid"
	ErrMissingToken     RespErr = "missing or malformed token"
	ErrNotFound         RespErr = "not found"
	ErrUnauthorized     RespErr = "unauthorized"
	ErrExists           RespErr = "already exists"
	ErrInvalidBody      RespErr = "invalid body"
	ErrInvalidMediaUUID RespErr = "unable to parse provided media ids into uuids"
	ErrMediaIDsNotFound RespErr = "unable to find media with provided ids"
)

type ErrResp struct {
	Error RespErr `json:"error"`
}
