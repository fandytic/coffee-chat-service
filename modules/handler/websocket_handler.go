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

	customerIDFloat, ok := claims["customer_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid customer_id in token")
	}
	customerID := uint(customerIDFloat)

	return websocket.New(ws.ServeWs(hub, customerID))(c)
}
