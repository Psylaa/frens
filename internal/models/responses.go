package models

import "github.com/google/uuid"

type Relationship struct {
	User  *RelationshipData `json:"user,omitempty"`
	Post  *RelationshipData `json:"post,omitempty"`
	Media *RelationshipData `json:"media,omitempty"`
}

type RelationshipData struct {
	Links RelationshipLinks   `json:"links"`
	Data  RelationshipDetails `json:"data,omitempty"`
}

type RelationshipLinks struct {
	Self string `json:"self"`
}

type RelationshipDetails struct {
	Type DataType  `json:"type"`
	ID   uuid.UUID `json:"id"`
}

// ToResponse takes UserData to create a UserResponse
func CreateUserResponse(userData UserData) *UserResponse {
	return &UserResponse{
		Links: UserLinks{
			Self: "todo",
		},
		Data: []UserData{userData},
	}
}

func CreatePostResponse(postData PostData, userData UserData) *PostResponse {
	return &PostResponse{
		Links: PostLinks{
			Self: "todo",
		},
		Data:     []PostData{postData},
		Included: []UserData{userData},
	}
}

func CreateLikeResponse(likeData LikeData) *LikeResponse {
	return &LikeResponse{
		Links: LikeLinks{
			Self: "todo",
		},
		Data: []LikeData{likeData},
	}
}
