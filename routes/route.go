package routes

import (
	"golang_clean_architechture/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(controller handler.TipeEventHandler, route *fiber.App) {

	v1 := route.Group("api/v1")

	v1.Get("/event", controller.Find)
	v1.Get("/eventall", controller.FindAll)
}
