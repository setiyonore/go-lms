package route

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	publicRoute := app.Group("/api")
	NewAuthorRoute(publicRoute)
	NewUserRoute(publicRoute)
}
