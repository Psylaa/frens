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
	rtr.Get("/:", pr.get)
	rtr.Post("", pr.create)
	rtr.Put("/:id", pr.update)
	rtr.Delete("/:id", pr.delete)
}

// @Summary Get a post by ID
// @Description Fetch a specific post by its ID.
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /posts/{id} [get]
func (pr *PostsRepo) get(c *fiber.Ctx) error {
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

// @Summary Create a post
// @Description Create a new post.
// @Tags Posts
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /posts [post]
func (pr *PostsRepo) create(c *fiber.Ctx) error {
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
