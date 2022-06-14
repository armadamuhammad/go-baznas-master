package controller

import (
	"api/app/lib"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// GetAPIInfo func
// @Summary show info response
// @Description show info response
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Failure 409 {object} lib.Response "conflict"
// @Failure 500 {object} lib.Response "internal error"
// @Router /info.json [get]
// @Tags API
func GetAPIInfo(c *fiber.Ctx) error {
	// example connection without services.InitDatabase()
	// database already initialized at app/routes/router.go
	// db := services.DB
	info := fiber.Map{
		"id":           viper.GetString("APP_ID"),
		"version":      viper.GetString("APP_VERSION"),
		"name":         viper.GetString("APP_NAME"),
		"description":  viper.GetString("APP_DESCRIPTION"),
		"dependencies": fiber.Map{},
	}

	return lib.OK(c, info)
}
