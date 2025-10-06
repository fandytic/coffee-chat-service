package handler

import (
	"github.com/gofiber/fiber/v2"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type AuthHandler struct {
	AuthService interfaces.AuthServiceInterface
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	resp, err := h.AuthService.Login(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Login successful", resp)
}
