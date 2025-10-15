package handler

import (
	"errors"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
	"coffee-chat-service/modules/utils"

	"github.com/gofiber/fiber/v2"
	// ...
)

type OrderHandler struct {
	OrderService *usecase.OrderUseCase
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	customerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	var req model.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	order, err := h.OrderService.CreateOrder(customerID, req)
	if err != nil {
		var validationErr *model.ValidationError
		if errors.As(err, &validationErr) {
			return model.ErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusCreated, "Order created successfully", order)
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.OrderService.GetAllOrders()
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve orders")
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Orders retrieved successfully", orders)
}
