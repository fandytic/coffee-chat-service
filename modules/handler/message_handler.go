package handler

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type MessageHandler struct {
	MessageService interfaces.MessageServiceInterface
}

func (h *MessageHandler) SendMessage(c *fiber.Ctx) error {
	var req model.SendMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if req.User == "" || req.Text == "" {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "User and text fields are required")
	}

	message, err := h.MessageService.SaveAndBroadcastMessage(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save message")
	}

	return model.SuccessResponse(c, fiber.StatusCreated, "Message sent successfully", message)
}

func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	messages, err := h.MessageService.GetAllMessages()
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve messages")
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Messages retrieved successfully", messages)
}

// Upgrade adalah middleware untuk memeriksa apakah koneksi boleh di-upgrade ke WebSocket.
func (h *MessageHandler) Upgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
