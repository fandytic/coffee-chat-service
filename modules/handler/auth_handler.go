package handler

import (
	"github.com/gofiber/fiber/v2"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/utils"
)

type AuthHandler struct {
	AuthService interfaces.AuthServiceInterface
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	resp, err := h.AuthService.Login(req)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Login successful", resp)
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return model.SuccessResponse(c, fiber.StatusOK, "Logout successful", nil)
}

// ResetPassword godoc
// @Summary Reset admin password
// @Description Reset admin password using username and new password
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /admin/reset-password [post]
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var req model.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.AuthService.ResetPassword(req); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Password reset successful", nil)
}

// UpdatePassword godoc
// @Summary Update admin password
// @Description Update admin password using old password and new password
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.UpdatePasswordRequest true "Update Password Request"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /admin/update-password [put]
func (h *AuthHandler) UpdatePassword(c *fiber.Ctx) error {
	adminID, err := utils.GetAdminIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	var req model.UpdatePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.AuthService.UpdatePassword(adminID, req); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Password updated successful", nil)
}

// UpdateUsername godoc
// @Summary Update admin username
// @Description Update admin username
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.UpdateUsernameRequest true "Update Username Request"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /admin/update-username [put]
func (h *AuthHandler) UpdateUsername(c *fiber.Ctx) error {
	adminID, err := utils.GetAdminIDFromToken(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	var req model.UpdateUsernameRequest
	if err := c.BodyParser(&req); err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.AuthService.UpdateUsername(adminID, req); err != nil {
		return model.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return model.SuccessResponse(c, fiber.StatusOK, "Username updated successful", nil)
}
