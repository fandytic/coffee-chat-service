package utils

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetCustomerIDFromToken(c *fiber.Ctx) (uint, error) {
	userClaim := c.Locals("user")
	if userClaim == nil {
		return 0, errors.New("user claim not found in token")
	}

	user := userClaim.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	customerIDClaim, ok := claims["customer_id"]
	if !ok {
		return 0, errors.New("customer_id not found in token")
	}

	customerIDFloat, ok := customerIDClaim.(float64)
	if !ok {
		return 0, fmt.Errorf("invalid customer_id type in token: %T", customerIDClaim)
	}

	return uint(customerIDFloat), nil
}

func GetAdminIDFromToken(c *fiber.Ctx) (uint, error) {
	userClaim := c.Locals("user")
	if userClaim == nil {
		return 0, errors.New("user claim not found in token")
	}

	user := userClaim.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	adminIDClaim, ok := claims["user_id"]
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	adminIDFloat, ok := adminIDClaim.(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user_id type in token: %T", adminIDClaim)
	}

	return uint(adminIDFloat), nil
}
