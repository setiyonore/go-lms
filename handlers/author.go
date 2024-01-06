package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"
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

func (h *AuthorHandler) AddAuthor(ctx *fiber.Ctx) error {
	var input entities.AddAuthorInput
	err := ctx.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("Failed parse data", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	err = h.authorService.AddAuhtor(input)
	if err != nil {
		response := helper.APIResponse("Failed to save author", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success save author", fiber.StatusOK, "success", nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthorHandler) UpdateAuthor(ctx *fiber.Ctx) error {
	Id, _ := strconv.Atoi(ctx.Params("id"))
	var input entities.AddAuthorInput
	err := ctx.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("Failed Parse data", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	err = h.authorService.UpdateAuthor(Id, input)
	if err != nil {
		response := helper.APIResponse("Failed Update author", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	response := helper.APIResponse("Success update author", fiber.StatusOK, "success", nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthorHandler) DeleteAuthor(ctx *fiber.Ctx) error {
	Id, _ := strconv.Atoi(ctx.Params("id"))
	err := h.authorService.DeleteAuthor(Id)
	if err != nil {
		response := helper.APIResponse("Failed delete author", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("Success delete author", fiber.StatusOK, "success", nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
