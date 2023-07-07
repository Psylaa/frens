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
	rtr.Get("/:postID", pr.get)
	rtr.Post("", pr.create)
	rtr.Put("/:postID", pr.update)
	rtr.Delete("/:postID", pr.delete)
}

// @Summary Get a post
// @Description Retrieve a post
// @Tags Posts
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /posts/{postID} [get]
func (pr *PostsRepo) get(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "get")

	// Parse the post ID from the URL parameter as a UUID
	postID, err := uuid.Parse(c.Params("postID"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing post id")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send the request to the service layer
	return pr.Srv.Posts.Get(c, &postID)
}

type CreatePostRequest struct {
	Text     string         `json:"text"`
	Privacy  shared.Privacy `json:"privacy"`
	MediaIDs []string       `json:"mediaIDs"`
}

// @Summary Create a post
// @Description Create a new post.
// @Tags Posts
// @Accept json
// @Produce json
// @Param text body string true "The text of the post"
// @Param privacy body string true "The privacy setting of the post"
// @Param mediaIDs body []string false "The UUIDs of the media files attached to the post"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /posts [post]
func (pr *PostsRepo) create(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "create")

	// Parse the request body
	var req CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error().Err(err).Msg("error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	// Fill in default values
	if req.Privacy == "" {
		req.Privacy = shared.PrivacyPublic
	}

	//// Validate the request body
	// Ensure one of text or media is provided
	if req.Text == "" && len(req.MediaIDs) == 0 {
		logger.Log.Error().Msg("no text or media provided in request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}
	// Ensure the privacy setting is valid
	if !req.Privacy.IsValid() {
		logger.Log.Error().Msg("invalid privacy setting provided in request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	// Convert the media ID's to UUID's
	mediaUUIDs := make([]*uuid.UUID, len(req.MediaIDs))
	for i, id := range req.MediaIDs {
		mediaID, err := uuid.Parse(id)
		if err != nil {
			logger.Log.Error().Err(err).Msg("error parsing media id")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}
		mediaUUIDs[i] = &mediaID
	}

	// Send the request to the service layer
	return pr.Srv.Posts.Create(c, req.Text, req.Privacy, mediaUUIDs)
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
// @Router /posts/{postID} [put]
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
// @Router /posts/{postID} [delete]
func (pr *PostsRepo) delete(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "delete")
	// Parse the post ID from the URL parameter.
	postID, err := uuid.Parse(c.Params("postID"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing post id")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}
	postIDPtr := &postID

	// Get the user ID from the JWT.
	requestorID := c.Locals("requestorID").(*uuid.UUID)

	// Send the request to the service layer.
	return pr.Srv.Posts.Delete(c, requestorID, postIDPtr)
}
