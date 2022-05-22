package main

import (
	"github.com/aligoren/go_ecommerce_microservice/auth-service/cmd/routes"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const port = ":80"

func main() {

	env := os.Getenv("ENV")

	if env == "" || env == "DEVELOPMENT" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf(".env file couldn't loaded %v", env)
		}
	}

	database.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}
