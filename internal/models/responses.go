package models

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
