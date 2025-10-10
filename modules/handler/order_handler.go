package handler

import (
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	// ...
)

type OrderHandler struct {
	OrderService *usecase.OrderUseCase
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	customerID := uint(claims["customer_id"].(float64))

	var req model.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	order, err := h.OrderService.CreateOrder(customerID, req)
	if err != nil {
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
