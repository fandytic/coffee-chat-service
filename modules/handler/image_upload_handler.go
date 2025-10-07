package handler

import (
	"github.com/gofiber/fiber/v2"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
)

type ImageUploadHandler struct {
	ImageUploadService *usecase.ImageUploadUseCase
}

func (h *ImageUploadHandler) UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "form file 'image' is required")
	}

	response, err := h.ImageUploadService.SaveImage(file)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Image uploaded successfully", response)
}
