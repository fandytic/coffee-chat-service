package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"coffee-chat-service/modules/model"
)

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

func AdminProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			if _, ok := claims["user_id"]; !ok {
				return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Admin access required")
			}
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return model.ErrorResponse(c, fiber.StatusBadRequest, "Missing or malformed JWT")
			}
			return model.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired JWT")
		},
	})
}

func CustomerProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			if _, ok := claims["customer_id"]; !ok {
				return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Customer access required")
			}
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return model.ErrorResponse(c, fiber.StatusBadRequest, "Missing or malformed JWT")
			}
			return model.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired JWT")
		},
	})
}
