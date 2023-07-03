package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type FollowsRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewFollowsRepo(db *database.Database, srv *service.Service) *FollowsRepo {
	return &FollowsRepo{
		DB:  db,
		Srv: srv,
	}
}

func (fr *FollowsRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("/", fr.get)
	rtr.Post("/:id", fr.create)
	rtr.Delete("/:id", fr.delete)
}

// @Summary Retrieve user's followers
// @Description Retrieve the followers of a user based on the provided user ID
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /follows/ [get]
func (fr *FollowsRepo) get(c *fiber.Ctx) error {
	/*
		id := c.Params("id")
		userID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		follows, err := db.Follows.GetFollows(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.JSON(response.GenerateFollowsResponse(follows))
	*/
	return nil
}

// @Summary Follow a user
// @Description Create a follow relationship with the target user
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 500
// @Security ApiKeyAuth
// @Router /follows/{userId} [post]
func (fr *FollowsRepo) create(c *fiber.Ctx) error {
	/*
		sourceID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		}

		id := c.Params("id")
		targetID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		// Check if the follower record already exists
		exists, err := db.Follows.DoesFollowExist(sourceID, targetID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		if exists {
			return c.Status(fiber.StatusConflict).JSON(response.CreateErrorResponse(response.ErrExists))
		}

		follow, err := db.Follows.CreateFollow(sourceID, targetID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusCreated).JSON(response.GenerateFollowResponse(follow))
	*/
	return nil
}

// @Summary Unfollow a user
// @Description Delete a follow relationship with the target user
// @Tags Follows
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /follows/{userId} [delete]
func (fr *FollowsRepo) delete(c *fiber.Ctx) error {
	/*
		SourceID, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrInvalidToken))
		}

		id := c.Params("id")
		TargetID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidID))
		}

		// Check if the follower record exists
		exists, err := db.Follows.DoesFollowExist(SourceID, TargetID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(response.CreateErrorResponse(response.ErrNotFound))
		}

		if err := db.Follows.DeleteFollow(SourceID, TargetID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.SendStatus(fiber.StatusOK)
	*/
	return nil
}
