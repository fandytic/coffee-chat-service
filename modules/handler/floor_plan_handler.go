package handler

import (
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/usecase"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FloorPlanHandler struct {
	FloorPlanService *usecase.FloorPlanUseCase
}

func (h *FloorPlanHandler) CreateFloorPlan(c *fiber.Ctx) error {
	// Ambil file dari form
	file, err := c.FormFile("floor_plan_image")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "floor_plan_image is required")
	}

	// Ambil field lain dari form
	floorNumberStr := c.FormValue("floor_number")
	tablesDataJSON := c.FormValue("tables")

	if floorNumberStr == "" || tablesDataJSON == "" {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "floor_number and tables are required")
	}

	floorNumber, err := strconv.Atoi(floorNumberStr)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "floor_number must be an integer")
	}

	response, err := h.FloorPlanService.CreateFloorPlan(floorNumber, file, tablesDataJSON)
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
	// Ambil ID meja dari parameter URL
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
		// Cek apakah error karena data tidak ditemukan
		if err.Error() == fmt.Sprintf("table with ID %d not found", tableID) {
			return model.ErrorResponse(c, fiber.StatusNotFound, err.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update table")
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Table updated successfully", updatedTable)
}

// DeleteTable menangani permintaan untuk menghapus meja.
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

	// Untuk DELETE, respons sukses biasanya tidak memiliki data (No Content)
	// Tapi kita akan kembalikan pesan sukses agar konsisten
	return model.SuccessResponse(c, fiber.StatusOK, "Table deleted successfully", nil)
}
