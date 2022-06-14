package lib

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestHTTPRequest(t *testing.T) {
	request := HTTPRequest("GET", "/", map[string]string{"Content-Type": "application/json"})
	request2 := HTTPRequest("POST", "/", map[string]string{"Content-Type": "application/json"}, `{}`)
	utils.AssertEqual(t, "application/json", request.Header.Get("content-type"), "content-type")
	utils.AssertEqual(t, "application/json", request2.Header.Get("content-type"), "content-type")
}

func TestGetTest(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{})
	})
	response, body, err := GetTest(app, "/", nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, true, nil != body, "getting response code")
}

func TestPostTest(t *testing.T) {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{})
	})
	response, body, err := PostTest(app, "/", nil, `{}`)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, true, nil != body, "getting response code")
}

func TestPutTest(t *testing.T) {
	app := fiber.New()
	app.Put("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{})
	})
	response, body, err := PutTest(app, "/", nil, `{}`)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, true, nil != body, "getting response code")
}

func TestDeleteTest(t *testing.T) {
	app := fiber.New()
	app.Delete("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{})
	})
	response, body, err := DeleteTest(app, "/", nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, true, nil != body, "getting response code")
}
