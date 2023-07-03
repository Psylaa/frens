package router

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UsersRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewUsersRepo(db *database.Database, srv *service.Service) *UsersRepo {
	return &UsersRepo{
		DB:  db,
		Srv: srv,
	}
}

func (ur *UsersRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Get("", ur.get)
	rtr.Post("", ur.create)
	rtr.Patch("/:userId", ur.update)
	rtr.Delete("/:userId", ur.delete)
}

// @Summary Get a user by ID
// @Description Fetch a specific user by their ID.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (ur *UsersRepo) get(c *fiber.Ctx) error {
	/*
		users, err := db.Users.GetUsers()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusOK).JSON(response.GenerateUsersResponse(users))
	*/
	return nil
}

// @Summary Create a user
// @Description Create a new user.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/ [post]
func (ur *UsersRepo) create(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		logger.Log.Error().Err(err).Msg("Error parsing request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
	}

	return ur.Srv.Users.Create(c, body.Username, body.Email, body.Password)
}

// @Summary Update a user
// @Description Update a user.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/ [patch]
func (ur *UsersRepo) update(c *fiber.Ctx) error {
	/*
		var body struct {
			Bio              *string `json:"bio"`
			AvatarID *string `json:"avatarId"`
			CoverID     *string `json:"coverId"`
		}

		if err := c.BodyParser(&body); err != nil {
			logger.Log.Error().Err(err).Msg("Error parsing request body")
			return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidBody))
		}

		userId, err := getUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.CreateErrorResponse(response.ErrUnauthorized))
		}

		if body.Bio != nil {
			if err := db.Users.UpdateBio(userId, body.Bio); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating bio")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		if body.AvatarID != nil {
			ppUUID, err := uuid.Parse(*body.AvatarID)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Error parsing AvatarID")
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
			}
			if err := db.Users.UpdateAvatar(userId, &ppUUID); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating profile picture")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		if body.CoverID != nil {
			ciUUID, err := uuid.Parse(*body.CoverID)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Error parsing CoverID")
				return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidUUID))
			}
			if err := db.Users.UpdateCover(userId, &ciUUID); err != nil {
				logger.Log.Error().Err(err).Msg("Error updating cover image")
				return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
			}
		}

		// Retrieve updated user
		user, err := db.Users.GetUser(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
		}

		return c.Status(fiber.StatusOK).JSON(response.GenerateUserResponse(user))
	*/
	return nil
}

// @Summary Delete a user
// @Description Delete a user.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Security ApiKeyAuth
// @Router /users [delete]
func (ur *UsersRepo) delete(c *fiber.Ctx) error {
	return nil
}
