package handlers

import (
	"sportmen/database"
	"sportmen/middleware"
	"sportmen/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetSportId(c *fiber.Ctx) error {

	//

	user := c.Locals("requestAuth")

	details, _ := user.(middleware.DeserializeUser)

	if details.Role != 2 {
		return c.Status(401).JSON(fiber.Map{"message": "no eres partner"})
	}

	//

	Sportmen := new(model.SportmenDelete)

	if err := c.BodyParser(Sportmen); err != nil {
		return err
	}

	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(Sportmen)

	if err1 != nil {

		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "field validation error"})

	}

	//

	sport := model.Sportmen{
		UserId: Sportmen.UserId,
	}

	db := database.DB

	db.Preload("Sport").First(&sport)

	return c.Status(200).JSON(fiber.Map{"sportmen": sport.SportmenWithoutId})
}
