package handler

import (
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	DashboardService interfaces.DashboardServiceInterface
}

func (h *DashboardHandler) GetStats(c *fiber.Ctx) error {
	// Panggil metode GetStats() dari service, bukan dari repository
	stats, err := h.DashboardService.GetStats()
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve dashboard statistics")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Dashboard statistics retrieved successfully", stats)
}
