package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewBoookRoute(app fiber.Router) {
	db := *config.Database
	bookRepository := repository.NewBook(&db)
	bookService := service.NewBook(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)
	app.Get("/books", func(c *fiber.Ctx) error {
		return bookHandler.GetBooks(c)
	})
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		return bookHandler.GetBookById(c)
	})
	app.Post("/books", func(c *fiber.Ctx) error {
		return bookHandler.AddBook(c)
	})
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		return bookHandler.UpdateBook(c)
	})
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		return bookHandler.DeleteBook(c)
	})
	app.Get("/books/checkBookStatus/:id", func(c *fiber.Ctx) error {
		return bookHandler.CheckBookAvalable(c)
	})
}
