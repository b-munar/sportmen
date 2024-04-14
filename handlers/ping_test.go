package handlers

import (
	"net/http/httptest"

	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/utils"
)

// go test -run -v Test_Handler
func TestPing(t *testing.T) {

	app := fiber.New()

	app.Get("/sportmen/ping", Ping)

	resp, err := app.Test(httptest.NewRequest("GET", "/sportmen/ping", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
