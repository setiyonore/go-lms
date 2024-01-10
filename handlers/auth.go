package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.Auth
	authJwt     helper.Auth
}

func NewAuthHandler(authService service.Auth, autJwt helper.Auth) *AuthHandler {
	return &AuthHandler{authService: authService, authJwt: autJwt}
}

func (a *AuthHandler) Login(ctx *fiber.Ctx) error {
	var intput entities.LoginInput
	err := ctx.BodyParser(&intput)
	if err != nil {
		response := helper.APIResponse("failed parse data ", fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	validate := validator.New()
	err = validate.Struct(&intput)
	if err != nil {
		response := helper.APIResponse(helper.FormatterError(err.(validator.ValidationErrors)), fiber.StatusBadRequest, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	user, err := a.authService.Login(intput)
	if err != nil {
		response := helper.APIResponse("failed to login", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	token, err := a.authJwt.GenerateToken(user)
	if err != nil {
		response := helper.APIResponse("failed to login", fiber.StatusInternalServerError, "error", nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("success login", fiber.StatusOK, "success", fiber.Map{"token": token})
	return ctx.Status(fiber.StatusOK).JSON(response)
}
