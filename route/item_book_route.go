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
	app.Post("/item_book", func(c *fiber.Ctx) error {
		return ItemBookHandler.AddItemBook(c)
	})
	app.Put("/item_book/:id", func(c *fiber.Ctx) error {
		return ItemBookHandler.UpdateItemBook(c)
	})
	app.Get("/item_book/update_status/:id/:status", func(c *fiber.Ctx) error {
		return ItemBookHandler.UpdateStatusItemBook(c)
	})
	app.Delete("/item_book/:id", func(c *fiber.Ctx) error {
		return ItemBookHandler.DeleteItemBook(c)
	})
	app.Get("/item_book/getByIdBook/:id", func(c *fiber.Ctx) error {
		return ItemBookHandler.GetItemBookByIdBook(c)
	})
	app.Get("/item_book/getByIdBookAvailable/:id", func(c *fiber.Ctx) error {
		return ItemBookHandler.GetItemBookByIdBookAvailable(c)
	})
}
