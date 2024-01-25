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
	bookBorrowingService := service.NewBookBorrowing(bookBorrowingsRepository)
	bookBorrowingHandler := handlers.NewBookBorrowingHandler(bookBorrowingService)
	app.Get("/book_borrowings", func(c *fiber.Ctx) error {
		return bookBorrowingHandler.GetBookBorrowings(c)
	})
}
