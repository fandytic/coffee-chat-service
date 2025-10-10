package handler

import (
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	ws "coffee-chat-service/modules/websocket"
)

func HandleWebSocketConnection(hub *ws.Hub, c *fiber.Ctx) error {
	if !websocket.IsWebSocketUpgrade(c) {
		return fiber.ErrUpgradeRequired
	}

	tokenString := c.Query("token")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing auth token")
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	if adminIDFloat, ok := claims["user_id"].(float64); ok {
		adminID := uint(adminIDFloat)

		return websocket.New(ws.ServeAdminWs(hub, adminID))(c)
	}

	if customerIDFloat, ok := claims["customer_id"].(float64); ok {
		customerID := uint(customerIDFloat)

		return websocket.New(ws.ServeCustomerWs(hub, customerID))(c)
	}

	return c.Status(fiber.StatusBadRequest).SendString("Invalid token type for WebSocket")
}
