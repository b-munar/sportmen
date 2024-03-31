package handlers

import (
	"sportmen/database"
	"sportmen/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateSportmen(c *fiber.Ctx) error {

	Sportmen := new(model.Sportmen)

	if err := c.BodyParser(Sportmen); err != nil {
		return err
	}

	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(Sportmen)

	err2 := validate.Struct(Sportmen.Sport)

	if err1 != nil && err2 != nil {

		return c.Status(400).JSON(fiber.Map{"message": "field validation error"})

	}

	db := database.DB

	var OldSportmen model.Sportmen

	result_1 := db.Find(&OldSportmen, Sportmen.UserId)

	if result_1.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{"message": "uuid"})
	}

	result_2 := db.Create(&Sportmen)

	if result_2.Error != nil {
		return c.Status(400).JSON(fiber.Map{"message": "error in creation"})
	}

	return c.Status(201).JSON(fiber.Map{"sportmen": Sportmen.SportmenWithoutId})
}
