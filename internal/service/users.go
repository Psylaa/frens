package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepo struct{}

func (ur *UserRepo) Get(c *fiber.Ctx, userID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "user", "GetByID")

	// Get user from database
	user, err := db.Users.GetByID(userID)
	if err != nil {
		logger.ErrorLogRequestError("service", "user", "GetByID", "error getting user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "user", "GetByID", "user retrieved from database")

	// Return the user
	return c.Status(fiber.StatusOK).JSON(response.CreateUsersResponse([]*database.User{user}))
}

func (ur *UserRepo) Create(c *fiber.Ctx, username string, email string, phoneNumber string, password string) error {
	logger.DebugLogRequestReceived("service", "user", "Create")

	// Check if username is taken
	if db.Users.UsernameExists(username) {
		logger.ErrorLogRequestError("service", "user", "Create", "username already taken", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrTakenUsername))
	}
	if db.Users.EmailExists(email) {
		logger.ErrorLogRequestError("service", "user", "Create", "email already taken", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrTakenEmail))
	}
	if db.Users.PhoneNumberExists(phoneNumber) {
		logger.ErrorLogRequestError("service", "user", "Create", "phone already taken", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrTakenPhone))
	}
	logger.DebugLogRequestUpdate("service", "user", "Create", "username and email are available")

	// Create user object
	newUser := database.NewUser(username, email, phoneNumber, password)

	// Create user in database
	err := db.Users.Create(&newUser)
	if err != nil {
		logger.ErrorLogRequestError("service", "user", "Create", "error creating user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "user", "Create", "user created in database")

	// Return the user
	return c.Status(fiber.StatusOK).JSON(response.CreateUsersResponse([]*database.User{&newUser}))
}

func (ur *UserRepo) Update(c *fiber.Ctx, bio *string, avatarID *uuid.UUID, coverID *uuid.UUID) error {
	logger.DebugLogRequestReceived("service", "user", "Update")
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Get the existing user from the database
	user, err := db.Users.GetByID(requestorID)
	if err != nil {
		logger.ErrorLogRequestError("service", "user", "Update", "error getting user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Update the user
	if bio != nil {
		user.Bio = bio
	}

	if avatarID != nil {
		user.AvatarID = avatarID
	}

	if coverID != nil {
		user.CoverID = coverID
	}

	// Update the user in the database
	err = db.Users.Update(user)
	if err != nil {
		logger.ErrorLogRequestError("service", "user", "Update", "error updating user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Return the user
	return c.Status(fiber.StatusOK).JSON(response.CreateUsersResponse([]*database.User{user}))
}

func (ur *UserRepo) Delete(c *fiber.Ctx) error {
	/*
		logger.DebugLogRequestReceived("service", "user", "Delete")

		// Get the userID from context
		userID := c.Locals("requestorID").(*uuid.UUID)
		if userID == nil {
			logger.ErrorLogRequestError("service", "user", "Delete", "userID not found in context", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		logger.DebugLogRequestUpdate("service", "user", "Delete", "userID found in context")

		// Delete user from database
		user, err := db.Users.Delete(userID)
		if err != nil {
			logger.ErrorLogRequestError("service", "user", "Delete", "error deleting user", err)
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		logger.DebugLogRequestUpdate("service", "user", "Delete", "user deleted from database")

		// Return the user
		return c.Status(fiber.StatusOK).JSON(response.CreateUsersResponse([]*database.User{user}))
	*/
	return nil
}
