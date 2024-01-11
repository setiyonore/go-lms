package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Publisher struct {
	publihserService service.Publisher
}

func NewPublihserHandler(publihserService service.Publisher) *Publisher {
	return &Publisher{publihserService: publihserService}
}

func (p *Publisher) GetAll(ctx *fiber.Ctx) error {
	publishers, err := p.publihserService.GetAll()
	if err != nil {
		response := helper.APIResponse("failed get publishers",
			fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("list of publisher", fiber.StatusOK, "success",
		entities.FormatterPublishers(publishers))
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (p *Publisher) GetById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	publisher, err := p.publihserService.GetById(id)
	if err != nil {
		response := helper.APIResponse("failed to get publisher", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("publisher", fiber.StatusOK, "success", entities.FormatterPublisher(publisher))
	return ctx.Status(fiber.StatusOK).JSON(response)
}
