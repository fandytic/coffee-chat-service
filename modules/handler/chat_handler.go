package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
)

type ChatHandler struct {
	ChatService *usecase.ChatUseCase
}

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

func (h *ChatHandler) GetMessageHistory(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	loggedInCustomerID := uint(claims["customer_id"].(float64))

	otherCustomerID, err := c.ParamsInt("id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid customer ID")
	}

	history, err := h.ChatService.GetMessageHistory(loggedInCustomerID, uint(otherCustomerID))
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve chat history")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Chat history retrieved successfully", history)
}
