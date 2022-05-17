package main

import (
	"golang_clean_architechture/config"
	"golang_clean_architechture/handler"
	"golang_clean_architechture/repository"
	"golang_clean_architechture/routes"
	"golang_clean_architechture/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.DatabaseConnection()

	app := fiber.New()

	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	tipeEventRepository := repository.NewTipeEventRepository()
	tipeEventService := service.NewTipeEventService(tipeEventRepository)
	tipeEventController := handler.NewTipeEventHandler(tipeEventService)

	routes.RouteInit(tipeEventController, app)

	app.Listen(":8083")
}
