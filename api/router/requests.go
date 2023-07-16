package router

type RegisterRequest struct {
	Username string `validate:"required,min=2,max=100"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
