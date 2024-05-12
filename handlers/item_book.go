package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"

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
