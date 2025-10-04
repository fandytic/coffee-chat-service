package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"coffee-chat-service/config"
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/handler"
	"coffee-chat-service/modules/repository"
	"coffee-chat-service/modules/usecase"
	ws "coffee-chat-service/modules/websocket"
	"coffee-chat-service/router" // <-- Import paket router baru
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

	// Setup semua routes dari file router.go
	router.SetupRoutes(app, messageHandler, hub) // <-- Panggil fungsi setup router

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
