package router

import (
	"time"

	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// login handles the HTTP request for user login.
func login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
			Error: ErrInvalidJSON,
		})
	}
	logger.Log.Debug().Interface("body", body).Msg("Parsed body")

	user, err := db.Users.VerifyUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().Interface("user", user).Msg("Verified user")

	// Create claims
	expiryDate := time.Now().Add(time.Hour * 24 * 7)
	claims := jwt.RegisteredClaims{
		Subject:   user.ID.String(),
		ExpiresAt: jwt.NewNumericDate(expiryDate),
	}
	logger.Log.Debug().Interface("claims", claims).Msg("Created claims")

	// Create token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.Server.JWTSecret))
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().Str("token", token).Msg("Created token")

	return c.JSON(APIResponse{
		Data: []APIResponseData{
			{
				Type: shared.DataTypeToken,
				ID:   &user.ID,
				Attributes: APIResponseDataAttributes{
					Token:     token,
					ExpiresAt: expiryDate.Format(time.RFC3339), // adding expiryDate to the response
				},
				Relationships: APIResponseDataRelationships{
					OwnerID: &user.ID,
				},
			},
		},
	})
}

// verifyToken handles the HTTP request for token verification.
func verifyToken(c *fiber.Ctx) error {
	id, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(APIResponse{
			Error: ErrInternal,
		})
	}
	logger.Log.Debug().Str("id", id.String()).Msg("Verified token")

	return c.JSON(APIResponse{
		Data: []APIResponseData{
			{
				Type: shared.DataTypeToken,
				Attributes: APIResponseDataAttributes{
					Token: c.Get("Authorization"),
				},
				Relationships: APIResponseDataRelationships{
					OwnerID: &id,
				},
			},
		},
	})
}
