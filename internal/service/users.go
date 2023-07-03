package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepo struct{}

func (ur *UserRepo) GetByID(c *fiber.Ctx, userID *uuid.UUID) error {
	logger.DebugLogRequestRecieved("service", "user", "GetByID")

	// Get user from database
	user, err := db.Users.GetByID(c.Locals("requestorId").(*uuid.UUID), userID)
	if err != nil {
		logger.ErrorLogRequestError("service", "user", "GetByID", "user not found", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "user", "GetByID", "user found")

	// Return the user
	return c.Status(fiber.StatusOK).JSON(response.CreateUsersResponse([]*database.User{user}))
}

func (ur *UserRepo) GetByUsername(username string) error {
	return nil
}

func (ur *UserRepo) GetFollowers(userID string) error {
	return nil
}

func (ur *UserRepo) GetFollowersCount(userID string) error {
	return nil
}

func (ur *UserRepo) GetFollowing(userID string) error {
	return nil
}

func (ur *UserRepo) GetFollowingCount(userID string) error {
	return nil
}

func (ur *UserRepo) GetPosts(userID string) error {
	return nil
}

func (ur *UserRepo) GetBookmarks(userID string) error {
	return nil
}

func (ur *UserRepo) GetBookmarksCount(userID string) error {
	return nil
}

func (ur *UserRepo) GetMedia(userID string) error {
	return nil
}

func (ur *UserRepo) GetNotifications(userID string) error {
	return nil
}

func (ur *UserRepo) GetSettings(userID string) error {
	return nil
}

func (ur *UserRepo) Create(c *fiber.Ctx, username string, email string, password string) error {
	logger.DebugLogRequestRecieved("service", "user", "Create")

	// Check if username is taken
	if db.Users.UsernameExists(&username) {
		logger.ErrorLogRequestError("service", "user", "Create", "username taken", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrTakenUsername))
	}
	logger.DebugLogRequestUpdate("service", "user", "Create", "username available")

	// Check if email is taken
	if db.Users.EmailExists(&email) {
		logger.ErrorLogRequestError("service", "user", "Create", "email taken", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrTakenEmail))
	}
	logger.DebugLogRequestUpdate("service", "user", "Create", "email available")

	// Create user in database
	user, err := db.Users.CreateUser(username, email, password)
	if err != nil {
		logger.ErrorLogRequestError("service", "user", "Create", "error creating user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}
	logger.DebugLogRequestUpdate("service", "user", "Create", "user created")

	// Return the user
	return c.Status(fiber.StatusOK).JSON(response.CreateUsersResponse([]*database.User{user}))
}

func (ur *UserRepo) Update() error {
	return nil
}

func (ur *UserRepo) Delete() error {
	return nil
}
