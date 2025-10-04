package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"coffee-chat-service/config"
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/handler"
	"coffee-chat-service/modules/repository"
	"coffee-chat-service/modules/usecase"
	ws "coffee-chat-service/modules/websocket"
)

func main() {
	// Inisialisasi Database
	db := config.InitDB()
	db.AutoMigrate(&entity.Message{})

	// Inisialisasi komponen aplikasi
	hub := ws.NewHub()
	go hub.Run()

	messageRepo := repository.NewMessageRepository(db)
	messageUseCase := &usecase.MessageUseCase{
		Repo: messageRepo,
		Hub:  hub,
	}
	messageHandler := &handler.MessageHandler{
		MessageService: messageUseCase,
	}

	// Inisialisasi Fiber
	app := fiber.New()
	app.Use(logger.New())

	// Set up the routes
	app.Post("/send", messageHandler.SendMessage)
	app.Get("/messages", messageHandler.GetMessages)

	// Route untuk WebSocket
	app.Use("/ws", messageHandler.Upgrade)
	app.Get("/ws", websocket.New(ws.ServeWs(hub)))

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
