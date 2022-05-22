package routes

import (
	"github.com/aligoren/go_ecommerce_microservice/auth-service/cmd/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-TOKEN",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	app.Use(middleware.HeartBeat("/ping"))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/auth", Login)

	v1.Get("/users", GetAllUsers)
	v1.Get("/users/:id", GetUserByID)
	v1.Post("/users", Create)
	v1.Put("/users", Update)
	v1.Delete("/users", Delete)
}
