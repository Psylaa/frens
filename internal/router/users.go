package router

import (
	"fmt"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// getUsers handles the HTTP request to fetch all users.
func getUsers(c *fiber.Ctx) error {
	users, err := db.Users.GetUsers()
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting all users")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	var data []APIResponseData
	for _, user := range users {
		data = append(data, createAPIResponseData(&user))
	}

	return c.JSON(APIResponse{
		Data: data,
	})
}

// getUser handles the HTTP request to fetch a specific user.
func getUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}
	logger.Log.Debug().Msgf("Successfully parsed user ID: %v", id)

	user, err := db.Users.GetUser(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting user")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().Msgf("Successfully retrieved user: %v", user)

	return c.JSON(APIResponse{
		Data: []APIResponseData{createAPIResponseData(user)},
	})
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
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidJSON,
		})
	}

	user, err := db.Users.CreateUser(body.Username, body.Email, body.Password)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error creating user: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Data: []APIResponseData{createAPIResponseData(user)},
	})
}

// updateUser handles the HTTP request to update a user's details.
func updateUser(c *fiber.Ctx) error {
	var body struct {
		Bio              *string `json:"bio"`
		ProfilePictureID *string `json:"profilePictureId"`
		CoverImageID     *string `json:"coverImageId"`
	}

	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidJSON,
		})
	}

	userId, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	if body.Bio != nil {
		if err := db.Users.UpdateBio(userId, body.Bio); err != nil {
			logger.Log.Error().Err(err).Msg("Error updating bio")
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
				Error: ErrInternal,
			})
		}
	}

	if body.ProfilePictureID != nil {
		ppUUID, err := uuid.Parse(*body.ProfilePictureID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing ProfilePictureID")
			return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
				Error: ErrInvalidUUID,
			})
		}
		if err := db.Users.UpdateProfilePicture(userId, &ppUUID); err != nil {
			logger.Log.Error().Err(err).Msg("Error updating profile picture")
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
				Error: ErrInternal,
			})
		}
	}

	if body.CoverImageID != nil {
		ciUUID, err := uuid.Parse(*body.CoverImageID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing CoverImageID")
			return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
				Error: ErrInvalidUUID,
			})
		}
		if err := db.Users.UpdateCoverImage(userId, &ciUUID); err != nil {
			logger.Log.Error().Err(err).Msg("Error updating cover image")
			return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
				Error: ErrInternal,
			})
		}
	}

	// Retrieve updated user
	user, err := db.Users.GetUser(userId)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting user")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.Status(fiber.StatusOK).JSON(createAPIResponseData(user))
}

// createAPIResponseData converts user to APIResponseData.
func createAPIResponseData(user *database.User) APIResponseData {
	selfLink := fmt.Sprintf("/users/%s", user.ID)
	postsLink := fmt.Sprintf("/users/%s/posts", user.ID)

	return APIResponseData{
		Type: shared.DataTypeUser,
		ID:   &user.ID,
		Attributes: APIResponseDataAttributes{
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
			Username:  user.Username,
			Bio:       user.Bio,
			Privacy:   user.Privacy,
			//FollowerCount:     followerCount,
			//FollowingCount:    followingCount,
		},
		Links: APIResponseDataLinks{
			Self:           selfLink,
			Posts:          postsLink,
			Following:      fmt.Sprintf("%s/following", selfLink),
			Followers:      fmt.Sprintf("%s/followers", selfLink),
			ProfilePicture: "/files/" + user.ProfilePicture.ID.String() + user.ProfilePicture.Extension,
			CoverImage:     "/files/" + user.CoverImage.ID.String() + user.CoverImage.Extension,
		},
	}
}
