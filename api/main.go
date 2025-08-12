package main

import (
	"github.com/Adityadangi14/centralized_logging_system/api/handler"
	"github.com/Adityadangi14/centralized_logging_system/api/initilizers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initilizers.LoadEnvVariables()
	initilizers.ConnectToPostgres()

}

func main() {
	app := fiber.New()

	app.Get("/logs", handler.GetLogsHandler)

	app.Listen(":4000")
}
