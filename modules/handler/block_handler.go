package handler

import (
	"github.com/gofiber/fiber/v2"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
	"coffee-chat-service/modules/utils"
)

type BlockHandler struct {
	BlockService *usecase.BlockUseCase
}

func (h *BlockHandler) BlockCustomer(c *fiber.Ctx) error {
	blockerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, "Invalid token")
	}

	blockedID, err := c.ParamsInt("id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid customer ID")
	}

	if err := h.BlockService.BlockCustomer(blockerID, uint(blockedID)); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to block customer")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Customer blocked successfully", nil)
}

func (h *BlockHandler) UnblockCustomer(c *fiber.Ctx) error {
	blockerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, "Invalid token")
	}

	blockedID, err := c.ParamsInt("id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid customer ID")
	}

	if err := h.BlockService.UnblockCustomer(blockerID, uint(blockedID)); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to unblock customer")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Customer unblocked successfully", nil)
}
