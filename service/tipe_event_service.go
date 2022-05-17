package service

import (
	"golang_clean_architechture/model/domain"

	"github.com/gofiber/fiber/v2"
)

type TipeEventService interface {
	Find(ctx *fiber.Ctx) domain.TipeEvent
	FindAll(ctx *fiber.Ctx) []domain.TipeEvent
}
