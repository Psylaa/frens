package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostsRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewPostsRepo(db *database.Database, srv *service.Service) *PostsRepo {
	return &PostsRepo{
		DB:  db,
		Srv: srv,
	}
}

func (pr *PostsRepo) ConfigureRoutes(rtr fiber.Router) {}

// getPost handles the HTTP request to fetch a specific post.
func (pr *PostsRepo) getPostByID(c *fiber.Ctx) error {
	return nil
	/*
		postID, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}
		logger.Log.Debug().
			Str("package", "router").
			Str("func", "getPost").
			Msgf("successfully parsed provided post id into uuid: %v", postID)

		post, err := db.Posts.GetPost(postID)
		switch err {
		case nil:
			break
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
		default:
			logger.Log.Error().Err(err).
				Str("package", "router").
				Str("func", "getPost").
				Msg("unknown error fetching post from database")
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.JSON(response.GeneratePostResponse(post))
	*/
}

func (pr *PostsRepo) getPostsByUserID(c *fiber.Ctx) error {
	logger.DebugLogRequestRecieved("router", "posts", "getPostsByUserID")

	// Get the user id from the request
	userID, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing userID: %v")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}
	logger.Log.Debug().Msgf("successfully parsed provided user id into uuid: %v", userID)

	// Check to make sure userID is not empty
	if userID == uuid.Nil {
		logger.Log.Error().Msg("userID was parsed as nil")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send to service package
	return pr.Srv.Posts.GetByUserID(c, &userID)
}

// getPosts handles the HTTP request to fetch all posts by a user.
func (pr *PostsRepo) GetPostsByUserID(c *fiber.Ctx) error {
	/*
		userID, err := uuid.Parse(c.Query("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}
		logger.Log.Debug().Msgf("successfully parsed provided user id into uuid: %v", userID)

		posts, err := db.Posts.GetPostsByUserID(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
		return c.JSON(response.GeneratePostsResponse(posts))
	*/
	return nil
}

// createPost handles the HTTP request to create a new post.
func (pr *PostsRepo) createPost(c *fiber.Ctx) error {
	var body struct {
		Text     string         `json:"text"`
		Privacy  shared.Privacy `json:"privacy"`
		MediaIDs []string       `json:"mediaIds"`
	}

	// Parse the request body.
	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	// Convert the media IDs to UUIDs.
	var mediaIDs []*uuid.UUID
	for _, id := range body.MediaIDs {
		mediaID, err := uuid.Parse(id)
		if err != nil {
			logger.Log.Error().Err(err).Interface("id", id).Msg("error parsing media id")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidMediaUUID))
		}
		mediaIDs = append(mediaIDs, &mediaID)
	}

	return pr.Srv.Posts.Create(c, body.Text, body.Privacy, mediaIDs)
}

// deletePost handles the HTTP request to delete a post.
func (pr *PostsRepo) deletePost(c *fiber.Ctx) error {
	/*
		// Parse the post ID from the URL parameter.
		postID, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		// Get the user ID from the JWT.
		userID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		}

		// First, check if the post exists and belongs to the user.
		post, err := db.Posts.GetPost(postID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
			}
		}

		// Check if the user owns the post.
		if post.AuthorID != userID {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrUnauthorized))
		}

		// Delete the post.
		err = db.Posts.DeletePost(postID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}
	*/
	return c.SendStatus(fiber.StatusNoContent)
}
