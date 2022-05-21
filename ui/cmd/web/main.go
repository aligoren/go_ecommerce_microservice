package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"log"
)

func main() {

	engine := django.New("./cmd/web/templates", ".html")

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
