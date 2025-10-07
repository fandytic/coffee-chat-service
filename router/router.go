package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"coffee-chat-service/modules/handler"
	"coffee-chat-service/modules/middleware"
	ws "coffee-chat-service/modules/websocket"
)

func SetupRoutes(app *fiber.App, messageHandler *handler.MessageHandler,
	authHandler *handler.AuthHandler, qrCodeHandler *handler.QRCodeHandler,
	floorPlanHandler *handler.FloorPlanHandler, imageUploadHandler *handler.ImageUploadHandler,
	hub *ws.Hub) {
	// Middleware untuk logging
	app.Use(logger.New())

	// Endpoint REST API
	app.Post("/login", authHandler.Login)
	app.Get("/messages", messageHandler.GetMessages)
	// Endpoint untuk koneksi WebSocket
	app.Use("/ws", messageHandler.Upgrade)
	app.Get("/ws", websocket.New(ws.ServeWs(hub)))

	protected := app.Group("", middleware.Protected())
	protected.Post("/logout", authHandler.Logout)

	protected.Post("/send", messageHandler.SendMessage)
	protected.Post("/generate-qr", qrCodeHandler.GenerateQRCode)

	protected.Post("/upload-image", imageUploadHandler.UploadImage)

	protected.Post("/floor-plans", floorPlanHandler.CreateFloorPlan)
	protected.Get("/floor-plans", floorPlanHandler.GetAllFloors)
	protected.Delete("/floor-plans/:floor_id", floorPlanHandler.DeleteFloor)
	protected.Get("/floor-plans/:floor_number", floorPlanHandler.GetFloorPlan)

	protected.Put("/tables/:table_id", floorPlanHandler.UpdateTable)
	protected.Delete("/tables/:table_id", floorPlanHandler.DeleteTable)
}
