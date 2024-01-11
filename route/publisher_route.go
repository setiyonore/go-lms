package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewPublihserRoute(app fiber.Router) {
	db := *config.Database
	publisherRepository := repository.NewPublisherRepository(&db)
	publihserService := service.NewService(publisherRepository)
	publisherHandler := handlers.NewPublihserHandler(publihserService)
	app.Get("/publishers", func(c *fiber.Ctx) error {
		return publisherHandler.GetAll(c)
	})
}
