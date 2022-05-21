package main

import (
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

const port = ":80"

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(port))
}
