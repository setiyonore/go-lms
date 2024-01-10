package route

import (
	"go-lms/config"
	"go-lms/handlers"
	"go-lms/helper"
	"go-lms/repository"
	"go-lms/service"

	"github.com/gofiber/fiber/v2"
)

func NewUserRoute(app fiber.Router) {
	db := *config.Database
	authJwt := helper.NewAuth()
	userRepository := repository.NewUserRepository(&db)
	userService := service.NewUser(userRepository)
	userHandler := handlers.NewUserHandler(userService, authJwt)
	app.Get("/users", func(ctx *fiber.Ctx) error {
		return userHandler.GetAllUser(ctx)
	})
	app.Get("/users/:id", func(ctx *fiber.Ctx) error {
		return userHandler.GetUserById(ctx)
	})
	app.Post("/users/getByEmail", func(ctx *fiber.Ctx) error {
		return userHandler.GetUserByEmail(ctx)
	})
	app.Post("/users", func(ctx *fiber.Ctx) error {
		return userHandler.AddUser(ctx)
	})
	app.Put("/users/:id", func(ctx *fiber.Ctx) error {
		return userHandler.UpdateUser(ctx)
	})
	app.Delete("/users/:id", func(ctx *fiber.Ctx) error {
		return userHandler.DeleteUser(ctx)
	})
	app.Post("/users/login", func(c *fiber.Ctx) error {
		return userHandler.Login(c)
	})
}
