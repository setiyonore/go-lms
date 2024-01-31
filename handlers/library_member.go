package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LibraryMemberHandler struct {
	libraryMemberService service.LibraryMember
}

func NewLibraryMemberHandler(libraryMemberService service.LibraryMember) *LibraryMemberHandler {
	return &LibraryMemberHandler{libraryMemberService: libraryMemberService}
}

func (h *LibraryMemberHandler) GetLibraryMembers(c *fiber.Ctx) error {
	libraryMembers, err := h.libraryMemberService.GetLibraryMember()
	if err != nil {
		response := helper.APIResponse("failed to get library members",
			fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("list of library member", fiber.StatusOK, "success",
		entities.FormatLibraryMembers(libraryMembers))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *LibraryMemberHandler) GetLibraryMemberById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	libraryMember, err := h.libraryMemberService.GetLibraryMemberById(id)
	if err != nil {
		response := helper.APIResponse("failed to get library member",
			fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success", fiber.StatusOK, "librarry member",
		entities.FormatLibraryMember(libraryMember))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *LibraryMemberHandler) GetLibraryMemberByName(c *fiber.Ctx) error {
	var input entities.LibrarryMemberSearchByName
	data := entities.LibrarryMembers{}
	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("failed to parse data", fiber.StatusBadRequest,
			"error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	data, err = h.libraryMemberService.GetLibraryMemberByName(input.Name)
	if err != nil {
		response := helper.APIResponse("failed to find library member",
			fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("librarry member", fiber.StatusOK, "success",
		entities.FormatLibraryMember(data))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *LibraryMemberHandler) AddLibrarryMember(c *fiber.Ctx) error {
	var input entities.AddLibraryMemberInput
	err := c.BodyParser(&input)
	if err != nil {
		response := helper.APIResponse("failed parse data", fiber.StatusInternalServerError,
			"error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	err = h.libraryMemberService.AddLibrarryMember(input)
	if err != nil {
		response := helper.APIResponse("failed to save librarry member", fiber.StatusInternalServerError, "error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success save librarry member", fiber.StatusOK, "success", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}
