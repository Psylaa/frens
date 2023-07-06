package router_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TestRouter_extractRequestorID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}

	tests := []struct {
		name        string
		requestorID string
		wantErr     bool
	}{
		{
			name:        "valid requestor id",
			requestorID: "123e4567-e89b-12d3-a456-426614174000",
		},
		{
			name:        "invalid requestor id",
			requestorID: "123e4567-e89b-12d3-a456-42661417400",
		}
	}
	for _, tt := range tests {

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["requestor_id"] = tt.requestorID

		// Create fiber context
		c := &fiber.Ctx{}
		c.Locals("user", token)

		// Create router

	}
}
