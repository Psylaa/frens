package router

import (
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UsersRepo struct {
	Service *service.Service
}

func (ur *UsersRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", ur.search)
	rtr.Get("/self", ur.getSelf)
	rtr.Put("/self", ur.update)
	rtr.Get("/:userID", ur.getByID)
	rtr.Get("/:userID/posts", ur.getPosts)
}

// @Summary Search Users
// @Description Search for users with query parameters.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userID query string false "User ID"
// @Param username query string false "Username"
// @Param count query string false "The number of users to return."
// @Param offset query string false "The number of users to offset the returned users by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users [get]
func (ur *UsersRepo) search(c *fiber.Ctx) error {
	return nil
}

// @Summary Get Self
// @Description Get the authenticated user's profile.
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self [get]
func (ur *UsersRepo) getSelf(c *fiber.Ctx) error {
	return nil
}

// @Summary Retrieve User by ID
// @Description Retrieves a user by ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID} [get]
func (ur *UsersRepo) getByID(c *fiber.Ctx) error {
	return nil
}

// @Summary Update User
// @Description Update the authenticated user's profile.
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self [put]
func (ur *UsersRepo) update(c *fiber.Ctx) error {
	return nil
}

// @Summary Delete User
// @Description Delete the authenticated user's profile.
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self [delete]
func (ur *UsersRepo) delete(c *fiber.Ctx) error {
	return nil
}

// @Summary Confirm Delete User
// @Description Confirm the deletion of the authenticated user's profile.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param confirmationCode query string true "Confirmation Code"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self/confirm [delete]
func (ur *UsersRepo) confirmDelete(c *fiber.Ctx) error {
	return nil
}

// @Summary Get Posts by User ID
// @Description Get posts by user ID
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param count query int false "The number of posts to return."
// @Param cursor query string false "Cursor to start the page from."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/posts [get]
func (ur *UsersRepo) getPosts(c *fiber.Ctx) error {
	return nil
}

// @Summary Block User
// @Description Blocks the specified user from interacting with the authenticated user.
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/blocks [post]
func (ur *UsersRepo) block(c *fiber.Ctx) error {
	return nil
}

// @Summary Unblock User
// @Description Removes block on the specified user, allowing them to interact with the authenticated user.
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/blocks [delete]
func (ur *UsersRepo) unblock(c *fiber.Ctx) error {
	return nil
}

// @Summary Retrieve Users Who are Following the Authenticated User
// @Description Retrieves a list of users following the authenticated user.
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param count query string false "The number of follows to return."
// @Param offset query string false "The number of follows to offset the returned follows by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self/followers [get]
func (fr *FollowsRepo) getSelfFollowers(c *fiber.Ctx) error {
	return nil
}

// @Summary Retrieve Users that the Authenticated User is Following
// @Description Retrieves a list of users the authenticated user is following.
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param count query string false "The number of follows to return."
// @Param offset query string false "The number of follows to offset the returned follows by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/self/following [get]
func (fr *FollowsRepo) getSelfFollowing(c *fiber.Ctx) error {
	return nil
}

// @Summary Get Users Who are Following the Specified User
// @Description Get a list of all users that are following a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [get]
func (ur *UsersRepo) getFollowersByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Get Users that the Specified User is Following
// @Description Get a list of all users that a user is following by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/following [get]
func (ur *UsersRepo) getFollowingByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Follow a user by user ID
// @Description Follow a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [post]
func (ur *UsersRepo) followUserByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Unfollow a user by user ID
// @Description Unfollow a user by user ID
// @Tags Follows
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/followers [delete]
func (ur *UsersRepo) unfollowUserByUserID(c *fiber.Ctx) error {
	return nil
}

// @Summary Get likes by user ID
// @Description Get likes by user ID
// @Tags Likes
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param count query string false "The number of likes to return."
// @Param offset query string false "The number of likes to offset the returned likes by."
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{userID}/likes [get]
func (ur *UsersRepo) getLikesByUserID(c *fiber.Ctx) error {
	return nil
}
