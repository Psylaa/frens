package router

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/gofiber/fiber/v2"
)

// getUsers handles the HTTP request to fetch all users.
func getUsers(c *fiber.Ctx) error {
	/*
		users, err := db.Users.GetUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusOK).JSON(response.GenerateUsersResponse(users))
	*/
	return nil
}

func getUser(c *fiber.Ctx) error {
	/*
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
		}
		logger.Log.Debug().Msgf("Successfully parsed user ID: %v", id)

		user, err := db.Users.GetUser(id)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error getting user")
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		logger.Log.Debug().Msgf("Successfully retrieved user: %v", user)

		return c.Status(fiber.StatusOK).JSON(response.GenerateUserResponse(user))
	*/
	return nil
}

// createUser handles the HTTP request to create a new user.
func createUser(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	return srv.Users.Create(c, body.Username, body.Email, body.Password)
}

// updateUser handles the HTTP request to update a user's details.
func updateUser(c *fiber.Ctx) error {
	/*
		var body struct {
			Bio              *string `json:"bio"`
			AvatarID *string `json:"avatarId"`
			CoverID     *string `json:"coverId"`
		}

		if err := c.BodyParser(&body); err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing request body")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
		}

		userId, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrUnauthorized))
		}

		if body.Bio != nil {
			if err := db.Users.UpdateBio(userId, body.Bio); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating bio")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		if body.AvatarID != nil {
			ppUUID, err := uuid.Parse(*body.AvatarID)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Error parsing AvatarID")
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
			}
			if err := db.Users.UpdateAvatar(userId, &ppUUID); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating profile picture")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		if body.CoverID != nil {
			ciUUID, err := uuid.Parse(*body.CoverID)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Error parsing CoverID")
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
			}
			if err := db.Users.UpdateCover(userId, &ciUUID); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating cover image")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		// Retrieve updated user
		user, err := db.Users.GetUser(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusOK).JSON(response.GenerateUserResponse(user))
	*/
	return nil
}
