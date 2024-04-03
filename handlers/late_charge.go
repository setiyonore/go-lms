package handlers

import (
	"go-lms/entities"
	"go-lms/helper"
	"go-lms/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LateChargeHandler struct {
	latechargeService service.LateCharge
}

func NewLateChargeHandler(latechargeService service.LateCharge) *LateChargeHandler {
	return &LateChargeHandler{latechargeService: latechargeService}
}

func (h *LateChargeHandler) GetLateCharge(c *fiber.Ctx) error {
	lateCharge, err := h.latechargeService.GetLateCharge()
	if err != nil {
		response := helper.APIResponse("failed to get late charge", fiber.StatusInternalServerError,
			"error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("list of late charge", fiber.StatusOK,
		"success", entities.FormatLateCharges(lateCharge))
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *LateChargeHandler) GetLateChargeById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	lateCharge, err := h.latechargeService.GetLateChargeById(id)
	if err != nil {
		response := helper.APIResponse("failed to ge late charge",
			fiber.StatusInternalServerError,
			"error", nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := helper.APIResponse("late charge", fiber.StatusOK, "success",
		entities.FormatLateCharge(lateCharge))
	return c.Status(fiber.StatusOK).JSON(response)
}
