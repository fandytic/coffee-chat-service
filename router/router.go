package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"coffee-chat-service/modules/handler"
	ws "coffee-chat-service/modules/websocket"
)

func SetupRoutes(app *fiber.App, messageHandler *handler.MessageHandler, hub *ws.Hub) {
	// Middleware untuk logging
	app.Use(logger.New())

	// Endpoint REST API
	app.Post("/send", messageHandler.SendMessage)
	app.Get("/messages", messageHandler.GetMessages)

	// Endpoint untuk koneksi WebSocket
	app.Use("/ws", messageHandler.Upgrade)
	app.Get("/ws", websocket.New(ws.ServeWs(hub)))
}
