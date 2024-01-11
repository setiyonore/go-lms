package route

import (
	"go-lms/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	apiRoute := app.Group("/api")

	publicRoute := apiRoute.Group("/")
	NewAuthRoute(publicRoute)
	privateRoute := apiRoute.Group("/", middleware.AuthMiddleware())
	NewPublihserRoute(privateRoute)
	NewAuthorRoute(privateRoute)
	NewUserRoute(privateRoute)
}
