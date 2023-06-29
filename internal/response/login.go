package response

import (
	"time"

	"github.com/google/uuid"
)

type LoginResp struct {
	Links    LoginResp_Links      `json:"links,omitempty"`
	Data     []LoginResp_Data     `json:"data,omitempty"`
	Included []LoginResp_Included `json:"included,omitempty"`
}

type LoginResp_Links struct {
	Self string `json:"self"`
}

type LoginResp_Data struct {
	Type       string                   `json:"type"`
	Attributes LoginResp_DataAttributes `json:"attributes"`
}

type LoginResp_DataAttributes struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
	UserID     uuid.UUID `json:"userId,omitempty"`
}

type LoginResp_Included struct {
	// For future use
}

func GenerateLoginResponse(token string, exp time.Time, userId uuid.UUID) *LoginResp {
	return &LoginResp{
		Links: LoginResp_Links{
			Self: "/login",
		},
		Data: []LoginResp_Data{
			{
				Type: "login",
				Attributes: LoginResp_DataAttributes{
					Token:      token,
					Expiration: exp,
					UserID:     userId,
				},
			},
		},
	}
}
