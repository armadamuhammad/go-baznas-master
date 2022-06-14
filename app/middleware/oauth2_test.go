//go:build !integration
// +build !integration

package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestOauth2Authentication(t *testing.T) {
	app := fiber.New()
	app.Use(Oauth2Authentication)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("success")
	})

	response, err := app.Test(httptest.NewRequest("GET", "/", nil))
	utils.AssertEqual(t, nil, err, "oauth2 authenticated")
	utils.AssertEqual(t, 200, response.StatusCode, "oauth2 authenticated")
}
