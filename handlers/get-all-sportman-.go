package handlers

import (
	"sportmen/database"
	"sportmen/middleware"
	"sportmen/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllSport(c *fiber.Ctx) error {

	//

	user := c.Locals("requestAuth")

	details, _ := user.(middleware.DeserializeUser)

	if details.Role != 2 {
		return c.Status(401).JSON(fiber.Map{"message": "no eres partner"})
	}

	var sport []model.Sportmen

	db := database.DB

	db.Preload("Sport").Find(&sport)

	return c.Status(200).JSON(fiber.Map{"sportmen": sport})
}
