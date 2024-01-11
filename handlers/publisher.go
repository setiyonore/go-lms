package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"

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
