package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
)

type FloorPlanHandler struct {
	FloorPlanService *usecase.FloorPlanUseCase
}

func (h *FloorPlanHandler) CreateFloorPlan(c *fiber.Ctx) error {
	var req model.CreateFloorPlanRequest

	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := h.FloorPlanService.CreateFloorPlan(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusCreated, "Floor plan created successfully", response)
}

func (h *FloorPlanHandler) GetFloorPlan(c *fiber.Ctx) error {
	floorNumber, err := strconv.Atoi(c.Params("floor_number"))
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid floor number")
	}

	response, err := h.FloorPlanService.GetFloorPlanByNumber(floorNumber)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusNotFound, "Floor plan not found")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Floor plan retrieved successfully", response)
}

func (h *FloorPlanHandler) GetAllFloors(c *fiber.Ctx) error {
	response, err := h.FloorPlanService.GetAllFloors()
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve floors")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Floors retrieved successfully", response)
}

func (h *FloorPlanHandler) UpdateTable(c *fiber.Ctx) error {
	tableID, err := c.ParamsInt("table_id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid table ID")
	}

	var req model.UpdateTableRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	updatedTable, err := h.FloorPlanService.UpdateTable(uint(tableID), req)
	if err != nil {
		if err.Error() == fmt.Sprintf("table with ID %d not found", tableID) {
			return model.ErrorResponse(c, fiber.StatusNotFound, err.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update table")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Table updated successfully", updatedTable)
}

func (h *FloorPlanHandler) DeleteTable(c *fiber.Ctx) error {
	tableID, err := c.ParamsInt("table_id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid table ID")
	}

	err = h.FloorPlanService.DeleteTable(uint(tableID))
	if err != nil {
		if err.Error() == fmt.Sprintf("table with ID %d not found", tableID) {
			return model.ErrorResponse(c, fiber.StatusNotFound, err.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete table")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Table deleted successfully", nil)
}

func (h *FloorPlanHandler) DeleteFloor(c *fiber.Ctx) error {
	floorID, err := c.ParamsInt("floor_id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid floor ID")
	}

	err = h.FloorPlanService.DeleteFloor(uint(floorID))
	if err != nil {
		if err.Error() == fmt.Sprintf("floor with ID %d not found", floorID) {
			return model.ErrorResponse(c, fiber.StatusNotFound, err.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Floor and its associated data deleted successfully", nil)
}
