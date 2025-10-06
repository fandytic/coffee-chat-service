package handler

import (
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"

	"github.com/gofiber/fiber/v2"
)

type QRCodeHandler struct {
	QRCodeService interfaces.QRCodeServiceInterface
}

func (h *QRCodeHandler) GenerateQRCode(c *fiber.Ctx) error {
	var req model.QRCodeRequest

	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Content == "" {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Content field is required")
	}

	pngBytes, err := h.QRCodeService.GenerateQRCode(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate QR code")
	}

	// Set header content type ke image/png dan kirim byte gambar
	c.Set("Content-Type", "image/png")
	return c.Send(pngBytes)
}
