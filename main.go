package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"coffee-chat-service/config"
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/handler"
	"coffee-chat-service/modules/repository"
	"coffee-chat-service/modules/usecase"
	ws "coffee-chat-service/modules/websocket"
	"coffee-chat-service/router"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(&entity.Message{}, &entity.Admin{}, &entity.Floor{},
		&entity.Table{}, &entity.Customer{}, &entity.ChatMessage{}, &entity.Menu{})

	// Seeder untuk membuat admin default jika belum ada
	createDefaultAdmin(db)

	hub := ws.NewHub(db)
	go hub.Run()

	// Inisialisasi Repositories
	messageRepo := repository.NewMessageRepository(db)
	adminRepo := repository.NewAdminRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	chatRepo := repository.NewChatRepository(db)
	floorPlanRepo := repository.NewFloorPlanRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)
	menuRepo := repository.NewMenuRepository(db)

	// Inisialisasi Use Cases
	messageUseCase := &usecase.MessageUseCase{Repo: messageRepo, Hub: hub}
	authUseCase := &usecase.AuthUseCase{AdminRepo: adminRepo}
	qrCodeUseCase := &usecase.QRCodeUseCase{}
	floorPlanUseCase := &usecase.FloorPlanUseCase{FloorPlanRepo: floorPlanRepo}
	imageUploadUseCase := &usecase.ImageUploadUseCase{}
	customerUseCase := &usecase.CustomerUseCase{CustomerRepo: customerRepo}
	dashboardUseCase := &usecase.DashboardUseCase{DashboardRepo: dashboardRepo}
	chatUseCase := &usecase.ChatUseCase{ChatRepo: chatRepo}
	menuUseCase := &usecase.MenuUseCase{MenuRepo: menuRepo}

	// Inisialisasi Handlers
	messageHandler := &handler.MessageHandler{MessageService: messageUseCase}
	authHandler := &handler.AuthHandler{AuthService: authUseCase}
	qrCodeHandler := &handler.QRCodeHandler{QRCodeService: qrCodeUseCase}
	floorPlanHandler := &handler.FloorPlanHandler{FloorPlanService: floorPlanUseCase}
	imageUploadHandler := &handler.ImageUploadHandler{ImageUploadService: imageUploadUseCase}
	customerHandler := &handler.CustomerHandler{CustomerService: customerUseCase}
	dashboardHandler := &handler.DashboardHandler{DashboardService: dashboardUseCase}
	chatHandler := &handler.ChatHandler{ChatService: chatUseCase}
	menuHandler := &handler.MenuHandler{MenuService: menuUseCase}

	app := fiber.New()
	app.Static("/public", "./public")
	router.SetupRoutes(app, messageHandler, authHandler, qrCodeHandler, floorPlanHandler,
		imageUploadHandler, customerHandler, dashboardHandler, chatHandler, menuHandler, hub)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}

// Fungsi untuk membuat admin default
func createDefaultAdmin(db *gorm.DB) {
	var admin entity.Admin
	if err := db.First(&admin, "username = ?", "admin").Error; err == gorm.ErrRecordNotFound {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		newAdmin := entity.Admin{Username: "admin", Password: string(hashedPassword)}
		if err := db.Create(&newAdmin).Error; err != nil {
			log.Fatalf("Failed to create default admin: %v", err)
		}
		log.Println("Default admin user created. Username: admin, Password: password123")
	}
}
