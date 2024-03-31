package handlers

import (
	"sportmen/database"
	"sportmen/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func DeleteSportmen(c *fiber.Ctx) error {

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

	db := database.DB

	var OldSportmenSport model.SportmenSport

	db.Where("user_id = ?", Sportmen.UserId).Delete(&OldSportmenSport)

	var OldSportmen model.Sportmen

	db.Delete(&OldSportmen, Sportmen.UserId)

	return c.Status(201).JSON(fiber.Map{"status": "success"})
}
