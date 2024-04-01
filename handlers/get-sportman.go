package handlers

import (
	"sportmen/database"
	"sportmen/middleware"
	"sportmen/model"

	"github.com/gofiber/fiber/v2"
)

func GetSport(c *fiber.Ctx) error {

	user := c.Locals("requestAuth")

	details, _ := user.(middleware.DeserializeUser)

	sport := model.Sportmen{
		UserId: details.ID,
	}

	db := database.DB

	db.Preload("Sport").First(&sport)

	return c.Status(200).JSON(fiber.Map{"sportmen": sport.SportmenWithoutId})
}
