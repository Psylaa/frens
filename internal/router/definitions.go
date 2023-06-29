package router

import (
	"time"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type APIResponseErr string

// Errors constants
const (
	ErrInternal     APIResponseErr = "internal server error"
	ErrNotFound     APIResponseErr = "not found"
	ErrInvalidID    APIResponseErr = "invalid id"
	ErrInvalidJSON  APIResponseErr = "invalid json"
	ErrInvalidToken APIResponseErr = "invalid token"
	ErrUnauthorized APIResponseErr = "unauthorized"
	ErrMissingToken APIResponseErr = "missing or malformed token"
	ErrForbidden    APIResponseErr = "forbidden"
	ErrInvalidUUID  APIResponseErr = "invalid uuid"
)

// Main API response structure
type APIResponse struct {
	Data     []APIResponseData         `json:"data,omitempty"`
	Included []APIResponseDataIncluded `json:"included,omitempty"`
	Error    APIResponseErr            `json:"errors,omitempty"`
}

// Structure representing individual resource object
type APIResponseData struct {
	Type          shared.DataType               `json:"type"`
	ID            *uuid.UUID                    `json:"id,omitempty"` // Only empty for token requests
	Attributes    APIResponseDataAttributes     `json:"attributes"`
	Relationships *APIResponseDataRelationships `json:"relationships,omitempty"`
	Included      []APIResponseDataIncluded     `json:"included,omitempty"`
	Links         APIResponseDataLinks          `json:"links,omitempty"`
}

type APIResponseDataIncluded struct {
	Type shared.DataType `json:"type"`
}

// Structure for individual resource object attributes
type APIResponseDataAttributes struct {
	// Common attributes
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	ExpiresAt string     `json:"expiresAt,omitempty"` // Must be string for RFC 7519 on JWT aka "NumericDate"

	// User specific attributes
	Username       string         `json:"username,omitempty"`
	Bio            string         `json:"bio,omitempty"`
	Privacy        shared.Privacy `json:"privacy,omitempty"`
	FollowerCount  int            `json:"followerCount,omitempty"`
	FollowingCount int            `json:"followingCount,omitempty"`

	// Follow specific attributes
	SourceID *uuid.UUID `json:"sourceId,omitempty"`
	TargetID *uuid.UUID `json:"targetId,omitempty"`

	// Post specific attributes
	Text          string `json:"text,omitempty"`
	LikeCount     int    `json:"likeCount,omitempty"`
	BookmarkCount int    `json:"bookmarkCount,omitempty"`
	ViewCount     int    `json:"viewCount,omitempty"`

	// File specific attributes
	Extension string `json:"extension,omitempty"`

	// Token
	Token string `json:"token,omitempty"`
}

// Structure for individual resource object relationships
type APIResponseDataRelationships struct {
	Author *APIResponseDataRelationshipsEntity `json:"author,omitempty"`
	Owner  *APIResponseDataRelationshipsEntity `json:"owner,omitempty"`
}

type APIResponseDataRelationshipsEntity struct {
	Data *APIResponseDataRelationshipsEntityData `json:"data"`
}

type APIResponseDataRelationshipsEntityData struct {
	Type shared.DataType `json:"type"`
	ID   *uuid.UUID      `json:"id,omitempty"`
}

// Structure for individual resource object links
type APIResponseDataLinks struct {
	Self           string `json:"self,omitempty"`
	Author         string `json:"author,omitempty"`
	Posts          string `json:"posts,omitempty"`
	Source         string `json:"source,omitempty"`
	Target         string `json:"target,omitempty"`
	Following      string `json:"following,omitempty"`
	Followers      string `json:"followers,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
	CoverImage     string `json:"coverImage,omitempty"`
}

// Structure representing the router
type Router struct {
	Config *config.Config
	App    *fiber.App
}
