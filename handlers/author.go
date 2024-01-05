package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AuthorHandler struct {
	authorService service.Author
}

func NewAuthorHandler(authorService service.Author) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}
func (h *AuthorHandler) GetAuthor(ctx *fiber.Ctx) error {
	//var authors []entities.Author
	//config.Database.Find(&authors)
	authors, err := h.authorService.GetAuthor()
	if err != nil {
		response := helper.APIResponse("Error to get author", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("List Of Author", fiber.StatusOK, "Success", entities.FormatAuthors(authors))
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthorHandler) GetAuthorByID(ctx *fiber.Ctx) error {
	Id, _ := strconv.Atoi(ctx.Params("id"))
	author, err := h.authorService.GetAuthorByID(Id)
	if err != nil {
		response := helper.APIResponse("Failed to Find Author", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Author", fiber.StatusOK, "Success", entities.FormatAuthor(author))
	return ctx.Status(fiber.StatusOK).JSON(response)
}
