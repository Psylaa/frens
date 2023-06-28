package router

import (
	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type APIResponseErr string

const (
	ErrInternal     APIResponseErr = "internal server error"
	ErrNotFound     APIResponseErr = "not found"
	ErrInvalidID    APIResponseErr = "invalid id"
	ErrInvalidJSON  APIResponseErr = "invalid json"
	ErrInvalidToken APIResponseErr = "invalid token"
	ErrUnauthorized APIResponseErr = "unauthorized"
	ErrMissingToken APIResponseErr = "missing or malformed token"
)

type APIResponse struct {
	Success bool              `json:"success"`
	Data    []APIResponseData `json:"data,omitempty"`
	Error   APIResponseErr    `json:"errors,omitempty"`
}

type APIResponseData struct {
	ID            uuid.UUID                    `json:"id"`
	Type          shared.DataType              `json:"type"`
	Attributes    APIResponseDataAttributes    `json:"attributes,omitempty"`
	Relationships APIResponseDataRelationships `json:"relationships,omitempty"`
}

type APIResponseDataAttributes struct {
	Privacy           shared.Privacy `json:"privacy,omitempty"`
	Text              string         `json:"text,omitempty"`
	Token             string         `json:"token,omitempty"`
	Filename          string         `json:"filename,omitempty"`
	Extension         string         `json:"extension,omitempty"`
	ProfilePictureURL string         `json:"profilePictureUrl,omitempty"`
	CoverImageURL     string         `json:"coverImageUrl,omitempty"`
	ExpiresAt         string         `json:"expiresAt,omitempty"`
}

type APIResponseDataRelationships struct {
	OwnerID uuid.UUID `json:"ownerId,omitempty"`
}

type Router struct {
	Config *config.Config
	App    *fiber.App
}
