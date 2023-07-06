package response

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type LoginResponse struct {
	Data []*LoginData `json:"data"`
}

type LoginData struct {
	Type       shared.DataType `json:"type"`
	ID         uuid.UUID       `json:"id"`
	Attributes *LoginAttr      `json:"attributes"`
}

type LoginAttr struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func CreateLoginResponse(token string, user *database.User, expiryDate time.Time) *LoginResponse {
	return &LoginResponse{
		Data: []*LoginData{
			{
				Type: shared.DataTypeToken,
				Attributes: &LoginAttr{
					Token:     token,
					CreatedAt: time.Now(),
					ExpiresAt: expiryDate,
				},
			},
		},
	}
}
