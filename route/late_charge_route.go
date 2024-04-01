package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewLateChargeRoute(app fiber.Router) {
	db := *config.Database
	lateChargeRepository := repository.NewLateCharge(&db)
	lateChargeService := service.NewLateCharge(lateChargeRepository)
	LateChargeHandler := handlers.NewLateChargeHandler(lateChargeService)
	app.Get("/late_charges", func(c *fiber.Ctx) error {
		return LateChargeHandler.GetLateCharge(c)
	})
}
