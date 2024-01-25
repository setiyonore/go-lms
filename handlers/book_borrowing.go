package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

type BookBorrowingsHandler struct {
	bookBorrowingService service.BookBorrowings
}

func NewBookBorrowingHandler(bookBorrowingService service.BookBorrowings) *BookBorrowingsHandler {
	return &BookBorrowingsHandler{bookBorrowingService: bookBorrowingService}
}

func (h *BookBorrowingsHandler) GetBookBorrowings(c *fiber.Ctx) error {
	bookBorrowings, err := h.bookBorrowingService.GetBookBorrowing()
	if err != nil {
		response := helper.APIResponse("failed to get book borrowings",
			fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("list of book borrowings", fiber.StatusOK, "success",
		entities.FormatBookBorrowings(bookBorrowings))
	return c.Status(fiber.StatusOK).JSON(response)
}
