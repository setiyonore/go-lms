package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewBookBorrowingsRoute(app fiber.Router) {
	db := *config.Database
	bookBorrowingsRepository := repository.NewBookBorrowing(&db)
	bookRepossitory := repository.NewBook(&db)
	bookBorrowingService := service.NewBookBorrowing(bookBorrowingsRepository,
		bookRepossitory)
	bookBorrowingHandler := handlers.NewBookBorrowingHandler(bookBorrowingService)
	app.Get("/book_borrowings", func(c *fiber.Ctx) error {
		return bookBorrowingHandler.GetBookBorrowings(c)
	})
	app.Get("/book_borrowings/:id", func(c *fiber.Ctx) error {
		return bookBorrowingHandler.GetDetailBorrowing(c)
	})
	app.Post("/book_borrowings", func(c *fiber.Ctx) error {
		return bookBorrowingHandler.Add(c)
	})
}
