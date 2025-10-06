package middleware

import (
	"os"

	"coffee-chat-service/modules/model"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected adalah middleware untuk memproteksi rute yang memerlukan otentikasi JWT.
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("JWT_SECRET_KEY")),
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return model.ErrorResponse(c, fiber.StatusBadRequest, "Missing or malformed JWT")
			}
			return model.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired JWT")
		},
	})
}
