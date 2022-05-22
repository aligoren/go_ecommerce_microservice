package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"models,omitempty"`
}

func Get(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(Response{
		Success: true,
		Message: "It works!",
		Data:    nil,
	})
}
