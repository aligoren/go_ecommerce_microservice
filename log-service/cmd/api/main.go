package main

import (
	"github.com/aligoren/go_ecommerce_microservice/log-service/cmd/routes"
	"github.com/aligoren/go_ecommerce_microservice/log-service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf(".env file couldn't loaded %v", err)
	}

	port := os.Getenv("WEB_PORT")

	database.ConnectoToMongoDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}
