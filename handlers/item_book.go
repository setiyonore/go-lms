package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ItemBookHandler struct {
	itemBookService service.ItemBook
}

func NewItemBookHandler(itemBookService service.ItemBook) *ItemBookHandler {
	return &ItemBookHandler{itemBookService: itemBookService}
}

func (h *ItemBookHandler) GetItemBook(c *fiber.Ctx) error {
	books, err := h.itemBookService.GetItemBook()
	if err != nil {
		response := helper.APIResponse("failed to get item books", fiber.StatusInternalServerError,
			"eror", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("list of item books", fiber.StatusOK, "success",
		entities.FormatItemBooks(books))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ItemBookHandler) GetItemBookById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	itemBook, err := h.itemBookService.GetItemBookById(id)
	if err != nil {
		response := helper.APIResponse("failed to find item book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("item book", fiber.StatusOK, "success", entities.FormatItemBook(itemBook))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ItemBookHandler) AddItemBook(c *fiber.Ctx) error {
	var input entities.AddItemBookInput
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
	err = h.itemBookService.AddItemBook(input)
	if err != nil {
		response := helper.APIResponse("failed to save item book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success save item book", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ItemBookHandler) UpdateItemBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var input entities.AddItemBookInput
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
	err = h.itemBookService.UpdateItemBook(id, input)
	if err != nil {
		response := helper.APIResponse("failed update item book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success update item book", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ItemBookHandler) UpdateStatusItemBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	status, _ := strconv.Atoi(c.Params("status"))
	err := h.itemBookService.UpdateStatusItemBook(id, status)
	if err != nil {
		response := helper.APIResponse("failed update status item book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success update status item book", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}
