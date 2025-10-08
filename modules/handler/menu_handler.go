package handler

import (
	"github.com/gofiber/fiber/v2"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type MenuHandler struct {
	MenuService interfaces.MenuServiceInterface
}

func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var req model.MenuRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}
	menu, err := h.MenuService.CreateMenu(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return model.SuccessResponse(c, fiber.StatusCreated, "Menu created successfully", menu)
}

func (h *MenuHandler) GetAllMenus(c *fiber.Ctx) error {
	search := c.Query("search")
	menus, err := h.MenuService.GetAllMenus(search)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Menus retrieved successfully", menus)
}

func (h *MenuHandler) GetMenuByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	menu, err := h.MenuService.GetMenuByID(uint(id))
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusNotFound, "Menu not found")
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Menu retrieved successfully", menu)
}

func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var req model.MenuRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}
	menu, err := h.MenuService.UpdateMenu(uint(id), req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Menu updated successfully", menu)
}

func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.MenuService.DeleteMenu(uint(id)); err != nil {
		return model.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}
	return model.SuccessResponse(c, fiber.StatusOK, "Menu deleted successfully", nil)
}
