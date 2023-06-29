package response

import (
	"fmt"
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/google/uuid"
)

type UserResp struct {
	Links    UserResp_Links      `json:"links,omitempty"`
	Data     []UserResp_Data     `json:"data,omitempty"`
	Included []UserResp_Included `json:"included,omitempty"`
}

type UserResp_Data struct {
	Type       string                  `json:"type"`
	ID         uuid.UUID               `json:"id,omitempty"`
	Attributes UserResp_DataAttributes `json:"attributes"`
}

type UserResp_DataAttributes struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Username  string         `json:"username"`
	Bio       string         `json:"bio"`
	Privacy   shared.Privacy `json:"privacy"`
}

type UserResp_Links struct {
	Self           string `json:"self"`
	Posts          string `json:"posts"`
	Following      string `json:"following"`
	Followers      string `json:"followers"`
	ProfilePicture string `json:"profilePicture"`
	CoverImage     string `json:"coverImage"`
}

type UserResp_Included struct {
}

func GenerateUserResponse(user *database.User) *UserResp {
	selfLink := fmt.Sprintf("/users/%s", user.ID)
	postsLink := fmt.Sprintf("/users/%s/posts", user.ID)

	return &UserResp{
		Links: UserResp_Links{
			Self:           selfLink,
			Posts:          postsLink,
			Following:      fmt.Sprintf("%s/following", selfLink),
			Followers:      fmt.Sprintf("%s/followers", selfLink),
			ProfilePicture: "/files/" + user.ProfilePicture.ID.String() + user.ProfilePicture.Extension,
			CoverImage:     "/files/" + user.CoverImage.ID.String() + user.CoverImage.Extension,
		},
		Data: []UserResp_Data{
			{
				Type: "user",
				ID:   user.ID,
				Attributes: UserResp_DataAttributes{
					CreatedAt: user.CreatedAt,
					UpdatedAt: user.UpdatedAt,
					Username:  user.Username,
					Bio:       user.Bio,
					Privacy:   user.Privacy,
					//FollowerCount:     followerCount,  // to be implemented
					//FollowingCount:    followingCount, // to be implemented
				},
			},
		},
	}
}

func GenerateUsersResponse(users []*database.User) *UserResp {
	usersResp := UserResp{
		Data: make([]UserResp_Data, 0, len(users)),
	}

	for _, user := range users {
		userResp := GenerateUserResponse(user)
		usersResp.Data = append(usersResp.Data, userResp.Data[0])
	}

	return &usersResp
}
