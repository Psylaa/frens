package activitypub

import (
	db "github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

type Collection struct {
	Context    string `json:"@context"`
	Type       string `json:"type"`
	TotalItems int    `json:"totalItems"`
	Items      []Note `json:"items"`
}

type Note struct {
	Type         string `json:"type"`
	ID           string `json:"id"`
	Content      string `json:"content"`
	AttributedTo string `json:"attributedTo"`
	To           []string
}

func HandleOutbox(c *fiber.Ctx) error {
	// Get the user's username from the URL
	username := c.Params("username")

	// Get the user's data
	user, err := db.GetUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user")
	}

	// Get the user's statuses
	statuses, err := db.GetPostsByUserID(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get statuses")
	}

	// Convert the statuses to notes
	notes := make([]Note, len(statuses))
	for i, status := range statuses {
		notes[i] = Note{
			Type:         "Note",
			ID:           "/statuses/" + status.ID.String(),
			Content:      status.Text,
			AttributedTo: "/users/" + username,
			To:           []string{"/users/" + username},
		}
	}

	// Create the collection
	collection := Collection{
		Context:    "https://www.w3.org/ns/activitystreams",
		Type:       "OrderedCollection",
		TotalItems: len(notes),
		Items:      notes,
	}

	// Return the collection
	return c.JSON(collection)
}
