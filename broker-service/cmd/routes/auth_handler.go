package routes

import (
	"encoding/json"
	"fmt"
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var baseUrl string

func AuthGetHandler(typeName string, jsonConfig map[string]models.ServiceModel, ctx *fiber.Ctx) error {

	baseUrl = jsonConfig["auth"].BaseUrl

	switch typeName {
	case "req-single-user":
		return getUserByID(ctx)
	case "req-all-users":
		return getAllUsers(ctx)
	default:
		return ctx.Status(http.StatusNotAcceptable).JSON(models.ResponseModel{
			StatusCode: http.StatusNotAcceptable,
			Message:    "Missed headers!",
			Error:      true,
			Data:       nil,
		})
	}
}

func getUserByID(ctx *fiber.Ctx) error {

	userID := ctx.Query("id", "0")

	response, err := http.Get(fmt.Sprintf("%s/users/%s", baseUrl, userID))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Error while fetching user info",
			Error:      true,
			Data:       nil,
		})
	}

	var jsonData interface{}
	_ = json.NewDecoder(response.Body).Decode(&jsonData)

	return ctx.Status(response.StatusCode).JSON(jsonData)

}

func getAllUsers(ctx *fiber.Ctx) error {

	response, err := http.Get(fmt.Sprintf("%s/users", baseUrl))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Error while fetching users",
			Error:      true,
			Data:       nil,
		})
	}

	var jsonData interface{}
	_ = json.NewDecoder(response.Body).Decode(&jsonData)

	return ctx.Status(response.StatusCode).JSON(jsonData)

}
