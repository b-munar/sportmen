package main

import (
	"sportmen/database"
	"sportmen/handlers"
	"sportmen/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Access-Control-Allow-Credentials",
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowMethods: "GET,POST",
	}))

	app.Use(logger.New())

	app.Get("/sportmen/ping", handlers.Ping)

	app.Post("/sportmen", handlers.CreateSportmen)

	app.Delete("/sportmen", handlers.DeleteSportmen)

	app.Use(middleware.New(middleware.Config{}))

	app.Post("/sportmen/info", handlers.GetSportId)

	app.Get("/sportmen/list", handlers.GetAllSport)

	app.Get("/sportmen", handlers.GetSport)
	app.Listen(":80")

}
