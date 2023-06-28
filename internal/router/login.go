package router

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// login handles the HTTP request for user login.
func login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidJSON,
		})
	}

	user, err := db.Users.VerifyUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(cfg.Server.JWTDuration)).Unix()

	t, err := token.SignedString([]byte(cfg.Server.JWTSecret))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to sign token")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Success: false,
			Error:   ErrInternal,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data: []APIResponseData{
			{
				ID: user.ID,
				Attributes: APIResponseDataAttributes{
					Token: t,
				},
			},
		},
	})
}

// verifyToken handles the HTTP request for token verification.
func verifyToken(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Success: false,
			Error:   ErrInvalidToken,
		})
	}

	return c.JSON(APIResponse{
		Success: true,
		Data: []APIResponseData{
			{
				ID:         id,
				Attributes: APIResponseDataAttributes{Token: c.Locals("token").(string)},
			},
		},
	})
}
