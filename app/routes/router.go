package routes

import (
	"api/app/controller"
	"api/app/controller/account"
	"api/app/controller/category"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()
	// services.InitRedis()

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)

	// Account
	api.Post("/accounts", account.PostAccount)
	api.Get("/accounts", account.GetAccount)
	api.Put("/accounts/:id", account.PutAccount)
	api.Get("/accounts/:id", account.GetAccountID)
	api.Delete("/accounts/:id", account.DeleteAccount)

	// Category
	api.Post("/categories", category.PostCategory)
	api.Get("/categories", category.GetCategory)
	api.Put("/categories/:id", category.PutCategory)
	api.Get("/categories/:id", category.GetCategoryID)
	api.Delete("/categories/:id", category.DeleteCategory)

}
