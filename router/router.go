package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"coffee-chat-service/modules/handler"
	"coffee-chat-service/modules/middleware"
	ws "coffee-chat-service/modules/websocket"
)

func SetupRoutes(app *fiber.App, messageHandler *handler.MessageHandler,
	authHandler *handler.AuthHandler, qrCodeHandler *handler.QRCodeHandler,
	floorPlanHandler *handler.FloorPlanHandler, imageUploadHandler *handler.ImageUploadHandler,
	customerHandler *handler.CustomerHandler, dashboardHandler *handler.DashboardHandler,
	chatHandler *handler.ChatHandler, menuHandler *handler.MenuHandler, orderHandler *handler.OrderHandler,
	hub *ws.Hub, db *gorm.DB) {
	// Middleware untuk logging
	app.Use(logger.New())

	// Endpoint REST API
	app.Post("/login", authHandler.Login)
	app.Post("/check-in", customerHandler.CheckIn)
	app.Get("/messages", messageHandler.GetMessages)

	app.Post("/upload-image", imageUploadHandler.UploadImage)
	// Endpoint untuk koneksi WebSocket
	app.Use("/ws", messageHandler.Upgrade)
	app.Get("/ws", func(c *fiber.Ctx) error {
		return handler.HandleWebSocketConnection(hub, c)
	})

	adminProtected := app.Group("/admin", middleware.AdminProtected())
	adminProtected.Post("/logout", authHandler.Logout)

	adminProtected.Post("/send", messageHandler.SendMessage)
	adminProtected.Post("/generate-qr", qrCodeHandler.GenerateQRCode)

	adminProtected.Post("/floor-plans", floorPlanHandler.CreateFloorPlan)
	adminProtected.Get("/floor-plans", floorPlanHandler.GetAllFloors)
	adminProtected.Delete("/floor-plans/:floor_id", floorPlanHandler.DeleteFloor)
	adminProtected.Get("/floor-plans/:floor_number", floorPlanHandler.GetFloorPlan)

	adminProtected.Put("/tables/:table_id", floorPlanHandler.UpdateTable)
	adminProtected.Delete("/tables/:table_id", floorPlanHandler.DeleteTable)

	adminProtected.Get("/dashboard/stats", dashboardHandler.GetStats)

	adminProtected.Post("/menus", menuHandler.CreateMenu)
	adminProtected.Get("/menus", menuHandler.GetAllMenus)
	adminProtected.Get("/menus/:id", menuHandler.GetMenuByID)
	adminProtected.Put("/menus/:id", menuHandler.UpdateMenu)
	adminProtected.Delete("/menus/:id", menuHandler.DeleteMenu)

	adminProtected.Get("/orders", orderHandler.GetAllOrders)

	adminProtected.Get("/customers", customerHandler.GetAllCustomers)

	adminProtected.Delete("/customers/:id", customerHandler.RevokeCustomerAccess)

	customerProtected := app.Group("/customer", middleware.CustomerProtected(db))
	customerProtected.Get("/active-list", customerHandler.GetActiveCustomers)
	customerProtected.Get("/stats", dashboardHandler.GetStats)
	customerProtected.Get("/floor-plans/:floor_number", floorPlanHandler.GetFloorPlan)
	customerProtected.Get("/floor-plans", floorPlanHandler.GetAllFloors)

	customerProtected.Post("/chats/:sender_id/mark-as-read", chatHandler.MarkMessagesAsRead)
	customerProtected.Get("/chats/:id", chatHandler.GetMessageHistory)

	customerProtected.Get("/menus", menuHandler.GetAllMenus)
	customerProtected.Post("/orders", orderHandler.CreateOrder)

	customerProtected.Get("/wishlists/:id", orderHandler.GetWishlistDetails)
	customerProtected.Post("/wishlists/:id/accept", orderHandler.AcceptWishlist)
}
