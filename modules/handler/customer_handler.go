package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
)

type CustomerHandler struct {
	CustomerService *usecase.CustomerUseCase
}

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

func (h *CustomerHandler) GetActiveCustomers(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	customerIDClaim, ok := claims["customer_id"]
	if !ok {
		return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Access is restricted to customers only")
	}

	customerIDFloat, ok := customerIDClaim.(float64)
	if !ok {
		return model.ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("Invalid customer_id type in token: %T", customerIDClaim))
	}

	loggedInCustomerID := uint(customerIDFloat)

	customers, err := h.CustomerService.GetActiveCustomers(loggedInCustomerID)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve active customers")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Active customers retrieved successfully", customers)
}

func (h *CustomerHandler) GetAllCustomers(c *fiber.Ctx) error {
	search := c.Query("search")
	customers, err := h.CustomerService.GetAllCustomers(search)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve customers")
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Customers retrieved successfully", customers)
}
