package handler

import (
	"github.com/gofiber/fiber/v2"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type DashboardHandler struct {
	DashboardService interfaces.DashboardServiceInterface
}

func (h *DashboardHandler) GetStats(c *fiber.Ctx) error {
	stats, err := h.DashboardService.GetStats()
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve dashboard statistics")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Dashboard statistics retrieved successfully", stats)
}
