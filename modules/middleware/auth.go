package middleware

import (
	"fmt"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
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

func CustomerProtected(db *gorm.DB) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			customerIDClaim, ok := claims["customer_id"]
			if !ok {
				return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Customer access required")
			}

			customerIDFloat, ok := customerIDClaim.(float64)
			if !ok {
				return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Invalid customer token")
			}
			customerID := uint(customerIDFloat)

			var customer entity.Customer
			if err := db.First(&customer, customerID).Error; err != nil {
				return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Customer not found")
			}
			if customer.Status != "active" {
				return model.ErrorResponse(c, fiber.StatusForbidden, fmt.Sprintf("Forbidden: Customer access has been %s", customer.Status))
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
