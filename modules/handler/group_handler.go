package handler

import (
	"github.com/gofiber/fiber/v2"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/utils"
)

type GroupHandler struct {
	GroupService interfaces.GroupServiceInterface
}

func NewGroupHandler(groupService interfaces.GroupServiceInterface) *GroupHandler {
	return &GroupHandler{GroupService: groupService}
}

func (h *GroupHandler) CreateGroup(c *fiber.Ctx) error {
	customerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	var req model.CreateGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	group, err := h.GroupService.CreateGroup(customerID, req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusCreated, "Group created successfully", group)
}

func (h *GroupHandler) InviteMembers(c *fiber.Ctx) error {
	customerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	groupID, err := c.ParamsInt("group_id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid group ID")
	}

	var req model.InviteToGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.GroupService.InviteMembers(customerID, uint(groupID), req); err != nil {
		if err.Error() == "forbidden: you are not a member of this group" {
			return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Members invited successfully", nil)
}

func (h *GroupHandler) GetGroupMembers(c *fiber.Ctx) error {
	customerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	groupID, err := c.ParamsInt("group_id")
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid group ID")
	}

	members, err := h.GroupService.GetGroupMembers(customerID, uint(groupID))
	if err != nil {
		if err.Error() == "forbidden: you are not a member of this group" {
			return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
		}
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Group members retrieved successfully", members)
}

func (h *GroupHandler) GetCustomerGroups(c *fiber.Ctx) error {
	customerID, err := utils.GetCustomerIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	groups, err := h.GroupService.GetCustomerGroups(customerID)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Groups retrieved successfully", groups)
}
