package activitypub

import (
	"strings"

	db "github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

func HandleInbox(c *fiber.Ctx) error {
	// Parse the incoming activity
	var activity Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid activity")
	}

	// Verify the activity
	if activity.Type != "Create" || activity.Object.Type != "Note" {
		return c.Status(fiber.StatusBadRequest).SendString("Unsupported activity type")
	}

	// We'll also need to verify that the `actor` is who they say they are,
	// and that the `to` field includes the user whose inbox we're posting to.
	// For simplicity, we'll skip this step for now.

	// Handle the activity
	if err := createStatusFromActivity(activity); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to handle activity")
	}

	return c.SendStatus(fiber.StatusOK)
}

func createStatusFromActivity(activity Activity) error {
	// Parse the actor's username from the activity
	// This assumes that the actor's URL is of the form "/users/:username"
	actorUsername := strings.TrimPrefix(activity.Actor, "/users/")

	// Get the actor's user data
	actor, err := db.GetUserByUsername(actorUsername)
	if err != nil {
		return err
	}

	// Create a new status from the activity
	_, err = db.CreateStatus(actor.ID, activity.Object.Content, nil)
	if err != nil {
		return err
	}

	return nil
}
