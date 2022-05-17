package service

import (
	"golang_clean_architechture/model/domain"
	"golang_clean_architechture/repository"

	"github.com/gofiber/fiber/v2"
)

type TipeEventServiceImpl struct {
	TipeEventRepository repository.TipeEventRepository
}

func NewTipeEventService(tipeEventRepository repository.TipeEventRepository) TipeEventService {
	return &TipeEventServiceImpl{
		TipeEventRepository: tipeEventRepository,
	}
}

func (s *TipeEventServiceImpl) Find(ctx *fiber.Ctx) domain.TipeEvent {
	tipeEvent := s.TipeEventRepository.Find(ctx)
	return tipeEvent
}

func (s *TipeEventServiceImpl) FindAll(ctx *fiber.Ctx) []domain.TipeEvent {
	tipeEvent := s.TipeEventRepository.FindAll(ctx)
	return tipeEvent
}
