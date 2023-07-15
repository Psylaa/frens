package router

import (
	"github.com/bwoff11/frens/internal/models"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type PostsRepo struct {
	Service *service.Service
}

func (pr *PostsRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", pr.search)
	rtr.Post("/", pr.create)
	rtr.Get("/:postID", pr.getByID)
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
	return nil
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

	var req models.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return models.ErrInvalidBody.SendResponse(c)
	}

	return pr.Service.Posts.Create(c, &req)
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
	return nil
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
	return nil
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
	return nil
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
	return nil
}
