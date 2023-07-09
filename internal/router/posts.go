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
	rtr.Get("/", pr.search)
	rtr.Post("/", pr.create)
	rtr.Put("/:postID", pr.update)
	rtr.Delete("/:postID", pr.delete)
	rtr.Post("/:postID/bookmarks", pr.createBookmark)
	rtr.Delete("/:postID/bookmarks", pr.deleteBookmark)
	rtr.Post("/:postID/likes", pr.createLike)
	rtr.Delete("/:postID/likes/", pr.deleteLike)
}

// @Summary Search Posts
// @Description Search for posts with query parameters.
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param postID query string false "Post ID"
// @Param userID query string false "User ID"
// @Param count query string false "The number of posts to return."
// @Param offset query string false "The number of posts to offset the returned posts by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts [get]
func (pr *PostsRepo) search(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "search")
	return nil
}

// @Summary Retrieve Post by ID
// @Description Retrieves a post by ID.
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/{postID} [get]
func (pr *PostsRepo) getByID(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "getByID")
	return nil
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
// @Param text formData string true "The text of the post"
// @Param privacy body string true "The privacy setting of the post"
// @Param privacy formData string true "The privacy setting of the post"
// @Param mediaIDs body []string false "The UUIDs of the media files attached to the post"
// @Param mediaIDs formData []string false "The UUIDs of the media files attached to the post"
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

	// Send the request to the service layer.
	return pr.Srv.Posts.Delete(c, postIDPtr)
}

// @Summary Bookmark a Post
// @Description Adds the specified post to the authenticated user's bookmarks.
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /posts/{postID}/bookmarks [post]
func (pr *PostsRepo) createBookmark(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "createBookmark")

	// Get the post ID from the URL parameter.
	postID, err := uuid.Parse(c.Params("postID"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing post id")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send the request to the service layer.
	return pr.Srv.Bookmarks.Create(c, &postID)
}

// @Summary Unbookmark a Post
// @Description Removes the specified post from the authenticated user's bookmarks.
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /posts/{postID}/bookmarks [delete]
func (pr *PostsRepo) deleteBookmark(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "deleteBookmark")
	return nil
}

// @Summary Like a Post
// @Description Create a new like for a post.
// @Tags Likes
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/{postID}/likes [post]
func (pr *PostsRepo) createLike(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "createLike")

	// Get the post ID from the URL parameter.
	postID, err := uuid.Parse(c.Params("postID"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing post id")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send the request to the service layer.
	return pr.Srv.Likes.Create(c, &postID)
}

// @Summary Unlike a Post
// @Description Delete a like for a post.
// @Tags Likes
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Param likeID path string true "Like ID"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /posts/{postID}/likes/ [delete]
func (pr *PostsRepo) deleteLike(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "posts", "deleteLike")

	// Get the post ID from the URL parameter.
	postID, err := uuid.Parse(c.Params("postID"))
	if err != nil {
		logger.Log.Error().Err(err).Msg("error parsing post id")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
	}

	// Send the request to the service layer.
	return pr.Srv.Likes.Delete(c, &postID)
}
