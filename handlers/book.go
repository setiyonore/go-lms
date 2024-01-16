package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	bookService service.Book
}

func NewBookHandler(bookService service.Book) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetBook()
	if err != nil {
		response := helper.APIResponse("failed to get books", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("list of books", fiber.StatusOK, "success", entities.FormatBooks(books))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *BookHandler) GetBookById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := h.bookService.GetBookById(id)
	if err != nil {
		response := helper.APIResponse("failed to find book", fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("book", fiber.StatusOK, "success", entities.FormatBook(book))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *BookHandler) AddBook(c *fiber.Ctx) error {
	var input entities.AddBookInput
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
	err = h.bookService.AddBook(input)
	if err != nil {
		response := helper.APIResponse("failed to save book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success save book", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var input entities.AddBookInput
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
	err = h.bookService.UpdateBook(id, input)
	if err != nil {
		response := helper.APIResponse("failed update book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success update book", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.bookService.DeleteBook(id)
	if err != nil {
		response := helper.APIResponse("failed delete book", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success delete book", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}
