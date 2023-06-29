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
		Bio            *string `json:"bio"`
		ProfilePicture *string `json:"profilePicture"`
		BannerImage    *string `json:"bannerImage"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidJSON,
		})
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	updatedUser, err := db.Users.UpdateUser(id, *body.Bio, *body.ProfilePicture, *body.BannerImage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Data: []APIResponseData{createAPIResponseData(updatedUser)},
	})
}

// createAPIResponseData converts user to APIResponseData.
func createAPIResponseData(user *database.User) APIResponseData {
	selfLink := fmt.Sprintf("/users/%s", user.ID)
	postsLink := fmt.Sprintf("/users/%s/posts", user.ID)

	return APIResponseData{
		Type: shared.DataTypeUser,
		ID:   &user.ID,
		Attributes: APIResponseDataAttributes{
			CreatedAt:         &user.CreatedAt,
			UpdatedAt:         &user.UpdatedAt,
			Username:          user.Username,
			Bio:               user.Bio,
			Privacy:           user.Privacy,
			ProfilePictureURL: user.ProfilePictureURL,
			CoverImageURL:     user.CoverImageURL,
			//FollowerCount:     followerCount,
			//FollowingCount:    followingCount,
		},
		Links: APIResponseDataLinks{
			Self:      selfLink,
			Posts:     postsLink,
			Following: fmt.Sprintf("%s/following", selfLink),
			Followers: fmt.Sprintf("%s/followers", selfLink),
		},
		Meta: APIResponseDataMeta{
			Version: "1.0",
		},
	}
}
