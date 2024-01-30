package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

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
	response := helper.APIResponse("success", fiber.StatusOK, "library member",
		entities.FormatLibraryMember(libraryMember))
	return c.Status(fiber.StatusOK).JSON(response)
}
