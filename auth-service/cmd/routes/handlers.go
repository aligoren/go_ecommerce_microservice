package routes

import (
	"github.com/aligoren/go_ecommerce_microservice/auth-service/models"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetUserByID(ctx *fiber.Ctx) error {

	id, _ := ctx.ParamsInt("id", 0)

	user, err := repository.GetOne(id)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "User not found",
			Error:      true,
			Data:       nil,
		})
	}

	if user.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "User not found",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

func GetAllUsers(ctx *fiber.Ctx) error {

	users, err := repository.GetAllUsers()

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "User not found",
			Error:      true,
			Data:       err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(users)
}
