package handler

import (
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase" // Ganti dengan interface nanti

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ChatHandler struct {
	ChatService *usecase.ChatUseCase // Ganti dengan interface nanti
}

// MarkMessagesAsRead menangani permintaan untuk menandai pesan sebagai telah dibaca
func (h *ChatHandler) MarkMessagesAsRead(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	recipientID := uint(claims["customer_id"].(float64))

	senderID, err := c.ParamsInt("sender_id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid sender ID")
	}

	if err := h.ChatService.MarkMessagesAsRead(uint(senderID), recipientID); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to mark messages as read")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Messages marked as read", nil)
}
