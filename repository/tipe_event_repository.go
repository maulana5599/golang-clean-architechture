package repository

import (
	"golang_clean_architechture/model/domain"

	"github.com/gofiber/fiber/v2"
)

type TipeEventRepository interface {
	Find(ctx *fiber.Ctx) domain.TipeEvent
	FindAll(ctx *fiber.Ctx) []domain.TipeEvent
}
