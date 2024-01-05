package handlers

import (
	"go-lms/service"

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
		return ctx.Status(fiber.StatusInternalServerError).JSON(authors)
	}
	return ctx.Status(fiber.StatusOK).JSON(authors)
}
