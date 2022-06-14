package controller

import (
	"api/app/lib"

	"github.com/gofiber/fiber/v2"
)

// GetAPIIndex index page
// @Summary show basic response
// @Description show basic response
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} lib.Response "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Failure 409 {object} lib.Response "conflict"
// @Failure 500 {object} lib.Response "internal error"
// @Router / [get]
// @Tags API
func GetAPIIndex(c *fiber.Ctx) error {
	// example connection without services.InitDatabase()
	// database already initialized at app/routes/router.go
	// db := services.DB
	return lib.OK(c)
}
