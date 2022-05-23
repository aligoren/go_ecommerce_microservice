package routes

import (
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/config"
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"models,omitempty"`
}

func Get(ctx *fiber.Ctx) error {

	headers := ctx.GetReqHeaders()
	service := headers["Service"]
	requestType := headers["Req-Type"]

	if err := config.IsPathValid(service, ctx); err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Error:      true,
			Data:       nil,
		})
	}

	switch service {
	case "auth":
		return AuthGetHandler(requestType, config.JsonConfig, ctx)
	default:
		return ctx.Status(http.StatusNotAcceptable).JSON(models.ResponseModel{
			StatusCode: http.StatusNotAcceptable,
			Message:    "Missed headers!",
			Error:      true,
			Data:       nil,
		})
	}
}
