package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// getPost handles the HTTP request to fetch a specific post.
func getPost(c *fiber.Ctx) error {
	postID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	post, err := db.Posts.GetPost(postID)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusNotFound).JSON(APIResponse{
			Success: false,
			Error:   ErrNotFound,
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data:    []APIResponseData{createAPIResponseDataPost(post)},
	})
}

// getPosts handles the HTTP request to fetch all posts by a user.
func getPosts(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Query("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidID,
		})
	}

	posts, err := db.Posts.GetPostsByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	var data []APIResponseData
	for _, post := range posts {
		data = append(data, createAPIResponseDataPost(&post))
	}

	return c.JSON(APIResponse{
		Success: true,
		Data:    data,
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
			Success: false,
			Error:   ErrInvalidJSON,
		})
	}

	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	newPost := database.Post{
		Text:    body.Text,
		Privacy: body.Privacy,
		OwnerID: userID,
	}

	if err := db.Posts.CreatePost(&newPost).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data:    []APIResponseData{createAPIResponseDataPost(&newPost)},
	})
}

// deletePost handles the HTTP request to delete a post.
func deletePost(c *fiber.Ctx) error {
	return nil
}

// createAPIResponseDataPost converts post to APIResponseData.
func createAPIResponseDataPost(post *database.Post) APIResponseData {
	return APIResponseData{
		Type: shared.DataTypePost,
		ID:   post.ID,
		Attributes: APIResponseDataAttributes{
			Text: post.Text,
		},
		Relationships: APIResponseDataRelationships{
			OwnerID: post.OwnerID,
		},
	}
}
