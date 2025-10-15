package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
	"coffee-chat-service/modules/utils"
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
	loggedInCustomerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	floorNumber, _ := strconv.Atoi(c.Query("floor")) // Konversi string ke integer
	filter := model.CustomerFilter{
		Search:      c.Query("search"),
		FloorNumber: floorNumber,
		TableNumber: c.Query("table"),
	}

	customers, err := h.CustomerService.GetActiveCustomers(loggedInCustomerID, filter)
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
