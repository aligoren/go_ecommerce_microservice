package routes

import (
	"fmt"
	"github.com/aligoren/go_ecommerce_microservice/log-service/models"
	"github.com/aligoren/go_ecommerce_microservice/log-service/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func CreateLog(ctx *fiber.Ctx) error {

	logModel := new(models.LogModel)

	if err := ctx.BodyParser(logModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Request body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	err := repository.InsertLog(models.LogModel{
		Name:      logModel.Name,
		Data:      logModel.Data,
		CreatedAt: time.Time{},
	})

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Error while inserting log %v", err),
			Error:      true,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Log created"),
		Error:      false,
	})
}
