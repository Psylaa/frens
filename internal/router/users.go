package router

import (
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
			Success: false,
			Error:   ErrInternal,
		})
	}

	var data []APIResponseData
	for _, user := range users {
		data = append(data, createAPIResponseData(&user))
	}

	return c.JSON(APIResponse{
		Success: true,
		Data:    data,
	})
}

// getUser handles the HTTP request to fetch a specific user.
func getUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	user, err := db.Users.GetUser(id)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error getting user")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data:    []APIResponseData{createAPIResponseData(user)},
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
			Success: false,
			Error:   ErrInvalidJSON,
		})
	}

	user, err := db.Users.CreateUser(body.Username, body.Email, body.Password)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error creating user: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(APIResponse{
		Success: true,
		Data:    []APIResponseData{createAPIResponseData(user)},
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
			Success: false,
			Error:   ErrInvalidJSON,
		})
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	updatedUser, err := db.Users.UpdateUser(id, *body.Bio, *body.ProfilePicture, *body.BannerImage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data:    []APIResponseData{createAPIResponseData(updatedUser)},
	})
}

// createAPIResponseData converts user to APIResponseData.
func createAPIResponseData(user *database.User) APIResponseData {
	return APIResponseData{
		Type: shared.DataTypeUser,
		ID:   user.ID,
		Attributes: APIResponseDataAttributes{
			Privacy: user.Privacy,
		},
		Relationships: APIResponseDataRelationships{
			OwnerID: user.ID,
		},
	}
}
