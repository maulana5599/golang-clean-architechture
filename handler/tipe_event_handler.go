package handler

import "github.com/gofiber/fiber/v2"

type TipeEventHandler interface {
	Find(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}
