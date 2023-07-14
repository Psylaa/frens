package models

type UserResponse struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}
