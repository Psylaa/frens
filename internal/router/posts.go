package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// getPost handles the HTTP request to fetch a specific post.
func getPost(c *fiber.Ctx) error {
	postID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Log.Debug().
			Str("package", "router").
			Msgf("error parsing provided post id into uuid: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	post, err := db.Posts.GetPost(postID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		logger.Log.Debug().
			Str("package", "router").
			Msgf("database returned record not found error: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Error: ErrNotFound,
		})
	default:
		logger.Log.Debug().
			Str("package", "router").
			Msgf("database returned unknown error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Data: []APIResponseData{createAPIResponseDataPost(post)},
	})
}

// getPosts handles the HTTP request to fetch all posts by a user.
func GetPostsByUserID(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Query("userId"))
	logger.Log.Debug().Msgf("userID: %v", userID)
	if err != nil {
		logger.Log.Debug().
			Str("package", "router").
			Msgf("error parsing provided user id into uuid: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}
	logger.Log.Debug().Msgf("successfully parsed provided user id into uuid: %v", userID)

	posts, err := db.Posts.GetPostsByUserID(userID)
	if err != nil {
		logger.Log.Debug().
			Str("package", "router").
			Msgf("error retrieving posts by user id: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().
		Str("package", "router").
		Msgf("successfully retrieved posts by user id: %v", userID)

	var data []APIResponseData
	for _, post := range posts {
		data = append(data, createAPIResponseDataPost(&post))
	}

	return c.JSON(APIResponse{
		Data: data,
	})
}

// createPost handles the HTTP request to create a new post.
func createPost(c *fiber.Ctx) error {
	var body struct {
		Text    string         `json:"text"`
		Privacy shared.Privacy `json:"privacy"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidJSON,
		})
	}
	logger.Log.Debug().Msgf("body: %v", body)

	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().Msgf("userID: %v", userID)

	post, err := db.Posts.CreatePost(userID, body.Text, body.Privacy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	// Retrieve the post so we can return the author's information.
	post, err = db.Posts.GetPost(post.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Data: []APIResponseData{createAPIResponseDataPost(post)},
	})
}

// deletePost handles the HTTP request to delete a post.
func deletePost(c *fiber.Ctx) error {
	// Parse the post ID from the URL parameter.
	postID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidID,
		})
	}

	// Get the user ID from the JWT.
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInvalidToken,
		})
	}

	// First, check if the post exists and belongs to the user.
	post, err := db.Posts.GetPost(postID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(APIResponse{
				Error: ErrNotFound,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	// Check if the user owns the post.
	if post.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(APIResponse{
			Error: ErrForbidden,
		})
	}

	// Delete the post.
	err = db.Posts.DeletePost(postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// createAPIResponseDataPost converts post to APIResponseData.
func createAPIResponseDataPost(post *database.Post) APIResponseData {
	return APIResponseData{
		Type: shared.DataTypePost,
		ID:   &post.ID,
		Attributes: APIResponseDataAttributes{
			CreatedAt: &post.CreatedAt,
			UpdatedAt: &post.UpdatedAt,
			Text:      post.Text,
			Privacy:   post.Privacy,
		},
		Links: APIResponseDataLinks{
			Self:   "/posts/" + post.ID.String(),
			Author: "/users/" + post.AuthorID.String(),
		},
	}
}
