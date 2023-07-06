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

func (pr *PostsRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/:postId", pr.get)
	rtr.Post("", pr.create)
	rtr.Put("/:postId", pr.update)
	rtr.Delete("/:postId", pr.delete)
}

// @Summary Get a post
// @Description Retrieve a post
// @Tags Posts
// @Accept json
// @Produce json
// @Param postId path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /posts/{:postId} [get]
func (pr *PostsRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "get")

	// Parse the post ID from the URL parameter as a UUID
	postID, err := uuid.Parse(c.Params("postId"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing post id")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send the request to the service layer
	return pr.Srv.Posts.Get(c, &postID)
}

// @Summary Create a post
// @Description Create a new post.
// @Tags Posts
// @Accept json
// @Produce json
// @Param text body string true "The text of the post"
// @Param privacy body string true "The privacy setting of the post"
// @Param mediaIds body []string false "The UUIDs of the media files attached to the post"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /posts [post]
func (pr *PostsRepo) create(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "create")

	// Parse the request body.
	var body struct {
		Text     string         `json:"text"`
		Privacy  shared.Privacy `json:"privacy"`
		MediaIDs []string       `json:"mediaIds"`
	}
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

	// Send the request to the service layer.
	return pr.Srv.Posts.Create(c, body.Text, body.Privacy, mediaIDs)
}

// @Summary Update a post
// @Description Update an existing post.
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /posts/{id} [put]
func (pr *PostsRepo) update(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "update")
	return nil
}

// @Summary Delete a post
// @Description Delete a post.
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /posts/{id} [delete]
func (pr *PostsRepo) delete(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "delete")
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
