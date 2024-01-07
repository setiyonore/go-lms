package route

import (
	"github.com/gofiber/fiber/v2"
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/repository"
	"go-lms/service"
)

func NewUserRoute(app fiber.Router) {
	db := *config.Database
	userRepository := repository.NewUserRepository(&db)
	userService := service.NewUser(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	app.Get("/users", func(ctx *fiber.Ctx) error {
		return userHandler.GetAllUser(ctx)
	})
	app.Get("/users/:id", func(ctx *fiber.Ctx) error {
		return userHandler.GetUserById(ctx)
	})
}
