package route

import (
	"github.com/gofiber/fiber/v2"
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"
)

func NewAuthorRoute(app fiber.Router) {
	db := *config.Database
	authorRepository := repository.NewAuthor(&db)
	authorService := service.NewAuthor(authorRepository)
	authorHandler := handlers.NewAuthorHandler(authorService)
	app.Get("/authors", func(ctx *fiber.Ctx) error {
		return authorHandler.GetAuthor(ctx)
	})
	app.Get("/authors/:id", func(ctx *fiber.Ctx) error {
		return authorHandler.GetAuthorByID(ctx)
	})
}
