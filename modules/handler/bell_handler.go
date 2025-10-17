package handler

import (
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
	"coffee-chat-service/modules/utils"

	"github.com/gofiber/fiber/v2"
)

type BellHandler struct {
	BellService *usecase.BellUseCase
}

func (h *BellHandler) CallWaiter(c *fiber.Ctx) error {
	customerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	if err := h.BellService.CallWaiter(customerID); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to process call")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Waiter has been called", nil)
}
