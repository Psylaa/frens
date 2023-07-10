package response

import (
	"fmt"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type UserResponse struct {
	Data []*UserData `json:"data"`
}

type UserData struct {
	Type       shared.DataType `json:"type"`
	ID         uuid.UUID       `json:"id"`
	Attributes UserAttr        `json:"attributes"`
	Links      UserLinks       `json:"links"`
}

type UserAttr struct {
	Username string `json:"username"`
}

type UserLinks struct {
	Self string `json:"self"`
}

func CreateUsersResponse(users []*database.User) *UserResponse {
	var userData []*UserData

	for _, user := range users {
		selfLink := fmt.Sprintf("%s/v1/users/%s", baseURL, user.ID.String())

		userData = append(userData, &UserData{
			Type: shared.DataTypeUser,
			ID:   user.ID,
			Attributes: UserAttr{
				Username: user.Username,
			},
			Links: UserLinks{
				Self: selfLink,
			},
		})
	}

	return &UserResponse{
		Data: userData,
	}
}
