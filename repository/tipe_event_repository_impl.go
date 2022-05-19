package repository

import (
	"golang_clean_architechture/config"
	"golang_clean_architechture/model/domain"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TipeEventRepositoryImpl struct {
}

func NewTipeEventRepository() TipeEventRepository {
	return &TipeEventRepositoryImpl{}
}

func (repo TipeEventRepositoryImpl) Find(ctx *fiber.Ctx) domain.TipeEvent {
	var result domain.TipeEvent
	tx := config.DB.Debug().Raw("SELECT * FROM tipe_event").Scan(&result)

	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	return result
}

func (repo TipeEventRepositoryImpl) FindAll(ctx *fiber.Ctx) []domain.TipeEvent {
	var result []domain.TipeEvent
	tx := config.DB.Debug().Raw("SELECT * FROM tipe_event").Scan(&result)

	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	return result
}
