package model

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(Response{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(Response{
		Success: false,
		Code:    code,
		Message: message,
	})
}
