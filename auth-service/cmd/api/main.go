package main

import (
	"github.com/aligoren/go_ecommerce_microservice/auth-service/cmd/routes"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

const port = ":80"

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file couldn't loaded")
	}

	database.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}
