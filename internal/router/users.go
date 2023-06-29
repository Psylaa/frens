package router

import (
	"errors"
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
	// parse form data
	// update to body parser at some point
	bio := c.FormValue("bio")
	profilePictureID := c.FormValue("profilePictureId")
	coverImageID := c.FormValue("coverImageId")

	// if form values are empty, set to nil
	var bioPtr *string
	if bio != "" {
		logger.Log.Debug().Msgf("Successfully parsed bio: %v", bio)
		bioPtr = &bio
	} else {
		logger.Log.Debug().Msg("Bio is empty. Not updating.")
	}

	profilePicturePtr, err := getProfilePicturePtr(profilePictureID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting profile picture")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}
	coverImagePtr, err := getCoverImagePtr(coverImageID)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting cover image")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	id, err := getUserID(c)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing user ID")
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}
	logger.Log.Debug().Msgf("Successfully parsed user ID: %v", id)

	updatedUser, err := db.Users.UpdateUser(id, bioPtr, profilePicturePtr, coverImagePtr)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error updating user")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().Msgf("Successfully updated user: %v", updatedUser)

	// Retrieve the user again to get the new file objects
	updatedUser, err = db.Users.GetUser(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting user after update")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Data: []APIResponseData{createAPIResponseData(updatedUser)},
	})
}

func getProfilePicturePtr(profilePictureID string) (*database.File, error) {
	var profilePicturePtr *database.File
	if profilePictureID != "" {
		logger.Log.Debug().Msgf("Successfully parsed profile picture ID: %v", profilePictureID)
		id, err := uuid.Parse(profilePictureID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing profile picture ID")
			return nil, errors.New("invalid profile picture ID format")
		}
		profilePicture, err := db.Files.GetFile(id)
		if err != nil || profilePicture == nil {
			logger.Log.Error().Err(err).Msg("Error getting profile picture")
			return nil, errors.New("profile picture not found")
		}
		logger.Log.Debug().Msgf("Successfully retrieved profile picture: %v", profilePicture)
		profilePicturePtr = profilePicture
	}
	return profilePicturePtr, nil
}

func getCoverImagePtr(coverImageID string) (*database.File, error) {
	var coverImagePtr *database.File
	if coverImageID != "" {
		logger.Log.Debug().Msgf("Successfully parsed cover image ID: %v", coverImageID)
		id, err := uuid.Parse(coverImageID)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing cover image ID")
			return nil, errors.New("invalid cover image ID format")
		}
		coverImage, err := db.Files.GetFile(id)
		if err != nil || coverImage == nil {
			logger.Log.Error().Err(err).Msg("Error getting cover image")
			return nil, errors.New("cover image not found")
		}
		logger.Log.Debug().Msgf("Successfully retrieved cover image: %v", coverImage)
		coverImagePtr = coverImage
	}
	return coverImagePtr, nil
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
			ProfilePicture: "/files/" + user.ProfilePicture.ID.String() + "." + user.ProfilePicture.Extension,
			CoverImage:     "/files/" + user.CoverImage.ID.String() + "." + user.CoverImage.Extension,
		},
		Meta: APIResponseDataMeta{
			Version: "1.0",
		},
	}
}
