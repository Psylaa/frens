package activitypub

import (
	"encoding/json"
	"strings"

	db "github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

type Activity struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	ID      string `json:"id"`
	To      []string
	Actor   string
	Object  Object
}

type Object struct {
	Type         string `json:"type"`
	ID           string `json:"id"`
	Content      string `json:"content"`
	AttributedTo string `json:"attributedTo"`
	To           []string
}

func GetUserProfile(c *fiber.Ctx) error {
	requestedUsername := strings.ToLower(c.Params("username"))
	if requestedUsername == "" {
		return c.Status(fiber.StatusBadRequest).SendString("No username specified in request")
	}

	format := c.Accepts("html", "application/activity+json")
	if format == "" {
		return c.Status(fiber.StatusNotAcceptable).SendString("Invalid Accept header")
	}

	if format == "html" {
		return c.Redirect("/@"+requestedUsername, fiber.StatusSeeOther)
	}

	// Get the user data
	user, err := db.GetUserByUsername(requestedUsername)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user")
	}

	// Marshal the user data into JSON
	userData, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to convert user data to JSON")
	}

	c.Set(fiber.HeaderContentType, format)
	return c.Send(userData)
}
