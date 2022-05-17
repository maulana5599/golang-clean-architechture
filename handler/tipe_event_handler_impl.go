package handler

import (
	"golang_clean_architechture/model/request"
	"golang_clean_architechture/model/response"
	"golang_clean_architechture/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TipeEventHandlerImpl struct {
	TipeEventService service.TipeEventService
}

func NewTipeEventHandler(tipeEventService service.TipeEventService) TipeEventHandler {
	return &TipeEventHandlerImpl{
		TipeEventService: tipeEventService,
	}
}

func (controller *TipeEventHandlerImpl) Find(ctx *fiber.Ctx) error {
	request := new(request.TipeEventRequest)

	if err := ctx.BodyParser(request); err != nil {
		return ctx.JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   "Body parser",
		})
	}

	Validator := validator.New()
	errValidate := Validator.Struct(request)

	if errValidate != nil {
		return ctx.JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": errValidate.Error,
			"error":   "Struct",
		})
	}

	eventResponse := controller.TipeEventService.Find(ctx)
	webResponse := response.TipeEventResponse{
		Status: http.StatusOK,
		Data:   eventResponse,
	}

	return ctx.JSON(fiber.Map{
		"message":  "Success",
		"response": webResponse,
	})

}

func (controller *TipeEventHandlerImpl) FindAll(ctx *fiber.Ctx) error {

	eventResponse := controller.TipeEventService.FindAll(ctx)
	webResponse := response.TipeEventResponse{
		Status: http.StatusOK,
		Data:   eventResponse,
	}

	return ctx.JSON(fiber.Map{
		"message":  "Success",
		"response": webResponse,
	})

}
