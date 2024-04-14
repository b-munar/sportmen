package handlers

import (
	"bytes"
	"net/http/httptest"
	"sportmen/database"

	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func TestSportmen(t *testing.T) {

	database.Migrate()

	app := fiber.New()

	contents := []byte(`
		{
		"user": "29a3ad78-6d3c-46e3-9c42-857ca3ec5220",
		"name": "Brahian",
		"last_name": "Munar",
		"age": 26,
		"weight": 63,
		"height": 163,
		"gender": "dog",
		"identification_type":"CC",
		"identification":"314159",
		"country_birth": "Colombia",
		"city_birth": "Cali",
		"country_residence": "Colombia",
		"city_residence": "Elvira",
		"length_residence": 26,
		"sports": [{"sport":"basketball"}]
		}
	`)

	app.Post("/sportmen", CreateSportmen)

	req := httptest.NewRequest("POST", "/sportmen", bytes.NewReader(contents))

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "app.Test 1")
	utils.AssertEqual(t, 201, resp.StatusCode, "Status code")

	app.Delete("/sportmen", DeleteSportmen)

	contents_delete := []byte(`
		{
		"user": "29a3ad78-6d3c-46e3-9c42-857ca3ec5220"
		}
	`)

	req_delete := httptest.NewRequest("DELETE", "/sportmen", bytes.NewReader(contents_delete))

	req_delete.Header.Set("Content-Type", "application/json")

	resp_delete, err_delete := app.Test(req_delete)

	utils.AssertEqual(t, nil, err_delete, "app.Test 2")
	utils.AssertEqual(t, 201, resp_delete.StatusCode, "Status code")

}
