package handler

import (
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	// Ambil data user dari token JWT yang sudah divalidasi oleh middleware
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	customerIDClaim, ok := claims["customer_id"]
	if !ok {
		// Jika field customer_id tidak ada di token, berarti ini bukan token pelanggan
		return model.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Access is restricted to customers only")
	}

	customerIDFloat, ok := customerIDClaim.(float64)
	if !ok {
		// Jika tipe datanya salah, token mungkin rusak atau tidak valid
		return model.ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("Invalid customer_id type in token: %T", customerIDClaim))
	}

	loggedInCustomerID := uint(customerIDFloat)

	// Teruskan ID pelanggan yang login ke service
	customers, err := h.CustomerService.GetActiveCustomers(loggedInCustomerID)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve active customers")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Active customers retrieved successfully", customers)
}
