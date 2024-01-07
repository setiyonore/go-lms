package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.User
}

func NewUserHandler(userService service.User) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	users, err := u.userService.GetAllUser()
	if err != nil {
		response := helper.APIResponse("Failed get user", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("List Of Users", fiber.StatusOK, "success", entities.FormatUsers(users))
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (u *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	Id, _ := strconv.Atoi(ctx.Params("id"))
	user, err := u.userService.GetUserById(Id)
	if err != nil {
		response := helper.APIResponse("Failed get user", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("User", fiber.StatusOK, "success", entities.FormatUser(user))
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (u *UserHandler) GetUserByEmail(ctx *fiber.Ctx) error {
	var input entities.GetUserByEmailInput
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
	var user entities.User
	user, err = u.userService.GetUserByEmail(input.Email)
	if err != nil {
		response := helper.APIResponse("Failed to get data", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("User", fiber.StatusOK, "success", entities.FormatUser(user))
	return ctx.Status(fiber.StatusOK).JSON(response)
}
