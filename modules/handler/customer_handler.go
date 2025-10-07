package handler

import (
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	CustomerService *usecase.CustomerUseCase
}

// CheckIn menangani proses check-in customer
func (h *CustomerHandler) CheckIn(c *fiber.Ctx) error {
	var req model.CustomerCheckInRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.TableID == 0 || req.Name == "" {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "table_id and name are required")
	}

	resp, err := h.CustomerService.CheckIn(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusCreated, "Check-in successful", resp)
}

// GetActiveCustomers menangani permintaan untuk melihat customer aktif
func (h *CustomerHandler) GetActiveCustomers(c *fiber.Ctx) error {
	customers, err := h.CustomerService.GetActiveCustomers()
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve active customers")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Active customers retrieved successfully", customers)
}
