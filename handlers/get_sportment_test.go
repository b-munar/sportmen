package handlers

import (
	"net/http/httptest"
	"sportmen/database"

	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func TestSportmenGet(t *testing.T) {

	database.Migrate()

	app := fiber.New()

	app.Get("/sportmen", GetSport)

	req := httptest.NewRequest("GET", "/sportmen", nil)

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIyOGEzYWQ3Ny03ZDNjLTQ3ZTMtOWM0Mi04NThjYTNlYzUyMjIiLCJpYXQiOjE3MTMxNDk5OTEsImV4cCI6MTcxMzE2MDc5MSwicm9sZSI6Mn0.Z2owpDEptx6mNuHm63QJzFKZoxK-7XFVgxasb5KKQF0")

	resp, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "app.Test 2")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

}
