package routes

import (
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/middleware"
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/routes/handlers/auth"
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

	api := app.Group("/api", middleware.HeartBeat("/ping"))

	api.Get("/users", auth.GetAllUsers)
	api.Get("/users/:id", auth.GetUserByID)
	api.Post("/users", auth.CreateUser)
	api.Put("/users", auth.UpdateUser)
	api.Delete("/users", auth.DeleteUser)
	api.Post("/auth", auth.LoginUser)
}
