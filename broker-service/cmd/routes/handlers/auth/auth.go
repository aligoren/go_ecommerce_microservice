package auth

import (
	"encoding/json"
	"fmt"
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/config"
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var baseUrl string

const service = "auth"

func init() {

	env := os.Getenv("ENV")

	var fileName string = "services.production"

	if env == "" || env == "DEVELOPMENT" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf(".env file couldn't loaded %v", env)
		}

		fileName = "services"
	}

	config.LoadServicesJson(fileName)

	baseUrl = config.JsonConfig[service].BaseUrl
}

func GetUserByID(ctx *fiber.Ctx) error {

	userID := ctx.Params("id", "0")

	response, err := http.Get(fmt.Sprintf("%s/users/%s", baseUrl, userID))

	if err != nil {

		log.Printf("Path %v, baseUrl: %s", err, baseUrl)
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

func GetAllUsers(ctx *fiber.Ctx) error {

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
