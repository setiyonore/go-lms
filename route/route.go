package route

import (
	"go-lms/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	publicRoute := app.Group("/api")
	privateRoute := app.Group("/api")
	privateRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Secret)},
	}))
	NewAuthorRoute(publicRoute)
	NewUserRoute(privateRoute)
}
