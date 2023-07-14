package models

type UpdateUserRequest struct {
	Bio      *string `json:"bio"`
	AvatarID *string `json:"avatar_id"`
	CoverID  *string `json:"cover_id"`
}
