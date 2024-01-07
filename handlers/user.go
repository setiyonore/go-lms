package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"
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
