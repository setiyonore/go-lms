package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewItemBookRoute(app fiber.Router) {
	db := *config.Database
	itemBookRepository := repository.NewItemBook(&db)
	itemBookService := service.NewItemBook(itemBookRepository)
	ItemBookHandler := handlers.NewItemBookHandler(itemBookService)
	app.Get("/item_book", func(c *fiber.Ctx) error {
		return ItemBookHandler.GetItemBook(c)
	})
	app.Get("/item_book/:id", func(c *fiber.Ctx) error {
		return ItemBookHandler.GetItemBookById(c)
	})
}
