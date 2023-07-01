package router

import (
	"github.com/gofiber/fiber/v2"
)

// getUsers handles the HTTP request to fetch all users.
func retrieveAllUsers(c *fiber.Ctx) error {
	/*
		users, err := db.Users.GetUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusOK).JSON(response.GenerateUsersResponse(users))
	*/
	return nil
}

func retrieveUserDetails(c *fiber.Ctx) error {
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
func registerUser(c *fiber.Ctx) error {
	/*
		var body struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&body); err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing request body")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
		}

		user, err := db.Users.CreateUser(body.Username, body.Email, body.Password)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error creating user: " + err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusOK).JSON(response.GenerateUserResponse(user))
	*/
	return nil
}

// updateUser handles the HTTP request to update a user's details.
func updateUserDetails(c *fiber.Ctx) error {
	/*
		var body struct {
			Bio              *string `json:"bio"`
			ProfilePictureID *string `json:"profilePictureId"`
			CoverImageID     *string `json:"coverImageId"`
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

		if body.ProfilePictureID != nil {
			ppUUID, err := uuid.Parse(*body.ProfilePictureID)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Error parsing ProfilePictureID")
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
			}
			if err := db.Users.UpdateProfilePicture(userId, &ppUUID); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating profile picture")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		if body.CoverImageID != nil {
			ciUUID, err := uuid.Parse(*body.CoverImageID)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Error parsing CoverImageID")
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
			}
			if err := db.Users.UpdateCoverImage(userId, &ciUUID); err != nil {
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
