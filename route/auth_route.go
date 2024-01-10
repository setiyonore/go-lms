package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/helper"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoute(app fiber.Router) {
	db := *config.Database
	authJwt := helper.NewAuth()
	userRepository := repository.NewUserRepository(&db)
	authService := service.NewAuth(userRepository)
	authHandler := handlers.NewAuthHandler(authService, authJwt)
	app.Post("/users/login", func(c *fiber.Ctx) error {
		return authHandler.Login(c)
	})
}
