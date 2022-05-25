package main

import (
	"github.com/aligoren/go_ecommerce_microservice/ui/cmd/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var baseUrl string

const service = "broker"

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

func main() {

	engine := django.New("./templates", ".html")

	if err := engine.Load(); err != nil {
		log.Fatalf("Template configs couldn't loaded! "+
			"%v", err)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title":   "Shopping Service",
			"Message": "Shopping Service is Working",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
