package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

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
