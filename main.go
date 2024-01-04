package main

import (
	"go-lms/config"
	"go-lms/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Connect()
	app.Get("/authors", handlers.GetAuthor)
	log.Fatal(app.Listen(":8080"))
}
