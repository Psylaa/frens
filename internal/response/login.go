package response

import "time"

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
	ID         string                   `json:"id"`
	Attributes LoginResp_DataAttributes `json:"attributes"`
}

type LoginResp_DataAttributes struct {
	// For future use
}

type LoginResp_Included struct {
	// For future use
}

func GenerateLoginResponse(token string, exp time.Time) *LoginResp {
	return &LoginResp{
		Links: LoginResp_Links{
			Self: "/login",
		},
		Data: []LoginResp_Data{
			{
				Type:       "login",
				ID:         "1",
				Attributes: LoginResp_DataAttributes{
					// For future use
				},
			},
		},
	}
}
