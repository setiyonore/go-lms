package route

import (
	"go-lms/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	apiRoute := app.Group("/api")

	publicRoute := apiRoute.Group("/")
	NewAuthRoute(publicRoute)
	privateRoute := apiRoute.Group("/", middleware.AuthMiddleware())
	NewLibraryMemberRoute(privateRoute)
	NewBookBorrowingsRoute(privateRoute)
	NewBoookRoute(privateRoute)
	NewPublihserRoute(privateRoute)
	NewAuthorRoute(privateRoute)
	NewUserRoute(privateRoute)
	NewLateChargeRoute(privateRoute)
}
