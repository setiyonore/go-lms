package route

import (
	"github.com/gofiber/fiber/v2"
	"go-lms/handlers"
)

func NewAuthorRoute(app fiber.Router) {

	app.Get("/authors", func(ctx *fiber.Ctx) error {
		return handlers.GetAuthor(ctx)
	})
}
