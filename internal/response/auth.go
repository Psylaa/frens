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
	Type       shared.DataType `json:"type"`
	Attributes *AuthAttr       `json:"attributes"`
}

type AuthAttr struct {
	Token     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func CreateAuthResponse(token string, user *database.User, expiryDate time.Time) *AuthResponse {
	return &AuthResponse{
		Data: []*AuthData{
			{
				Type: shared.DataTypeToken,
				Attributes: &AuthAttr{
					Token:     token,
					CreatedAt: time.Now(),
					ExpiresAt: expiryDate,
				},
			},
		},
	}
}
