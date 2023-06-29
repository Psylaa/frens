package response

type APIResponseErr string

const (
	ErrForbidden     APIResponseErr = "forbidden"
	ErrInternal      APIResponseErr = "internal server error"
	ErrInvalidCursor APIResponseErr = "invalid cursor"
	ErrInvalidID     APIResponseErr = "invalid id"
	ErrInvalidToken  APIResponseErr = "invalid token"
	ErrInvalidUUID   APIResponseErr = "invalid uuid"
	ErrMissingToken  APIResponseErr = "missing or malformed token"
	ErrNotFound      APIResponseErr = "not found"
	ErrUnauthorized  APIResponseErr = "unauthorized"
	ErrExists        APIResponseErr = "already exists"
	ErrInvalidBody   APIResponseErr = "invalid body"
)

type ErrResp struct {
	Error APIResponseErr `json:"error"`
}

var baseURL string

func Init(baseUrl string) {
	baseURL = baseUrl
}

func GenerateErrorResponse(err APIResponseErr) *ErrResp {
	return &ErrResp{
		Error: err,
	}
}
