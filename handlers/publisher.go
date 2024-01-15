package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/go-playground/validator/v10"
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

func (p *Publisher) AddPublisher(c *fiber.Ctx) error {
	input := entities.AddPublisherInput{}
	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("failed parse data", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	err = p.publihserService.AddPublisher(input)
	if err != nil {
		response := helper.APIResponse("failed to save publihser", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success save publisher", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (p *Publisher) UpdatePublisher(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var input entities.AddPublisherInput
	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("failed parse data", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	err = p.publihserService.UpdatePublisher(id, input)
	if err != nil {
		response := helper.APIResponse("failed update publisher", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success update publisher", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (p *Publisher) DeletePublisher(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := p.publihserService.DeletePublisher(id)
	if err != nil {
		response := helper.APIResponse("failed delete publisher", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success delete publihser", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}
