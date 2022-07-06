//go:build !integration
// +build !integration

package controller

import (
	"api/app/lib"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetAPIIndex(t *testing.T) {
	// example connection
	services.DBConnectTest()
	// db := services.DB
	app := fiber.New()
	app.Get("/", GetAPIIndex)

	response, _, err := lib.GetTest(app, "/", nil)

	utils.AssertEqual(t, nil, err, "GET /")
	utils.AssertEqual(t, 200, response.StatusCode, "HTTP Status")
}
