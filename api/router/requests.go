package router

type RegisterRequest struct {
	Username string `validate:"required,min=2,max=100"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type CreatePostRequest struct {
	Text    string `validate:"required,min=1,max=1000"`
	Privacy string `validate:"omitempty,oneof=public protected private"`
}
