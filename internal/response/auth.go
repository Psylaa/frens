package response

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
)

type AuthResponse struct {
	Data []*AuthData `json:"data"`
}

type AuthData struct {
	Type          shared.DataType `json:"type"`
	Attributes    *AuthAttr       `json:"attributes"`
	Relationships *AuthRel        `json:"relationships"`
}

type AuthAttr struct {
	Token     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AuthRel struct {
	Owner *UserResponse `json:"owner"`
}

func CreateAuthResponse(token string, user *database.User, expiryDate time.Time) *AuthResponse {
	return nil
}
