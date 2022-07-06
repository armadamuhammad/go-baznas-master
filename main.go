package main

import (
	"api/app/config"
	"api/app/lib"
	"api/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	lib.LoadEnvironment(config.Environment)
}

// @title Baznas Master Service
// @version 1.0.0
// @description API Documentation
// @termsOfService https://dospecs.monstercode.net/en/guide/tnc.html
// @contact.name Armada Muhammad
// @contact.email armada_muhammad@student.uns.ac.id
// @host localhost:8000
// @schemes http
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
