package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserRespone struct {
	Links UserLinks  `json:"links"`
	Data  []UserData `json:"data"`
}

type UserLinks struct {
	Self string `json:"self"`
}

type UserData struct {
	Type       DataTypes      `json:"type"`
	ID         uuid.UUID      `json:"id"`
	Attributes UserAttributes `json:"attributes"`
}

type UserAttributes struct {
	Role      Role    `json:"role"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Verrified bool    `json:"verified"`
	Token     *string `json:"token,omitempty"`
}

func (ur *UserRespone) AddToken(signingKey []byte, duration time.Duration) error {
	claims := jwt.MapClaims{
		"name": ur.Data[0].Attributes.Username,
		"sub":  ur.Data[0].ID.String(),
		"role": ur.Data[0].Attributes.Role,
		"exp":  time.Now().Add(time.Hour * duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	ur.Data[0].Attributes.Token = &t
	return nil
}

func (ur *UserRespone) Send(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ur)
}
