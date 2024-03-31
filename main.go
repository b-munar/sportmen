package main

import (
	"sportmen/database"
	"sportmen/handlers"
	"sportmen/middleware"
	"sportmen/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDB()

	db := database.DB

	db.AutoMigrate(&model.Sportmen{}, &model.SportmenSport{})

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/sportmen/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("ping")
	})

	app.Post("/sportmen", handlers.CreateSportmen)

	app.Delete("/sportmen", handlers.DeleteSportmen)

	app.Use(middleware.New(middleware.Config{}))

	app.Get("/sportmen", handlers.GetSport)

	app.Listen(":80")
}
